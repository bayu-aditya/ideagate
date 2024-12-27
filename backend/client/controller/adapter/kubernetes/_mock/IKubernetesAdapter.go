// Code generated by mockery v2.42.1. DO NOT EDIT.

package mockery

import (
	context "context"

	corev1 "k8s.io/api/core/v1"

	mock "github.com/stretchr/testify/mock"

	v1 "k8s.io/api/apps/v1"
)

// IKubernetesAdapter is an autogenerated mock type for the IKubernetesAdapter type
type IKubernetesAdapter struct {
	mock.Mock
}

type IKubernetesAdapter_Expecter struct {
	mock *mock.Mock
}

func (_m *IKubernetesAdapter) EXPECT() *IKubernetesAdapter_Expecter {
	return &IKubernetesAdapter_Expecter{mock: &_m.Mock}
}

// CreateDeployment provides a mock function with given fields: ctx, deployment
func (_m *IKubernetesAdapter) CreateDeployment(ctx context.Context, deployment v1.Deployment) error {
	ret := _m.Called(ctx, deployment)

	if len(ret) == 0 {
		panic("no return value specified for CreateDeployment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, v1.Deployment) error); ok {
		r0 = rf(ctx, deployment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IKubernetesAdapter_CreateDeployment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDeployment'
type IKubernetesAdapter_CreateDeployment_Call struct {
	*mock.Call
}

// CreateDeployment is a helper method to define mock.On call
//   - ctx context.Context
//   - deployment v1.Deployment
func (_e *IKubernetesAdapter_Expecter) CreateDeployment(ctx interface{}, deployment interface{}) *IKubernetesAdapter_CreateDeployment_Call {
	return &IKubernetesAdapter_CreateDeployment_Call{Call: _e.mock.On("CreateDeployment", ctx, deployment)}
}

func (_c *IKubernetesAdapter_CreateDeployment_Call) Run(run func(ctx context.Context, deployment v1.Deployment)) *IKubernetesAdapter_CreateDeployment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(v1.Deployment))
	})
	return _c
}

func (_c *IKubernetesAdapter_CreateDeployment_Call) Return(_a0 error) *IKubernetesAdapter_CreateDeployment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IKubernetesAdapter_CreateDeployment_Call) RunAndReturn(run func(context.Context, v1.Deployment) error) *IKubernetesAdapter_CreateDeployment_Call {
	_c.Call.Return(run)
	return _c
}

// ListPods provides a mock function with given fields: ctx, appName
func (_m *IKubernetesAdapter) ListPods(ctx context.Context, appName string) ([]corev1.Pod, error) {
	ret := _m.Called(ctx, appName)

	if len(ret) == 0 {
		panic("no return value specified for ListPods")
	}

	var r0 []corev1.Pod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]corev1.Pod, error)); ok {
		return rf(ctx, appName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []corev1.Pod); ok {
		r0 = rf(ctx, appName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]corev1.Pod)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, appName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IKubernetesAdapter_ListPods_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPods'
type IKubernetesAdapter_ListPods_Call struct {
	*mock.Call
}

// ListPods is a helper method to define mock.On call
//   - ctx context.Context
//   - appName string
func (_e *IKubernetesAdapter_Expecter) ListPods(ctx interface{}, appName interface{}) *IKubernetesAdapter_ListPods_Call {
	return &IKubernetesAdapter_ListPods_Call{Call: _e.mock.On("ListPods", ctx, appName)}
}

func (_c *IKubernetesAdapter_ListPods_Call) Run(run func(ctx context.Context, appName string)) *IKubernetesAdapter_ListPods_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *IKubernetesAdapter_ListPods_Call) Return(_a0 []corev1.Pod, _a1 error) *IKubernetesAdapter_ListPods_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IKubernetesAdapter_ListPods_Call) RunAndReturn(run func(context.Context, string) ([]corev1.Pod, error)) *IKubernetesAdapter_ListPods_Call {
	_c.Call.Return(run)
	return _c
}

// RestartDeployment provides a mock function with given fields: ctx, deploymentName
func (_m *IKubernetesAdapter) RestartDeployment(ctx context.Context, deploymentName string) error {
	ret := _m.Called(ctx, deploymentName)

	if len(ret) == 0 {
		panic("no return value specified for RestartDeployment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, deploymentName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IKubernetesAdapter_RestartDeployment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RestartDeployment'
type IKubernetesAdapter_RestartDeployment_Call struct {
	*mock.Call
}

// RestartDeployment is a helper method to define mock.On call
//   - ctx context.Context
//   - deploymentName string
func (_e *IKubernetesAdapter_Expecter) RestartDeployment(ctx interface{}, deploymentName interface{}) *IKubernetesAdapter_RestartDeployment_Call {
	return &IKubernetesAdapter_RestartDeployment_Call{Call: _e.mock.On("RestartDeployment", ctx, deploymentName)}
}

func (_c *IKubernetesAdapter_RestartDeployment_Call) Run(run func(ctx context.Context, deploymentName string)) *IKubernetesAdapter_RestartDeployment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *IKubernetesAdapter_RestartDeployment_Call) Return(_a0 error) *IKubernetesAdapter_RestartDeployment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IKubernetesAdapter_RestartDeployment_Call) RunAndReturn(run func(context.Context, string) error) *IKubernetesAdapter_RestartDeployment_Call {
	_c.Call.Return(run)
	return _c
}

// SetReplicas provides a mock function with given fields: ctx, deploymentName, replica
func (_m *IKubernetesAdapter) SetReplicas(ctx context.Context, deploymentName string, replica int32) error {
	ret := _m.Called(ctx, deploymentName, replica)

	if len(ret) == 0 {
		panic("no return value specified for SetReplicas")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int32) error); ok {
		r0 = rf(ctx, deploymentName, replica)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IKubernetesAdapter_SetReplicas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetReplicas'
type IKubernetesAdapter_SetReplicas_Call struct {
	*mock.Call
}

// SetReplicas is a helper method to define mock.On call
//   - ctx context.Context
//   - deploymentName string
//   - replica int32
func (_e *IKubernetesAdapter_Expecter) SetReplicas(ctx interface{}, deploymentName interface{}, replica interface{}) *IKubernetesAdapter_SetReplicas_Call {
	return &IKubernetesAdapter_SetReplicas_Call{Call: _e.mock.On("SetReplicas", ctx, deploymentName, replica)}
}

func (_c *IKubernetesAdapter_SetReplicas_Call) Run(run func(ctx context.Context, deploymentName string, replica int32)) *IKubernetesAdapter_SetReplicas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int32))
	})
	return _c
}

func (_c *IKubernetesAdapter_SetReplicas_Call) Return(_a0 error) *IKubernetesAdapter_SetReplicas_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IKubernetesAdapter_SetReplicas_Call) RunAndReturn(run func(context.Context, string, int32) error) *IKubernetesAdapter_SetReplicas_Call {
	_c.Call.Return(run)
	return _c
}

// NewIKubernetesAdapter creates a new instance of IKubernetesAdapter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIKubernetesAdapter(t interface {
	mock.TestingT
	Cleanup(func())
}) *IKubernetesAdapter {
	mock := &IKubernetesAdapter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}