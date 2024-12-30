package handler

import (
	"context"

	adapterController "github.com/bayu-aditya/ideagate/backend/client/worker-rest/adapter/controller"
	entityEndpoint "github.com/bayu-aditya/ideagate/backend/core/model/entity/endpoint"
	"github.com/bayu-aditya/ideagate/backend/core/utils/errors"
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
		endpoint := entityEndpoint.Endpoint{}
		endpoint.FromPB(endpointPb)

		router.Handle(endpointPb.GetMethod(), endpointPb.GetPath(), h.handler(endpoint))
	}

	return nil
}

func (h *handler) handler(endpoint entityEndpoint.Endpoint) gin.HandlerFunc {
	return func(c *gin.Context) {
		mgr, _ := newManager(c, endpoint)
		mgr.RunHandler()
	}
}
