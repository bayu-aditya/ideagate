package entity

import (
	"encoding/json"
	"github.com/bayu-aditya/ideagate/backend/pkg/utils/template"
	"sync"
)

// ContextData data
type ContextData struct {
	sync.RWMutex
	isJsonFormInvalidated bool
	jsonForm              map[string]interface{}

	Request  ContextRequestData         `json:"req,omitempty"`  // data from http request
	Variable map[string]interface{}     `json:"var,omitempty"`  // map[Var]Value. Data from workflow variables
	Step     map[string]ContextStepData `json:"step,omitempty"` // map[StepId]StepData
}

type ContextRequestData struct {
	Query map[string]interface{} `json:"query,omitempty"` // map[queryVar]Value
	Json  map[string]interface{} `json:"json,omitempty"`  // map[jsonVar]Value
}

type ContextStepData struct {
	Variable map[string]interface{}           `json:"var,omitempty"`     // map[Var]Value. Data from step variables
	Actions  map[string]ContextStepDataAction `json:"actions,omitempty"` // map[ActionId]
	Output   map[string]interface{}           `json:"output,omitempty"`  // map[OutputVar]Value
}

type ContextStepDataAction struct {
	StatusCode int         `json:"status_code,omitempty"` // status code for action rest type
	Body       interface{} `json:"body,omitempty"`        // body response. For database in JSON form
}

func (ctxData *ContextData) parseJson() {
	if !ctxData.isJsonFormInvalidated {
		return
	}

	jsonBytes, _ := json.Marshal(ctxData)
	_ = json.Unmarshal(jsonBytes, &ctxData.jsonForm)

	ctxData.isJsonFormInvalidated = false
}

func (ctxData *ContextData) SetRequestQuery(query map[string]interface{}) {
	ctxData.Lock()
	ctxData.Request.Query = query
	ctxData.isJsonFormInvalidated = true
	ctxData.Unlock()
}

func (ctxData *ContextData) SetRequestJson(json map[string]interface{}) {
	ctxData.Lock()
	ctxData.Request.Json = json
	ctxData.isJsonFormInvalidated = true
	ctxData.Unlock()
}

func (ctxData *ContextData) SetVariable(json map[string]interface{}) {
	ctxData.Lock()
	ctxData.Request.Json = json
	ctxData.isJsonFormInvalidated = true
	ctxData.Unlock()
}

func (ctxData *ContextData) SetStepOutput(stepId string, data map[string]interface{}) {
	ctxData.Lock()
	stepData, isFound := ctxData.Step[stepId]
	if !isFound {
		stepData = ContextStepData{}
	}
	stepData.Output = data
	ctxData.Step[stepId] = stepData
	ctxData.isJsonFormInvalidated = true
	ctxData.Unlock()
}

func (ctxData *ContextData) SetStepActionData(stepId, actionId string, data ContextStepDataAction) {
	ctxData.Lock()

	if ctxData.Step == nil {
		ctxData.Step = map[string]ContextStepData{}
	}

	step, isExist := ctxData.Step[stepId]
	if !isExist {
		ctxData.Step[stepId] = ContextStepData{}
	}

	if step.Actions == nil {
		step.Actions = map[string]ContextStepDataAction{}
	}

	if _, isActionExist := step.Actions[actionId]; !isActionExist {
		step.Actions[actionId] = ContextStepDataAction{}
	}

	step.Actions[actionId] = data

	ctxData.Step[stepId] = step
	ctxData.isJsonFormInvalidated = true

	ctxData.Unlock()
}

func (ctxData *ContextData) GetValues(query string) string {
	ctxData.RLock()
	ctxData.parseJson()
	value := template.GetValue(ctxData.jsonForm, query)
	ctxData.RUnlock()
	return value
}
