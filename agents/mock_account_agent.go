// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/caser789/jw-session/agents (interfaces: AccountAgent)

// Package agents is a generated GoMock package.
package agents

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccountAgent is a mock of AccountAgent interface.
type MockAccountAgent struct {
	ctrl     *gomock.Controller
	recorder *MockAccountAgentMockRecorder
}

// MockAccountAgentMockRecorder is the mock recorder for MockAccountAgent.
type MockAccountAgentMockRecorder struct {
	mock *MockAccountAgent
}

// NewMockAccountAgent creates a new mock instance.
func NewMockAccountAgent(ctrl *gomock.Controller) *MockAccountAgent {
	mock := &MockAccountAgent{ctrl: ctrl}
	mock.recorder = &MockAccountAgentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountAgent) EXPECT() *MockAccountAgentMockRecorder {
	return m.recorder
}

// VerifyToken mocks base method.
func (m *MockAccountAgent) VerifyToken(arg0 string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockAccountAgentMockRecorder) VerifyToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockAccountAgent)(nil).VerifyToken), arg0)
}
