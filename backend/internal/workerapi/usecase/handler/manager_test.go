package handler

import (
	mockAdapterEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/adapter/endpoint/_mock"
	"github.com/bayu-aditya/ideagate/backend/internal/shared/domain/constant"
	entityEndpoint "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/endpoint"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"net/url"
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

			mockAdapterEndpoint := mockAdapterEndpoint.NewIEndpointAdapter(t)
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
							TimeoutMs: 1000,
						},
					},
					{
						Id:   "sleep_2",
						Type: constant.JobTypeSleep,
						Action: entityEndpoint.Action{
							TimeoutMs: 500,
						},
					},
					{
						Id:   "end",
						Type: constant.JobTypeEnd,
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
		//
		// total duration must be 8000 ms
		It("start - sleep - end", func() {
			t := GinkgoT()

			gin.SetMode(gin.TestMode)
			resp := httptest.NewRecorder()
			mockCtxGin, _ := gin.CreateTestContext(resp)
			mockCtxGin.Request = &http.Request{
				URL: &url.URL{},
			}

			mockAdapterEndpoint := mockAdapterEndpoint.NewIEndpointAdapter(t)
			mockEndpoint := entityEndpoint.Endpoint{
				Id: "mock_endpoint_id",
				Setting: entityEndpoint.Setting{
					NumWorkers: 3,
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
								TimeoutMs: 3000,
							},
						},
						{
							Id:   "sleep_2",
							Type: constant.JobTypeSleep,
							Action: entityEndpoint.Action{
								TimeoutMs: 3000,
							},
						},
						{
							Id:   "sleep_3",
							Type: constant.JobTypeSleep,
							Action: entityEndpoint.Action{
								TimeoutMs: 1000,
							},
						},
						{
							Id:   "sleep_4",
							Type: constant.JobTypeSleep,
							Action: entityEndpoint.Action{
								TimeoutMs: 1000,
							},
						},
						{
							Id:   "sleep_5",
							Type: constant.JobTypeSleep,
							Action: entityEndpoint.Action{
								TimeoutMs: 500,
							},
						},
						{
							Id:   "sleep_6",
							Type: constant.JobTypeSleep,
							Action: entityEndpoint.Action{
								TimeoutMs: 1500,
							},
						},
						{
							Id:   "end",
							Type: constant.JobTypeEnd,
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

			mockAdapterEndpoint.EXPECT().GetEndpoint(mock.Anything, mockEndpoint.Id).Once().
				Return(mockEndpoint, nil)

			manager, err := newManager(mockCtxGin, mockAdapterEndpoint, mockEndpoint.Id)
			if err != nil {
				t.Error("new manager failed", err)
			}

			manager.RunHandler()
			Expect(resp.Code).To(Equal(http.StatusOK))
		})
	})
})
