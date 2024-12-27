package handler

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/bayu-aditya/ideagate/backend/client/worker-rest/model"
	handlerJob "github.com/bayu-aditya/ideagate/backend/client/worker-rest/usecase/handler/job"
	adapterEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/adapter/endpoint"
	"github.com/bayu-aditya/ideagate/backend/internal/shared/domain/constant"
	entityContext "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/context"
	"github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/datasource"
	entityEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/endpoint"
	"github.com/bayu-aditya/ideagate/backend/pkg/utils/errors"
	"github.com/bayu-aditya/ideagate/backend/pkg/utils/pubsub"
	"github.com/gin-gonic/gin"
)

type iManager interface {
	RunHandler()
}

type manager struct {
	mu                    sync.RWMutex
	ctxGin                *gin.Context
	ctx                   context.Context
	ctxData               *entityContext.ContextData
	response              model.HttpResponse
	endpoint              entityEndpoint.Endpoint
	steps                 map[string]entityEndpoint.Step // map[stepId]step
	stepStatus            map[string]StepStatusType      // map[stepId]StepStatus
	pubSub                pubsub.IPubSub
	pubSubTopicStepStatus string
	edgesNext             map[string][]string // map[stepId]nextStepIds
	edgesPrev             map[string][]string // map[stepId]prevStepIds
}

func newManager(c *gin.Context, endpointAdapter adapterEndpoint.IEndpointAdapter, endpointId string) (iManager, error) {
	mgr := &manager{
		ctxGin:                c,
		ctx:                   c.Request.Context(),
		ctxData:               new(entityContext.ContextData),
		steps:                 make(map[string]entityEndpoint.Step),
		stepStatus:            make(map[string]StepStatusType),
		pubSub:                pubsub.New(),
		pubSubTopicStepStatus: "step_status",
		edgesNext:             make(map[string][]string),
		edgesPrev:             make(map[string][]string),
	}

	// construct endpoint detail
	endpointDetail, err := endpointAdapter.GetEndpoint(mgr.ctx, endpointId)
	if err != nil {
		mgr.response.AddErrors(err).GinErrorInternal(mgr.ctxGin)
		return nil, err
	}
	mgr.endpoint = endpointDetail

	// construct map of steps
	for _, step := range endpointDetail.Workflow.Steps {
		mgr.steps[step.Id] = step
	}

	// construct edges
	for _, edge := range endpointDetail.Workflow.Edges {
		mgr.edgesNext[edge.Source] = append(mgr.edgesNext[edge.Source], edge.Dest)
		mgr.edgesPrev[edge.Dest] = append(mgr.edgesPrev[edge.Dest], edge.Source)
	}

	return mgr, nil
}

func (m *manager) RunHandler() {
	// construct timeout channel based on endpoint detail setting
	timeoutMs := m.endpoint.Setting.TimeoutMs
	if timeoutMs == 0 {
		timeoutMs = 10000 // default 10 sec
	}
	timeoutChan := time.After(time.Duration(timeoutMs) * time.Millisecond)

	chanResult := make(chan handlerResult)
	go func() {
		chanResult <- m.process()
	}()

	select {
	case result := <-chanResult:
		if len(result.errors) > 0 {
			m.response.GinErrorInternal(m.ctxGin)
			return
		}

		m.response.AddData(result.data).GinSuccess(m.ctxGin)
		return

	case <-timeoutChan:
		m.response.GinErrorTimeout(m.ctxGin)
		return
	}
}

type handlerResult struct {
	data   interface{}
	errors []error
}

