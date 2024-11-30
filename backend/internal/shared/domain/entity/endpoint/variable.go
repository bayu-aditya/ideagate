package endpoint

import (
	"bytes"
	"reflect"
	"strings"

	"text/template"

	entityContext "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/context"
	"github.com/spf13/cast"
)

type Variable struct {
	Type     string      `json:",omitempty"` // related with VariableType
	Required bool        `json:",omitempty"`
	Value    interface{} `json:",omitempty"`
	Default  interface{} `json:",omitempty"`
}

type VariableType string

var (
	VariableTypeString VariableType = "string"
	VariableTypeInt    VariableType = "int"
	VariableTypeFloat  VariableType = "float"
	VariableTypeBool   VariableType = "bool"
	VariableTypeObject VariableType = "object"
)

func (v *Variable) GetValue(stepId string, ctxData *entityContext.ContextData) (interface{}, error) {
	value := v.Value
	varType := VariableType(v.Type)

	// get value from context
	if valStr, ok := v.Value.(string); ok {
		value = v.getValueFromTemplate(stepId, ctxData, valStr)
	}

	// parse value by type
	value, err := v.parseValueByType(value, varType)
	if err != nil {
		return nil, err
	}

	// check is value empty
	if v.isEmptyValue(value) && v.Required {
		// set into default and parse the default value
		value = v.Default
		value, err = v.parseValueByType(value, varType)
		if err != nil {
			return nil, err
		}
	}

	return value, nil
}

func (v *Variable) GetValueString(stepId string, ctxData *entityContext.ContextData) (string, error) {
	value, err := v.GetValue(stepId, ctxData)
	if err != nil {
		return "", err
	}

	return cast.ToStringE(value)
}

func (v *Variable) getValueFromTemplate(stepId string, ctxData *entityContext.ContextData, templateValue string) interface{} {
	tmpl, err := template.New("").Parse(templateValue)

	if err != nil {
		return nil
	}

	type dataTemplateType struct {
		Req  entityContext.ContextRequestData
		Step map[string]entityContext.ContextStepData
		Var  map[string]any
		Data entityContext.ContextStepDataBody
	}

	data := dataTemplateType{
		Req:  ctxData.Req,
		Step: ctxData.Step,
		Var:  ctxData.Step[stepId].Var,
		Data: ctxData.Step[stepId].Data,
	}

	var resultBuffer bytes.Buffer
	if err = tmpl.Execute(&resultBuffer, data); err != nil {
		return nil
	}

	result := resultBuffer.String()
	result = strings.ReplaceAll(result, "<no value>", "")

	if result == "" {
		return nil
	}

	return result
}

func (v *Variable) parseValueByType(value interface{}, varType VariableType) (interface{}, error) {
	if value == nil {
		return nil, nil
	}

	switch varType {
	case VariableTypeString:
		return cast.ToStringE(value)

	case VariableTypeInt:
		return cast.ToInt64E(value)

	case VariableTypeFloat:
		return cast.ToFloat64E(value)

	case VariableTypeBool:
		return cast.ToBoolE(value)

	case VariableTypeObject:
		return value, nil
	}

	return value, nil
}

func (v *Variable) isEmptyValue(value interface{}) bool {
	if value == nil {
		return true
	}

	return reflect.ValueOf(value).IsZero()
}
