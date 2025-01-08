package model

import (
	"time"

	"github.com/jackc/pgtype"
)

type Application struct {
	ID          string
	ProjectID   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
}

func (a *Application) TableName() string {
	return "application"
}

type Endpoint struct {
	ID            string
	ApplicationID string
	ProjectID     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Method        string
	Path          string
	Name          string
	Description   string
	Settings      pgtype.JSONB `gorm:"type:jsonb"`
}

func (e *Endpoint) TableName() string {
	return "endpoint"
}

type Workflow struct {
	Version       int
	EndpointID    string
	ApplicationID string
	ProjectID     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Data          pgtype.JSONB `gorm:"type:jsonb"`
}

func (w *Workflow) TableName() string {
	return "workflow"
}
