package handler

import (
	"context"

	adapterEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/adapter/endpoint"
	"github.com/bayu-aditya/ideagate/backend/pkg/utils/errors"
	"github.com/gin-gonic/gin"
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
	return func(c *gin.Context) {
		mgr, _ := newManager(c, h.endpointAdapter, endpointId)
		mgr.RunHandler()
	}
}
