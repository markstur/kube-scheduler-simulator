// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubernetes-sigs/kube-scheduler-simulator/export (interfaces: SchedulerService)

// Package mock_export is a generated GoMock package.
package mock_export

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta3 "k8s.io/kube-scheduler/config/v1beta3"
)

// MockSchedulerService is a mock of SchedulerService interface.
type MockSchedulerService struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerServiceMockRecorder
}

// MockSchedulerServiceMockRecorder is the mock recorder for MockSchedulerService.
type MockSchedulerServiceMockRecorder struct {
	mock *MockSchedulerService
}

// NewMockSchedulerService creates a new mock instance.
func NewMockSchedulerService(ctrl *gomock.Controller) *MockSchedulerService {
	mock := &MockSchedulerService{ctrl: ctrl}
	mock.recorder = &MockSchedulerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchedulerService) EXPECT() *MockSchedulerServiceMockRecorder {
	return m.recorder
}

// GetSchedulerConfig mocks base method.
func (m *MockSchedulerService) GetSchedulerConfig() *v1beta3.KubeSchedulerConfiguration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedulerConfig")
	ret0, _ := ret[0].(*v1beta3.KubeSchedulerConfiguration)
	return ret0
}

// GetSchedulerConfig indicates an expected call of GetSchedulerConfig.
func (mr *MockSchedulerServiceMockRecorder) GetSchedulerConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedulerConfig", reflect.TypeOf((*MockSchedulerService)(nil).GetSchedulerConfig))
}

// RestartScheduler mocks base method.
func (m *MockSchedulerService) RestartScheduler(arg0 *v1beta3.KubeSchedulerConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestartScheduler", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestartScheduler indicates an expected call of RestartScheduler.
func (mr *MockSchedulerServiceMockRecorder) RestartScheduler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestartScheduler", reflect.TypeOf((*MockSchedulerService)(nil).RestartScheduler), arg0)
}
