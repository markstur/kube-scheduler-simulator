// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubernetes-sigs/kube-scheduler-simulator/export (interfaces: NodeService)

// Package mock_export is a generated GoMock package.
package mock_export

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/client-go/applyconfigurations/core/v1"
)

// MockNodeService is a mock of NodeService interface.
type MockNodeService struct {
	ctrl     *gomock.Controller
	recorder *MockNodeServiceMockRecorder
}

// MockNodeServiceMockRecorder is the mock recorder for MockNodeService.
type MockNodeServiceMockRecorder struct {
	mock *MockNodeService
}

// NewMockNodeService creates a new mock instance.
func NewMockNodeService(ctrl *gomock.Controller) *MockNodeService {
	mock := &MockNodeService{ctrl: ctrl}
	mock.recorder = &MockNodeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeService) EXPECT() *MockNodeServiceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockNodeService) Apply(arg0 context.Context, arg1 *v10.NodeApplyConfiguration) (*v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1)
	ret0, _ := ret[0].(*v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockNodeServiceMockRecorder) Apply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockNodeService)(nil).Apply), arg0, arg1)
}

// List mocks base method.
func (m *MockNodeService) List(arg0 context.Context) (*v1.NodeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1.NodeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockNodeServiceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNodeService)(nil).List), arg0)
}
