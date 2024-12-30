package endpoint

import (
	"github.com/bayu-aditya/ideagate/backend/core/model/constant"
	pbEndpoint "github.com/bayu-aditya/ideagate/backend/model/gen-go/core/endpoint"
)

// Endpoint entity for table `endpoint`
type Endpoint struct {
	Id        string
	ProjectId string
	Method    string
	Path      string
	Setting   Setting
	Workflow  Workflow
}

func (e *Endpoint) FromPB(in *pbEndpoint.Endpoint) {
	e.Id = in.Id
	e.ProjectId = in.ProjectId
	e.Method = in.Method
	e.Path = in.Path

	e.Setting = Setting{}
	e.Setting.FromPB(in.Setting)

	e.Workflow = Workflow{}
	e.Workflow.FromPB(in.Workflow)
}

type Setting struct {
	Name        string         `json:",omitempty"`
	Description string         `json:",omitempty"`
	TimeoutMs   int64          `json:",omitempty"`
	NumWorkers  int            `json:",omitempty"`
	Request     SettingRequest `json:",omitempty"`
}

func (s *Setting) FromPB(in *pbEndpoint.Setting) {
	s.Name = in.Name
	s.Description = in.Description
	s.TimeoutMs = in.TimeoutMs
	s.NumWorkers = int(in.NumWorkers)

	s.Request = SettingRequest{}
	s.Request.FromPB(in.Request)
}

type SettingRequest struct {
	Query map[string]Variable `json:",omitempty"`
	Json  map[string]Variable `json:",omitempty"`
}

func (s *SettingRequest) FromPB(in *pbEndpoint.SettingRequest) {
	s.Query = make(map[string]Variable)
	for k, v := range in.Query {
		variable := Variable{}
		variable.FromPB(v)
		s.Query[k] = variable
	}

	s.Json = make(map[string]Variable)
	for k, v := range in.Json {
		variable := Variable{}
		variable.FromPB(v)
		s.Json[k] = variable
	}
}

type Workflow struct {
	Steps []Step `json:",omitempty"`
	Edges []Edge `json:",omitempty"`
}

func (w *Workflow) FromPB(in *pbEndpoint.Workflow) {
	w.Steps = make([]Step, len(in.Steps))
	for i, v := range in.Steps {
		step := Step{}
		step.FromPB(v)
		w.Steps[i] = step
	}

	w.Edges = make([]Edge, len(in.Edges))
	for i, v := range in.Edges {
		edge := Edge{}
		edge.FromPB(v)
		w.Edges[i] = edge
	}
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

func (s *Step) FromPB(in *pbEndpoint.Step) {
	s.Id = in.Id
	s.Name = in.Name
	s.Type = constant.JobType(in.Type)

	s.Variables = make(map[string]Variable)
	for k, v := range in.Variables {
		variable := Variable{}
		variable.FromPB(v)
		s.Variables[k] = variable
	}

	s.Action = Action{}
	s.Action.FromPB(in.Action)

	s.Outputs = make(map[string]Variable)
	for k, v := range in.Outputs {
		variable := Variable{}
		variable.FromPB(v)
		s.Outputs[k] = variable
	}

	s.Returns = make([]Return, len(in.Returns))
	for i, v := range in.Returns {
		ret := Return{}
		ret.FromPB(v)
		s.Returns[i] = ret
	}
}

type Action struct {
	DataSourceId string `json:",omitempty"` // reference datasource id

	End   *ActionEnd   `json:",omitempty"`
	Mysql *ActionMysql `json:",omitempty"`
	Rest  *ActionRest  `json:",omitempty"`
	Sleep *ActionSleep `json:",omitempty"`
}

func (a *Action) FromPB(in *pbEndpoint.Action) {
	a.DataSourceId = in.DataSourceId

	if in.End != nil {
		a.End = &ActionEnd{}
		a.End.FromPB(in.End)
	}

	if in.Mysql != nil {
		a.Mysql = &ActionMysql{}
		a.Mysql.FromPB(in.Mysql)
	}

	if in.Rest != nil {
		a.Rest = &ActionRest{}
		a.Rest.FromPB(in.Rest)
	}

	if in.Sleep != nil {
		a.Sleep = &ActionSleep{}
		a.Sleep.FromPB(in.Sleep)
	}
}

type ActionEnd struct {
	ReturnDataFromStepIds []string `json:",omitempty"`
}

func (a *ActionEnd) FromPB(in *pbEndpoint.ActionEnd) {
	a.ReturnDataFromStepIds = in.ReturnDataFromStepIds
}

type ActionMysql struct {
	Queries []Query `json:",omitempty"` // This queries will run in transaction
}

func (a *ActionMysql) FromPB(in *pbEndpoint.ActionMysql) {
	a.Queries = make([]Query, len(in.Queries))
	for i, v := range in.Queries {
		query := Query{}
		query.FromPB(v)
		a.Queries[i] = query
	}
}

type ActionRest struct {
	Path        Variable            `json:",omitempty"` // Text template supported. Ex: /user?id={{.Var.user_id}}
	Method      string              `json:",omitempty"` // Ex: GET, POST, PUT, etc
	Headers     map[string]Variable `json:",omitempty"` //
	RequestType string              `json:",omitempty"` // Ex: application/json
	RequestBody string              `json:",omitempty"` // REST body
}

func (a *ActionRest) FromPB(in *pbEndpoint.ActionRest) {
	a.Path.FromPB(in.Path)
	a.Method = in.Method

	a.Headers = make(map[string]Variable)
	for k, v := range in.Headers {
		variable := Variable{}
		variable.FromPB(v)
		a.Headers[k] = variable
	}

	a.RequestType = in.RequestType
	a.RequestBody = in.RequestBody
}

type ActionSleep struct {
	TimeoutMs int64 `json:",omitempty"`
}

func (a *ActionSleep) FromPB(in *pbEndpoint.ActionSleep) {
	a.TimeoutMs = in.TimeoutMs
}

type Query struct {
	Query      Variable   `json:"query,omitempty"`      // query template Ex: SELECT * FROM user WHERE user_id = ?
	Parameters []Variable `json:"parameters,omitempty"` // query parameters
}

func (q *Query) FromPB(in *pbEndpoint.Query) {
	q.Query.FromPB(in.Query)
	q.Parameters = make([]Variable, len(in.Parameters))
	for i, v := range in.Parameters {
		variable := Variable{}
		variable.FromPB(v)
		q.Parameters[i] = variable
	}
}

type Return struct {
	Id                string `json:",omitempty"`
	Name              string `json:",omitempty"`
	IsFinishCondition string `json:",omitempty"`
	NextStepId        string `json:",omitempty"`
}

func (r *Return) FromPB(in *pbEndpoint.Return) {
	r.Id = in.Id
	r.Name = in.Name
	r.IsFinishCondition = in.IsFinishCondition
	r.NextStepId = in.NextStepId
}

type Edge struct {
	Id          string `json:",omitempty"`
	ConditionId string `json:",omitempty"`
	Source      string `json:",omitempty"`
	Dest        string `json:",omitempty"`
}

func (e *Edge) FromPB(in *pbEndpoint.Edge) {
	e.Id = in.Id
	e.ConditionId = in.ConditionId
	e.Source = in.Source
	e.Dest = in.Dest
}
