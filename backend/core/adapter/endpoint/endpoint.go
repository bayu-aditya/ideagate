package endpoint

import (
	"context"

	"github.com/bayu-aditya/ideagate/backend/core/model/entity/endpoint"
)

type IEndpointAdapter interface {
	GetListEndpoint(ctx context.Context) ([]endpoint.Endpoint, error)
	GetEndpoint(ctx context.Context, id string) (endpoint.Endpoint, error)
}