func (m *manager) process() handlerResult {
	type jobFinishChanType struct {
		stepId string
		data   any
		err    error
	}

	var (
		numWorkers    = m.endpoint.Setting.NumWorkers
		jobSeedsChan  = make(chan string, 1000)      // value is stepId
		jobFinishChan = make(chan jobFinishChanType) // value is based on latest step
	)

	defer func() {
		close(jobSeedsChan)
		close(jobFinishChan)
	}()

	if numWorkers == 0 {
		numWorkers = 1
	}

	// Step Request must be existed
	if _, isFound := m.steps[constant.StepIdStart]; !isFound {
		return handlerResult{
			errors: []error{errors.New("step request not found")},
		}
	}

	// Spawn workers
	for i := 0; i < numWorkers; i++ {
		go func() {
			for stepId := range jobSeedsChan {
				workerName := fmt.Sprintf("worker-%d", i)

				data, err := m.stepWorker(stepId, workerName)
				if err != nil {
					jobFinishChan <- jobFinishChanType{stepId: stepId, err: err}
					return
				}

				if m.steps[stepId].Type == constant.JobTypeEnd {
					jobFinishChan <- jobFinishChanType{stepId: stepId, data: data}
					return
				}

				for _, nextStepId := range m.edgesNext[stepId] {
					if m.getOrSetStepStatusToQueue(nextStepId) != StepStatusUnknown {
						continue
					}
					jobSeedsChan <- nextStepId
				}
			}
		}()
	}

	// Seed first job by start type
	jobSeedsChan <- constant.StepIdStart

	// Wait until finish channel is published
	finishData := <-jobFinishChan

	if finishData.err != nil {
		return handlerResult{
			errors: []error{finishData.err},
		}
	}

	return handlerResult{
		data: finishData.data,
	}
}

func (m *manager) stepWorker(stepId, workerName string) (any, error) {
	step, isStepFound := m.steps[stepId]
	if !isStepFound {
		return nil, fmt.Errorf("step id '%s' not found", stepId)
	}

	m.setStepStatus(stepId, StepStatusInit)

	// wait all previous steps must be already executed
	m.waitAllDependencies(stepId, workerName)

	// construct job
	jobInput := handlerJob.StartInput{
		Ctx:        m.ctx,
		GinCtx:     m.ctxGin,
		DataCtx:    m.ctxData,
		DataSource: &datasource.DataSource{}, // TODO fill this
		Endpoint:   m.endpoint,
		Step:       step,
	}

	job, err := handlerJob.New(step.Type, jobInput)
	if err != nil {
		return nil, err
	}

	// start job
	m.setStepStatus(stepId, StepStatusProgress)

	output, err := job.Start()
	if err != nil {
		m.setStepStatus(stepId, StepStatusError)
		return nil, err
	}

	m.setStepStatus(stepId, StepStatusFinish)

	return output.Data, nil
}

func (m *manager) waitAllDependencies(stepId, workerName string) {
	stepsWait := make(map[string]bool)                            // key is list of step id must be waited
	numBufferChanSubscriber := len(m.endpoint.Workflow.Steps) * 3 // num step * 3

	subscriber := m.pubSub.Subscribe(context.Background(), m.pubSubTopicStepStatus, workerName, pubsub.SubscribeSetting{
		NumBufferChan: numBufferChanSubscriber,
	})
	defer subscriber.Close()

	for _, prevStepId := range m.edgesPrev[stepId] {
		stepStatus := m.getStepStatus(prevStepId)
		if stepStatus == StepStatusSkip || stepStatus == StepStatusFinish {
			continue
		}

		stepsWait[prevStepId] = true
	}

	if len(stepsWait) > 0 {
		for data := range subscriber.GetData() {
			stepIdUpdated := data.(string)

			if _, isExist := stepsWait[stepIdUpdated]; !isExist {
				continue
			}

			stepStatus := m.getStepStatus(stepIdUpdated)
			if stepStatus != StepStatusSkip && stepStatus != StepStatusFinish {
				continue
			}

			delete(stepsWait, stepIdUpdated)

			if len(stepsWait) == 0 {
				break
			}
		}
	}
}

func (m *manager) getStepStatus(stepId string) (status StepStatusType) {
	m.mu.RLock()
	status = m.stepStatus[stepId]
	m.mu.RUnlock()

	return
}

func (m *manager) getOrSetStepStatusToQueue(stepId string) (status StepStatusType) {
	m.mu.Lock()
	status = m.stepStatus[stepId]

	// if status is unknown (default) then set to queue
	if status == StepStatusUnknown {
		m.stepStatus[stepId] = StepStatusQueue
	}

	m.mu.Unlock()
	return
}

func (m *manager) setStepStatus(stepId string, status StepStatusType) {
	m.mu.Lock()
	m.stepStatus[stepId] = status
	m.mu.Unlock()

	m.pubSub.Publish(m.ctx, m.pubSubTopicStepStatus, stepId)
}