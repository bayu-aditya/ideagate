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
	Edges []Edge `json:",omitempty"`
}

type Step struct {
	Id        string              `json:",omitempty"` // step id
	Name      string              `json:",omitempty"`
	Type      constant.JobType    `json:",omitempty"`
	Variables map[string]Variable `json:",omitempty"`
	Action    Action              `json:",omitempty"`
	Outputs   map[string]Variable `json:",omitempty"`
	Returns   []Return            `json:",omitempty"` // deprecated instead using Workflow.Edges
}

type Action struct {
	DataSourceId string `json:",omitempty"` // reference datasource id

	End   *ActionEnd   `json:",omitempty"`
	Mysql *ActionMysql `json:",omitempty"`
	Rest  *ActionRest  `json:",omitempty"`
	Sleep *ActionSleep `json:",omitempty"`
}

type ActionEnd struct {
	ReturnDataFromStepIds []string `json:",omitempty"`
}

type ActionMysql struct {
	Queries []Query `json:",omitempty"` // This queries will run in transaction
}

type ActionRest struct {
	Path        Variable            `json:",omitempty"` // Text template supported. Ex: /user?id={{.Var.user_id}}
	Method      string              `json:",omitempty"` // Ex: GET, POST, PUT, etc
	Headers     map[string]Variable `json:",omitempty"` //
	RequestType string              `json:",omitempty"` // Ex: application/json
	RequestBody string              `json:",omitempty"` // REST body
}

type ActionSleep struct {
	TimeoutMs int64 `json:",omitempty"`
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

type Edge struct {
	Id          string `json:",omitempty"`
	ConditionId string `json:",omitempty"`
	Source      string `json:",omitempty"`
	Dest        string `json:",omitempty"`
}
