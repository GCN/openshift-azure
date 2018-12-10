// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/openshift-azure/pkg/util/kubeclient (interfaces: Kubeclient)

// Package mock_kubeclient is a generated GoMock package.
package mock_kubeclient

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	api "github.com/openshift/openshift-azure/pkg/api"
	kubeclient "github.com/openshift/openshift-azure/pkg/util/kubeclient"
)

// MockKubeclient is a mock of Kubeclient interface
type MockKubeclient struct {
	ctrl     *gomock.Controller
	recorder *MockKubeclientMockRecorder
}

// MockKubeclientMockRecorder is the mock recorder for MockKubeclient
type MockKubeclientMockRecorder struct {
	mock *MockKubeclient
}

// NewMockKubeclient creates a new mock instance
func NewMockKubeclient(ctrl *gomock.Controller) *MockKubeclient {
	mock := &MockKubeclient{ctrl: ctrl}
	mock.recorder = &MockKubeclientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubeclient) EXPECT() *MockKubeclientMockRecorder {
	return m.recorder
}

// Drain mocks base method
func (m *MockKubeclient) Drain(arg0 context.Context, arg1 api.AgentPoolProfileRole, arg2 kubeclient.ComputerName) error {
	ret := m.ctrl.Call(m, "Drain", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Drain indicates an expected call of Drain
func (mr *MockKubeclientMockRecorder) Drain(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drain", reflect.TypeOf((*MockKubeclient)(nil).Drain), arg0, arg1, arg2)
}

// MasterIsReady mocks base method
func (m *MockKubeclient) MasterIsReady(arg0 kubeclient.ComputerName) (bool, error) {
	ret := m.ctrl.Call(m, "MasterIsReady", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MasterIsReady indicates an expected call of MasterIsReady
func (mr *MockKubeclientMockRecorder) MasterIsReady(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MasterIsReady", reflect.TypeOf((*MockKubeclient)(nil).MasterIsReady), arg0)
}

// WaitForInfraServices mocks base method
func (m *MockKubeclient) WaitForInfraServices(arg0 context.Context) *api.PluginError {
	ret := m.ctrl.Call(m, "WaitForInfraServices", arg0)
	ret0, _ := ret[0].(*api.PluginError)
	return ret0
}

// WaitForInfraServices indicates an expected call of WaitForInfraServices
func (mr *MockKubeclientMockRecorder) WaitForInfraServices(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForInfraServices", reflect.TypeOf((*MockKubeclient)(nil).WaitForInfraServices), arg0)
}

// WaitForReady mocks base method
func (m *MockKubeclient) WaitForReady(arg0 context.Context, arg1 api.AgentPoolProfileRole, arg2 kubeclient.ComputerName) error {
	ret := m.ctrl.Call(m, "WaitForReady", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForReady indicates an expected call of WaitForReady
func (mr *MockKubeclientMockRecorder) WaitForReady(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForReady", reflect.TypeOf((*MockKubeclient)(nil).WaitForReady), arg0, arg1, arg2)
}
