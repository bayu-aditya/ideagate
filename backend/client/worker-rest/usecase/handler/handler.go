package handler

import (
	"context"

	adapterController "github.com/bayu-aditya/ideagate/backend/client/worker-rest/adapter/controller"
	"github.com/bayu-aditya/ideagate/backend/core/utils/errors"
	pbEndpoint "github.com/bayu-aditya/ideagate/backend/model/gen-go/core/endpoint"
	"github.com/gin-gonic/gin"
)

type IHandlerUsecase interface {
	GenerateEndpoint(ctx context.Context, router *gin.Engine) error
}

func New(adapterController adapterController.IControllerAdapter) IHandlerUsecase {
	return &handler{
		prefix:            "handler",
		adapterController: adapterController,
	}
}

type handler struct {
	prefix            string
	adapterController adapterController.IControllerAdapter
}

func (h *handler) GenerateEndpoint(ctx context.Context, router *gin.Engine) error {
	prefix := h.prefix + ".GenerateEndpoint"

	resultListEndpoint, err := h.adapterController.GetListEndpoint(ctx)
	if err != nil {
		return errors.Wrap(prefix, err, "get list endpoint")
	}

	for _, endpointPb := range resultListEndpoint.GetEndpoints() {
		router.Handle(endpointPb.GetMethod(), endpointPb.GetPath(), h.handler(endpointPb))
	}

	return nil
}

func (h *handler) handler(endpoint *pbEndpoint.Endpoint) gin.HandlerFunc {
	return func(c *gin.Context) {
		mgr, _ := newManager(c, endpoint)
		mgr.RunHandler()
	}
}
