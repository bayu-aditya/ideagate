package handler

import (
	"context"
	"github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/datasource"
	"github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/endpoint"
	"github.com/bayu-aditya/ideagate/backend/internal/workerapi/domain/entity"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_jobRest_Start(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	type args struct {
		ctx        context.Context
		ctxData    *entity.ContextData
		dataSource datasource.DataSource
		step       endpoint.Step
		action     endpoint.Action
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantCtxData *entity.ContextData
		mockFunc    func()
	}{
		{
			name: "negative - invalid method",
			args: args{
				ctx:     context.TODO(),
				ctxData: &entity.ContextData{},
				dataSource: datasource.DataSource{
					Config: datasource.Config{
						Host: "https://mockhost.com/api",
					},
				},
				step: endpoint.Step{
					Id: "mock-step-id",
				},
				action: endpoint.Action{
					Id:     "mock-action-id",
					Method: "unknown",
					Path:   "/user/detail",
				},
			},
			mockFunc:    func() {},
			wantErr:     true,
			wantCtxData: &entity.ContextData{},
		},
		{
			name: "success - empty context data",
			args: args{
				ctx:     context.TODO(),
				ctxData: &entity.ContextData{},
				dataSource: datasource.DataSource{
					Config: datasource.Config{
						Host: "https://mockhost.com/api",
					},
				},
				step: endpoint.Step{
					Id: "mock-step-id",
				},
				action: endpoint.Action{
					Id:     "mock-action-id",
					Method: "GET",
					Path:   "/user/detail",
				},
			},
			mockFunc: func() {
				httpmock.RegisterResponder("GET", "https://mockhost.com/api/user/detail",
					httpmock.NewStringResponder(200, `{
						"data": {
							"id": "mock-user-id",
							"name": "mock-user-name"
						}
					}`))
			},
			wantCtxData: &entity.ContextData{
				Step: map[string]entity.ContextStepData{
					"mock-step-id": {
						Actions: map[string]entity.ContextStepDataAction{
							"mock-action-id": {
								StatusCode: 200,
								Body: map[string]interface{}{
									"data": map[string]interface{}{
										"id":   "mock-user-id",
										"name": "mock-user-name",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "success - path is template",
			args: args{
				ctx: context.TODO(),
				ctxData: &entity.ContextData{
					Variable: map[string]interface{}{
						"user_id": 123,
					},
				},
				dataSource: datasource.DataSource{
					Config: datasource.Config{
						Host: "https://mockhost.com/api",
					},
				},
				step: endpoint.Step{
					Id: "mock-step-id",
				},
				action: endpoint.Action{
					Id:     "mock-action-id",
					Method: "GET",
					Path:   "/user/detail?user_id={{.var.user_id}}",
				},
			},
			mockFunc: func() {
				httpmock.RegisterResponder("GET", "https://mockhost.com/api/user/detail?user_id=123",
					httpmock.NewStringResponder(200, `{
						"data": {
							"id": "mock-user-id",
							"name": "mock-user-name"
						}
					}`))
			},
			wantCtxData: &entity.ContextData{
				Variable: map[string]interface{}{
					"user_id": 123,
				},
				Step: map[string]entity.ContextStepData{
					"mock-step-id": {
						Actions: map[string]entity.ContextStepDataAction{
							"mock-action-id": {
								StatusCode: 200,
								Body: map[string]interface{}{
									"data": map[string]interface{}{
										"id":   "mock-user-id",
										"name": "mock-user-name",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "success - path is template not found variable",
			args: args{
				ctx:     context.TODO(),
				ctxData: &entity.ContextData{},
				dataSource: datasource.DataSource{
					Config: datasource.Config{
						Host: "https://mockhost.com/api",
					},
				},
				step: endpoint.Step{
					Id: "mock-step-id",
				},
				action: endpoint.Action{
					Id:     "mock-action-id",
					Method: "GET",
					Path:   "/user/detail?user_id={{.var.user_id}}",
				},
			},
			mockFunc: func() {
				httpmock.RegisterResponder("GET", "https://mockhost.com/api/user/detail?user_id=",
					httpmock.NewBytesResponder(500, nil))
			},
			wantCtxData: &entity.ContextData{
				Step: map[string]entity.ContextStepData{
					"mock-step-id": {
						Actions: map[string]entity.ContextStepDataAction{
							"mock-action-id": {
								StatusCode: 500,
								Body:       nil,
							},
						},
					},
				},
			},
		},
		{
			name: "success - 500 status code",
			args: args{
				ctx:     context.TODO(),
				ctxData: &entity.ContextData{},
				dataSource: datasource.DataSource{
					Config: datasource.Config{
						Host: "https://mockhost.com/api",
					},
				},
				step: endpoint.Step{
					Id: "mock-step-id",
				},
				action: endpoint.Action{
					Id:     "mock-action-id",
					Method: "GET",
					Path:   "/user/detail",
				},
			},
			mockFunc: func() {
				httpmock.RegisterResponder("GET", "https://mockhost.com/api/user/detail",
					httpmock.NewBytesResponder(500, nil))
			},
			wantCtxData: &entity.ContextData{
				Step: map[string]entity.ContextStepData{
					"mock-step-id": {
						Actions: map[string]entity.ContextStepDataAction{
							"mock-action-id": {
								StatusCode: 500,
								Body:       nil,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jobRest{}
			tt.mockFunc()
			if err := j.Start(tt.args.ctx, tt.args.ctxData, tt.args.dataSource, tt.args.step, tt.args.action); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.EqualExportedValues(t, tt.wantCtxData, tt.args.ctxData)
		})
	}
}
