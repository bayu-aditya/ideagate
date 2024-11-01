package endpoint

import (
	"reflect"
	"testing"

	entityContext "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/context"
)

func TestVariable_GetValue(t *testing.T) {
	mockStepId := "mockStepId"
	mockCtxData := &entityContext.ContextData{
		Req: entityContext.ContextRequestData{
			Query: map[string]interface{}{
				"query_1": "value_query_1",
				"query_2": "value_query_2",
			},
			Json: map[string]interface{}{
				"json_1": "value_json_1",
				"json_2": map[string]interface{}{
					"json_2_a": 123,
				},
			},
		},
		Step: map[string]entityContext.ContextStepData{
			mockStepId: {
				Var: map[string]interface{}{
					"var_1": "value_var_1",
					"var_2": "value_var_2",
				},
				StatusCode: 200,
				Data:       nil,
				Output:     nil,
			},
		},
	}

	type fields struct {
		Type     string
		Required bool
		Value    interface{}
		Default  interface{}
	}
	type args struct {
		stepId  string
		ctxData *entityContext.ContextData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "test 123",
			fields: fields{
				Value: "test 123",
			},
			args: args{
				stepId:  mockStepId,
				ctxData: mockCtxData,
			},
			want: "test 123",
		},
		{
			name: "test 123 {{.Ctx.Step.mockStepId.Var.var_2}}",
			fields: fields{
				Value: "test 123 {{.Ctx.Step.mockStepId.Var.var_2}}",
			},
			args: args{
				stepId:  mockStepId,
				ctxData: mockCtxData,
			},
			want: "test 123 value_var_2",
		},
		{
			name: "test 123 {{.Ctx.Step.mockStepId.Var.var_unknown}}: default",
			fields: fields{
				Value:    "test 123 {{.Ctx.Step.mockStepId.Var.var_unknown}}",
				Required: false,
				Default:  "value_default",
			},
			args: args{
				stepId:  mockStepId,
				ctxData: mockCtxData,
			},
			want: "test 123 ",
		},
		{
			name: "test 123 {{.Ctx.Step.mockStepId.Var.var_unknown}}: required default",
			fields: fields{
				Value:    "test 123 {{.Ctx.Step.mockStepId.Var.var_unknown}}",
				Required: true,
				Default:  "value_default",
			},
			args: args{
				stepId:  mockStepId,
				ctxData: mockCtxData,
			},
			want: "test 123 ",
		},
		{
			name: "{{.Ctx.Step.mockStepId.Var.var_unknown}}",
			fields: fields{
				Value:    "{{.Ctx.Step.mockStepId.Var.var_unknown}}",
				Required: false,
				Default:  "value_default",
			},
			args: args{
				stepId:  mockStepId,
				ctxData: mockCtxData,
			},
			want: "",
		},
		{
			name: "{{.Ctx.Step.mockStepId.Var.var_unknown}}: required default",
			fields: fields{
				Value:    "{{.Ctx.Step.mockStepId.Var.var_unknown}}",
				Required: true,
				Default:  "value_default",
				Type:     string(VariableTypeString),
			},
			args: args{
				stepId:  mockStepId,
				ctxData: mockCtxData,
			},
			want: "value_default",
		},

		{
			name: ".Step.StatusCode: found",
			fields: fields{
				Value:    "{{.Step.StatusCode}}",
				Required: true,
				Default:  500,
				Type:     string(VariableTypeInt),
			},
			args: args{
				stepId:  mockStepId,
				ctxData: mockCtxData,
			},
			want: 200,
		},
		{
			name: ".Step.StatusCode: not found, but have default",
			fields: fields{
				Value:    "{{.Step.StatusCode.Unknown}}",
				Required: true,
				Default:  500,
				Type:     string(VariableTypeInt),
			},
			args: args{
				stepId:  mockStepId,
				ctxData: mockCtxData,
			},
			want: 500,
		},
		{
			name: ".Req.Json.json_2.json_2_a",
			fields: fields{
				Value:    "{{.Req.Json.json_2.json_2_a}}",
				Required: true,
				Default:  "value_default",
				Type:     string(VariableTypeInt),
			},
			args: args{
				stepId:  mockStepId,
				ctxData: mockCtxData,
			},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Variable{
				Type:     tt.fields.Type,
				Required: tt.fields.Required,
				Value:    tt.fields.Value,
				Default:  tt.fields.Default,
			}
			got, err := v.GetValue(tt.args.stepId, tt.args.ctxData)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}
