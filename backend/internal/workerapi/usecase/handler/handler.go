package handler

import (
	"context"
	adapterEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/adapter/endpoint"
	entityDataSource "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/datasource"
	entityEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/endpoint"
	"github.com/bayu-aditya/ideagate/backend/internal/workerapi/domain/entity"
	"github.com/bayu-aditya/ideagate/backend/pkg/utils/errors"
	"github.com/bayu-aditya/ideagate/backend/pkg/utils/template"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"time"
)

type IHandlerUsecase interface {
	GenerateEndpoint(ctx context.Context, router *gin.Engine) error
}

func New() IHandlerUsecase {
	return &handler{
		prefix:          "handler",
		endpointAdapter: nil, // TODO fill this
	}
}

type handler struct {
	prefix          string
	endpointAdapter adapterEndpoint.IEndpointAdapter
}

func (h *handler) GenerateEndpoint(ctx context.Context, router *gin.Engine) error {
	prefix := h.prefix + ".GenerateEndpoint"

	endpoints, err := h.endpointAdapter.GetListEndpoint(ctx)
	if err != nil {
		return errors.Wrap(prefix, err, "get list endpoint")
	}

	for _, endpointItem := range endpoints {
		router.Handle(endpointItem.Method, endpointItem.Path, h.handler(endpointItem.Id))
	}

	return nil
}

func (h *handler) handler(endpointId string) gin.HandlerFunc {

	var (
		dataSources = map[string]entityDataSource.DataSource{} // map[dataSourceId]dataSource  TODO construct this
	)

	return func(c *gin.Context) {
		ctx := c.Request.Context()

		// get endpoint detail
		endpointDetail, errGetEndpoint := h.endpointAdapter.GetEndpoint(ctx, endpointId)
		if errGetEndpoint != nil {
			// TODO status code 500
			return
		}

		var (
			entryPoint     = endpointDetail.Workflow.Entrypoint
			timeoutMs      = endpointDetail.Setting.TimeoutMs
			settingRequest = endpointDetail.Setting.Request
			steps          = map[string]entityEndpoint.Step{} // map[currentStepId]step TODO construct this

		)

		ctx, cancel := context.WithTimeout(ctx, time.Duration(timeoutMs)*time.Second)
		defer cancel()

		// construct context data
		ctxData := &entity.ContextData{}
		if err := constructRequestCtxData(c, ctxData, settingRequest); err != nil {
			// TODO status code 400
			return
		}

		// run step
		currentStepId := entryPoint
		for {
			step, isStepExist := steps[currentStepId]
			if !isStepExist {
				// TODO status code 500, step ID {:currentStepId} not exist
				return
			}

			// run action concurrently
			for _, action := range step.Actions {
				// get datasource by ID
				dataSource, isDataSourceExist := dataSources[action.DataSourceId]
				if !isDataSourceExist {
					// TODO status code 500, data source ID {:step.DataSource.RefId} not exist
					return
				}

				// running job
				job := newJob(dataSource.Type)
				if job == nil {
					// TODO status code 500, job type {:dataSource.Type} not exist
					return
				}

				if err := job.Start(ctx, ctxData, dataSource, step, action); err != nil {
					// TODO status code 500, job {:currentStepId} failed
					return
				}
			}

			// run returns
			for _, returnItem := range step.Returns {
				// check condition based on context data
				isReturnStr := template.GetValue(ctxData, returnItem.IsFinishCondition)
				if cast.ToBool(isReturnStr) {
					return
				}

				// next step
				currentStepId = returnItem.NextStepId
			}
		}
	}
}
