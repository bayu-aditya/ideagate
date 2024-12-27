package handler

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	adapterEndpointMock "github.com/bayu-aditya/ideagate/backend/core/adapter/endpoint/_mock"
	"github.com/bayu-aditya/ideagate/backend/core/model/constant"
	entityEndpoint "github.com/bayu-aditya/ideagate/backend/core/model/entity/endpoint"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Manager - Process", func() {
	Context("Linear", func() {
		It("start - sleep - sleep - end", func() {
			t := GinkgoT()

			gin.SetMode(gin.TestMode)
			resp := httptest.NewRecorder()
			mockCtxGin, _ := gin.CreateTestContext(resp)
			mockCtxGin.Request = &http.Request{
				URL: &url.URL{},
			}

			mockAdapterEndpoint := adapterEndpointMock.NewIEndpointAdapter(t)
			mockEndpointId := "mock_endpoint_id"
			mockWorkflow := entityEndpoint.Workflow{
				Steps: []entityEndpoint.Step{
					{
						Id:   constant.StepIdStart,
						Type: constant.JobTypeStart,
					},
					{
						Id:   "sleep_1",
						Type: constant.JobTypeSleep,
						Action: entityEndpoint.Action{
							Sleep: &entityEndpoint.ActionSleep{
								TimeoutMs: 1000,
							},
						},
					},
					{
						Id:   "sleep_2",
						Type: constant.JobTypeSleep,
						Action: entityEndpoint.Action{
							Sleep: &entityEndpoint.ActionSleep{
								TimeoutMs: 500,
							},
						},
					},
					{
						Id:   "end",
						Type: constant.JobTypeEnd,
						Action: entityEndpoint.Action{
							End: &entityEndpoint.ActionEnd{},
						},
					},
				},
				Edges: []entityEndpoint.Edge{
					{Id: "edge_1", Source: constant.StepIdStart, Dest: "sleep_1"},
					{Id: "edge_2", Source: "sleep_1", Dest: "sleep_2"},
					{Id: "edge_3", Source: "sleep_2", Dest: constant.StepIdEnd},
				},
			}

			mockAdapterEndpoint.EXPECT().GetEndpoint(mock.Anything, mockEndpointId).Once().
				Return(entityEndpoint.Endpoint{Workflow: mockWorkflow}, nil)

			manager, err := newManager(mockCtxGin, mockAdapterEndpoint, mockEndpointId)
			if err != nil {
				t.Error("new manager failed", err)
			}

			manager.RunHandler()
			Expect(resp.Code).To(Equal(http.StatusOK))
		})
	})
	Context("Parallel", func() {
		// start -- sleep1(3000) -- sleep2(3000) -- sleep5(500) -- sleep6(1500) -- end
		//       |- sleep3(1000) -- sleep4(1000) -|
		Context("start - sleep - end", func() {
			var (
				mockCtxGin       *gin.Context
				t                = GinkgoT()
				httpRecorder     = httptest.NewRecorder()
				mockDataEndpoint = entityEndpoint.Endpoint{
					Id: "mock_endpoint_id",
					Setting: entityEndpoint.Setting{
						NumWorkers: 1,
						TimeoutMs:  8100,
					},
					Workflow: entityEndpoint.Workflow{
						Steps: []entityEndpoint.Step{
							{
								Id:   constant.StepIdStart,
								Type: constant.JobTypeStart,
							},
							{
								Id:   "sleep_1",
								Type: constant.JobTypeSleep,
								Action: entityEndpoint.Action{
									Sleep: &entityEndpoint.ActionSleep{
										TimeoutMs: 3000,
									},
								},
							},
							{
								Id:   "sleep_2",
								Type: constant.JobTypeSleep,
								Action: entityEndpoint.Action{
									Sleep: &entityEndpoint.ActionSleep{
										TimeoutMs: 3000,
									},
								},
							},
							{
								Id:   "sleep_3",
								Type: constant.JobTypeSleep,
								Action: entityEndpoint.Action{
									Sleep: &entityEndpoint.ActionSleep{
										TimeoutMs: 1000,
									},
								},
							},
							{
								Id:   "sleep_4",
								Type: constant.JobTypeSleep,
								Action: entityEndpoint.Action{
									Sleep: &entityEndpoint.ActionSleep{
										TimeoutMs: 1000,
									},
								},
							},
							{
								Id:   "sleep_5",
								Type: constant.JobTypeSleep,
								Action: entityEndpoint.Action{
									Sleep: &entityEndpoint.ActionSleep{
										TimeoutMs: 500,
									},
								},
							},
							{
								Id:   "sleep_6",
								Type: constant.JobTypeSleep,
								Action: entityEndpoint.Action{
									Sleep: &entityEndpoint.ActionSleep{
										TimeoutMs: 1500,
									},
								},
							},
							{
								Id:   "end",
								Type: constant.JobTypeEnd,
								Action: entityEndpoint.Action{
									End: &entityEndpoint.ActionEnd{},
								},
							},
						},
						Edges: []entityEndpoint.Edge{
							{Id: "edge_1", Source: constant.StepIdStart, Dest: "sleep_1"},
							{Id: "edge_2", Source: "sleep_1", Dest: "sleep_2"},
							{Id: "edge_3", Source: "sleep_2", Dest: "sleep_5"},
							{Id: "edge_4", Source: constant.StepIdStart, Dest: "sleep_3"},
							{Id: "edge_5", Source: "sleep_3", Dest: "sleep_4"},
							{Id: "edge_6", Source: "sleep_4", Dest: "sleep_5"},
							{Id: "edge_7", Source: "sleep_5", Dest: "sleep_6"},
							{Id: "edge_4", Source: "sleep_6", Dest: constant.StepIdEnd},
						},
					},
				}
			)

			BeforeEach(func() {
				gin.SetMode(gin.TestMode)
				mockCtxGin, _ = gin.CreateTestContext(httpRecorder)
				mockCtxGin.Request = &http.Request{
					URL: &url.URL{},
				}
			})

			It("single worker: 10 secs", func() {
				mockDataEndpoint.Setting.NumWorkers = 1
				mockDataEndpoint.Setting.TimeoutMs = 10100 // must be finished around 10 secs

				mockAdapterEndpoint := adapterEndpointMock.NewIEndpointAdapter(t)
				mockAdapterEndpoint.EXPECT().GetEndpoint(mock.Anything, mockDataEndpoint.Id).Once().Return(mockDataEndpoint, nil)

				manager, err := newManager(mockCtxGin, mockAdapterEndpoint, mockDataEndpoint.Id)
				if err != nil {
					t.Error("new manager failed", err)
				}

				manager.RunHandler()
				Expect(httpRecorder.Code).To(Equal(http.StatusOK))
			})

			It("multiple worker: 8 secs", func() {
				mockDataEndpoint.Setting.NumWorkers = 3
				mockDataEndpoint.Setting.TimeoutMs = 8100 // must be finished around 8 secs

				mockAdapterEndpoint := adapterEndpointMock.NewIEndpointAdapter(t)
				mockAdapterEndpoint.EXPECT().GetEndpoint(mock.Anything, mockDataEndpoint.Id).Once().Return(mockDataEndpoint, nil)

				manager, err := newManager(mockCtxGin, mockAdapterEndpoint, mockDataEndpoint.Id)
				if err != nil {
					t.Error("new manager failed", err)
				}

				manager.RunHandler()
				Expect(httpRecorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
