package job

import (
	"context"

	entityContext "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/context"
	entityDataSource "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/datasource"
	entityEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/endpoint"
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
}
