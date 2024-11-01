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
)

func (v *Variable) GetValue(stepId string, ctxData *entityContext.ContextData) (interface{}, error) {
	value := v.Value
	varType := VariableType(v.Type)

	// get value from context
	if valStr, ok := v.Value.(string); ok {
		value = v.getValueFromTemplate(stepId, ctxData, valStr)
	}

	// check is value empty
	if v.isEmptyValue(value) && v.Required {
		value = v.Default
	}

	// parse value by type
	value, err := v.parseValueByType(value, varType)
	if err != nil {
		return nil, err
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

func (v *Variable) getValueFromTemplate(stepId string, ctxData *entityContext.ContextData, templateValue string) string {
	tmpl, err := template.New("").Parse(templateValue)

	if err != nil {
		return ""
	}

	type dataTemplateType struct {
		Step entityContext.ContextStepData
		Ctx  *entityContext.ContextData
		Req  entityContext.ContextRequestData
	}

	data := dataTemplateType{
		Step: ctxData.Step[stepId],
		Ctx:  ctxData,
		Req:  ctxData.Req,
	}

	var resultBuffer bytes.Buffer
	if err = tmpl.Execute(&resultBuffer, data); err != nil {
		return ""
	}

	result := resultBuffer.String()
	result = strings.ReplaceAll(result, "<no value>", "")

	return result
}

func (v *Variable) parseValueByType(value interface{}, varType VariableType) (interface{}, error) {
	switch varType {
	case VariableTypeString:
		return cast.ToStringE(value)

	case VariableTypeInt:
		return cast.ToIntE(value)

	case VariableTypeFloat:
		return cast.ToFloat32E(value)

	case VariableTypeBool:
		return cast.ToBoolE(value)
	}

	return value, nil
}

func (v *Variable) isEmptyValue(value interface{}) bool {
	return reflect.ValueOf(value).IsZero()
}
