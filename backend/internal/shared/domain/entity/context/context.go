package context

import (
	"sync"
)

// ContextData data
type ContextData struct {
	sync.RWMutex
	Req  ContextRequestData         `json:",omitempty"` // data from http request
	Step map[string]ContextStepData `json:",omitempty"` // map[StepId]StepData
}

type ContextRequestData struct {
	Query map[string]any `json:",omitempty"` // map[queryVar]Value
	Json  map[string]any `json:",omitempty"` // map[jsonVar]Value
}

type ContextStepData struct {
	Var        map[string]any `json:",omitempty"` // map[Var]Value. Data from step variables
	StatusCode int            `json:",omitempty"` // status code for action rest type
	Data       any            `json:",omitempty"` // body response. For database in JSON form
	Output     map[string]any `json:",omitempty"` // map[OutputVar]Value
}

func (ctxData *ContextData) SetRequestQuery(query map[string]any) {
	ctxData.Lock()
	ctxData.Req.Query = query
	ctxData.Unlock()
}

func (ctxData *ContextData) SetRequestJson(json map[string]any) {
	ctxData.Lock()
	ctxData.Req.Json = json
	ctxData.Unlock()
}

func (ctxData *ContextData) GetStep(stepId string) ContextStepData {
	if ctxData.Step == nil {
		ctxData.Step = make(map[string]ContextStepData)
	}

	stepData, _ := ctxData.Step[stepId]
	return stepData
}

func (ctxData *ContextData) SetStepStatusCode(stepId string, statusCode int) {
	ctxData.Lock()
	stepData := ctxData.GetStep(stepId)
	stepData.StatusCode = statusCode
	ctxData.Step[stepId] = stepData
	ctxData.Unlock()
}

func (ctxData *ContextData) SetStepData(stepId string, data any) {
	ctxData.Lock()
	stepData := ctxData.GetStep(stepId)
	stepData.Data = data
	ctxData.Step[stepId] = stepData
	ctxData.Unlock()
}

func (ctxData *ContextData) SetStepVariable(stepId string, data map[string]any) {
	ctxData.Lock()
	stepData := ctxData.GetStep(stepId)
	stepData.Var = data
	ctxData.Step[stepId] = stepData
	ctxData.Unlock()
}

func (ctxData *ContextData) SetStepOutput(stepId string, data map[string]any) {
	ctxData.Lock()
	stepData := ctxData.GetStep(stepId)
	stepData.Output = data
	ctxData.Step[stepId] = stepData
	ctxData.Unlock()
}
