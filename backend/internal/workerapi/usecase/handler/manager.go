package handler

import (
	"context"
	"fmt"
	"time"

	adapterEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/adapter/endpoint"
	entityContext "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/context"
	entityEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/endpoint"
	handlerJob "github.com/bayu-aditya/ideagate/backend/internal/workerapi/usecase/handler/job"

	"github.com/bayu-aditya/ideagate/backend/internal/shared/domain/constant"
	"github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/datasource"
	"github.com/bayu-aditya/ideagate/backend/internal/workerapi/domain/entity"
	"github.com/bayu-aditya/ideagate/backend/pkg/utils/errors"
	"github.com/gin-gonic/gin"
)

type iManager interface {
	RunHandler()
}

type manager struct {
	ctxGin          *gin.Context
	ctx             context.Context
	ctxData         *entityContext.ContextData
	endpointAdapter adapterEndpoint.IEndpointAdapter
	response        entity.HttpResponse
	endpointId      string
	endpoint        entityEndpoint.Endpoint
	steps           map[string]entityEndpoint.Step //map[stepId]step
	isStepSuccess   map[string]bool                //map[stepId]isStepSuccess
}

func newManager(c *gin.Context, endpointAdapter adapterEndpoint.IEndpointAdapter, endpointId string) (iManager, error) {
	mgr := &manager{
		ctxGin:          c,
		ctx:             c.Request.Context(),
		ctxData:         new(entityContext.ContextData),
		endpointAdapter: endpointAdapter,
		endpointId:      endpointId,
		steps:           make(map[string]entityEndpoint.Step),
		isStepSuccess:   make(map[string]bool),
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

	return mgr, nil
}

func (m *manager) RunHandler() {
	// construct timeout channel based on endpoint detail setting
	timeoutChan := time.After(time.Duration(m.endpoint.Setting.TimeoutMs) * time.Millisecond)

	processWithChanFn := func() (chanResult chan handlerResult) {
		chanResult <- m.process()
		return
	}

	select {
	case result := <-processWithChanFn():
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

func (m *manager) process() (result handlerResult) {
	type jobFinishChanType struct {
		lastStepId string
		err        error
	}

	var (
		numWorkers    = m.endpoint.Setting.NumWorkers
		jobSeedsChan  = make(chan string, 1000) // value is stepId
		jobFinishChan = make(chan jobFinishChanType)
	)

	defer close(jobSeedsChan)
	defer close(jobFinishChan)

	// Step Request must be existed
	if _, isFound := m.steps[constant.StepIdRequest]; !isFound {
		return handlerResult{
			errors: []error{errors.New("step request not found")},
		}
	}

	// Spawn workers
	for i := 0; i < numWorkers; i++ {
		go func() {
			for {
				stepId := <-jobSeedsChan
				step, isStepFound := m.steps[stepId]
				if !isStepFound {
					jobFinishChan <- jobFinishChanType{
						lastStepId: stepId,
						err:        fmt.Errorf("step id (%s) is not found", stepId),
					}
					return
				}

				// check same step id cannot execute twice
				if !m.isStepSuccess[stepId] {
					continue
				}

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
					jobFinishChan <- jobFinishChanType{lastStepId: stepId, err: err}
					return
				}

				// start job
				jobOutput, err := job.Start()
				if err != nil {
					m.isStepSuccess[stepId] = false
					jobFinishChan <- jobFinishChanType{lastStepId: stepId, err: err}
					return
				}
				m.isStepSuccess[stepId] = true

				// determine next step job
				for _, nextStepId := range jobOutput.NextStepIds {
					if nextStepId == string(constant.JobTypeEnd) {
						jobFinishChan <- jobFinishChanType{lastStepId: stepId}
						return
					}

					jobSeedsChan <- nextStepId
				}
			}
		}()
	}

	// Seed first job by start type
	jobSeedsChan <- string(constant.JobTypeStart)

	// Wait until finish channel is published
	// and capture the data if no error
	jobFinish := <-jobFinishChan
	result.data = m.ctxData.Step[jobFinish.lastStepId].Data

	return
}
