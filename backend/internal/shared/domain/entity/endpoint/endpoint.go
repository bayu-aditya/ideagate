package endpoint

import "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/constant"

// Endpoint entity for table `endpoint`
type Endpoint struct {
	Id        string
	ProjectId string
	Method    string
	Path      string
	Setting   Setting
	Workflow  Workflow
}

type Setting struct {
	Name        string         `json:",omitempty"`
	Description string         `json:",omitempty"`
	TimeoutMs   int64          `json:",omitempty"`
	NumWorkers  int            `json:",omitempty"`
	Request     SettingRequest `json:",omitempty"`
}

type SettingRequest struct {
	Query map[string]Variable `json:",omitempty"`
	Json  map[string]Variable `json:",omitempty"`
}

type Workflow struct {
	Steps []Step `json:",omitempty"`
}

type Step struct {
	Id        string              `json:",omitempty"` // step id
	Name      string              `json:",omitempty"`
	Type      constant.JobType    `json:",omitempty"`
	Variables map[string]Variable `json:",omitempty"`
	Action    Action              `json:",omitempty"`
	Outputs   map[string]Variable `json:",omitempty"`
	Returns   []Return            `json:",omitempty"`
}

type Action struct {
	DataSourceId string              `json:",omitempty"` // reference datasource id
	Queries      []Query             `json:",omitempty"` // for database type. This queries will run in transaction
	Path         Variable            `json:",omitempty"` // for rest type. Text template supported. Ex: /user?id={{.var.user_id}}
	Method       string              `json:",omitempty"` // for rest type. Ex: GET, POST, PUT, etc
	Headers      map[string]Variable `json:",omitempty"` // for rest type
	RequestType  string              `json:",omitempty"` // for rest type. Ex: application/json
	RequestBody  string              `json:",omitempty"` // for rest body
}

type Query struct {
	Query      Variable   `json:"query,omitempty"`      // query template Ex: SELECT * FROM user WHERE user_id = ?
	Parameters []Variable `json:"parameters,omitempty"` // query parameters
}

type Return struct {
	Id                string `json:",omitempty"`
	Name              string `json:",omitempty"`
	IsFinishCondition string `json:",omitempty"`
	NextStepId        string `json:",omitempty"`
}
