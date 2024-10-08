package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContextData_GetValues(t *testing.T) {
	type args struct {
		query string
	}

	mockFields := ContextData{}
	mockFields.SetRequestQuery(map[string]interface{}{
		"var_1": 123,
	})

	tests := []struct {
		name   string
		fields ContextData
		args   args
		want   string
	}{
		{
			name:   "static value",
			fields: mockFields,
			args: args{
				query: "static value",
			},
			want: "static value",
		},
		{
			name:   "unknown field",
			fields: mockFields,
			args: args{
				query: "{{.unknown}}",
			},
			want: "",
		},
		{
			name:   "request query found",
			fields: mockFields,
			args: args{
				query: "{{.req.query.var_1}}",
			},
			want: "123",
		},
		{
			name:   "request query not found",
			fields: mockFields,
			args: args{
				query: "{{.req.query.unknown}}",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctxData := &tt.fields
			if got := ctxData.GetValues(tt.args.query); got != tt.want {
				t.Errorf("GetValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContextData_SetStepActionData(t *testing.T) {
	type args struct {
		stepId   string
		actionId string
		data     ContextStepDataAction
	}
	tests := []struct {
		name   string
		fields ContextData
		args   args
		want   ContextData
	}{
		{
			name: "success - step is empty",
			fields: ContextData{
				Step: nil,
			},
			args: args{
				stepId:   "mock-step-id",
				actionId: "mock-action-id",
				data: ContextStepDataAction{
					StatusCode: 200,
					Body: map[string]interface{}{
						"foo": "bar",
					},
				},
			},
			want: ContextData{
				Step: map[string]ContextStepData{
					"mock-step-id": {
						Actions: map[string]ContextStepDataAction{
							"mock-action-id": {
								StatusCode: 200,
								Body: map[string]interface{}{
									"foo": "bar",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "success - previous step exist",
			fields: ContextData{
				Step: map[string]ContextStepData{
					"mock-step-id-before": {
						Actions: map[string]ContextStepDataAction{
							"mock-action-id-before": {
								StatusCode: 123,
							},
						},
					},
				},
			},
			args: args{
				stepId:   "mock-step-id",
				actionId: "mock-action-id",
				data: ContextStepDataAction{
					StatusCode: 200,
					Body: map[string]interface{}{
						"foo": "bar",
					},
				},
			},
			want: ContextData{
				Step: map[string]ContextStepData{
					"mock-step-id-before": {
						Actions: map[string]ContextStepDataAction{
							"mock-action-id-before": {
								StatusCode: 123,
							},
						},
					},
					"mock-step-id": {
						Actions: map[string]ContextStepDataAction{
							"mock-action-id": {
								StatusCode: 200,
								Body: map[string]interface{}{
									"foo": "bar",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "success - current step exist and action empty",
			fields: ContextData{
				Step: map[string]ContextStepData{
					"mock-step-id": {
						Actions: nil,
					},
				},
			},
			args: args{
				stepId:   "mock-step-id",
				actionId: "mock-action-id",
				data: ContextStepDataAction{
					StatusCode: 200,
					Body: map[string]interface{}{
						"foo": "bar",
					},
				},
			},
			want: ContextData{
				Step: map[string]ContextStepData{
					"mock-step-id": {
						Actions: map[string]ContextStepDataAction{
							"mock-action-id": {
								StatusCode: 200,
								Body: map[string]interface{}{
									"foo": "bar",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctxData := tt.fields
			ctxData.SetStepActionData(tt.args.stepId, tt.args.actionId, tt.args.data)
			assert.EqualExportedValues(t, tt.want, ctxData)
		})
	}
}
