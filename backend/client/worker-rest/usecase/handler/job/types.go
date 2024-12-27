package job

import (
	"context"

	entityContext "github.com/bayu-aditya/ideagate/backend/core/model/entity/context"
	entityDataSource "github.com/bayu-aditya/ideagate/backend/core/model/entity/datasource"
	entityEndpoint "github.com/bayu-aditya/ideagate/backend/core/model/entity/endpoint"
	"github.com/gin-gonic/gin"
)

type StartInput struct {
	Ctx        context.Context
	GinCtx     *gin.Context
	DataCtx    *entityContext.ContextData
	DataSource *entityDataSource.DataSource
	Endpoint   entityEndpoint.Endpoint
	Step       entityEndpoint.Step
}

type StartOutput struct {
	NextStepIds []string
	Data        any
}
