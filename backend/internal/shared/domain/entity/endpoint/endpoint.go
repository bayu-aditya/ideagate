package endpoint

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
	Request     SettingRequest `json:",omitempty"`
}

type SettingRequest struct {
	Query map[string]Variable `json:",omitempty"`
	Json  map[string]Variable `json:",omitempty"`
}

type Workflow struct {
	Variables  map[string]Variable `json:",omitempty"`
	Entrypoint string              `json:",omitempty"` // step id for entrypoint
	Steps      []Step              `json:",omitempty"`
}

type Step struct {
	Id        string              `json:",omitempty"` // step id
	Name      string              `json:",omitempty"`
	Variables map[string]Variable `json:",omitempty"`
	Actions   []Action            `json:",omitempty"`
	Outputs   map[string]Variable `json:",omitempty"`
	Returns   []Return            `json:",omitempty"`
}

type Variable struct {
	Type     string `json:",omitempty"`
	Required bool   `json:",omitempty"`
	Value    string `json:",omitempty"`
	Default  string `json:",omitempty"`
}

type Action struct {
	Id           string                 `json:",omitempty"` // action id
	DataSourceId string                 `json:",omitempty"` // reference datasource id
	Query        string                 `json:",omitempty"` // for database type. Text template supported. Ex: SELECT * FROM user WHERE user_id = {{.var.user_id}}
	Path         string                 `json:",omitempty"` // for rest type. Text template supported. Ex: /user?id={{.var.user_id}}
	Method       string                 `json:",omitempty"` // for rest type. Ex: GET, POST, PUT, etc
	Headers      map[string]interface{} `json:",omitempty"` // for rest type
	RequestType  string                 `json:",omitempty"` // for rest type. Ex: application/json
	RequestBody  string                 `json:",omitempty"` // for rest body
}

type Return struct {
	Id                string `json:",omitempty"`
	Name              string `json:",omitempty"`
	IsFinishCondition string `json:",omitempty"`
	NextStepId        string `json:",omitempty"`
}
