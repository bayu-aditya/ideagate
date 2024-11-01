package handler

import (
	"context"
	"testing"

	adapterEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/adapter/endpoint"
	entityContext "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/context"
	entityEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/endpoint"

	"github.com/bayu-aditya/ideagate/backend/internal/workerapi/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_manager_process(t *testing.T) {
	type fields struct {
		ctxGin          *gin.Context
		ctx             context.Context
		ctxData         *entityContext.ContextData
		endpointAdapter adapterEndpoint.IEndpointAdapter
		response        entity.HttpResponse
		endpointId      string
		endpoint        entityEndpoint.Endpoint
		steps           map[string]entityEndpoint.Step
		isStepSuccess   map[string]bool
	}
	tests := []struct {
		name       string
		fields     fields
		wantResult handlerResult
	}{
		{
			name:       "request only",
			fields:     fields{},
			wantResult: handlerResult{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &manager{
				ctxGin:          tt.fields.ctxGin,
				ctx:             tt.fields.ctx,
				ctxData:         tt.fields.ctxData,
				endpointAdapter: tt.fields.endpointAdapter,
				response:        tt.fields.response,
				endpointId:      tt.fields.endpointId,
				endpoint:        tt.fields.endpoint,
				steps:           tt.fields.steps,
				isStepSuccess:   tt.fields.isStepSuccess,
			}
			assert.Equalf(t, tt.wantResult, m.process(), "process()")
		})
	}
}
