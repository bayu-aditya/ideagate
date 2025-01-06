package sql

import "context"

type SqlAdapter interface {
	FindApplications(ctx context.Context)
}
