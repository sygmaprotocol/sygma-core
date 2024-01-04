// Code generated by MockGen. DO NOT EDIT.
// Source: ./relayer/relayer.go
//
// Generated by this command:
//
//	mockgen -destination=./mock/relayer.go -source=./relayer/relayer.go -package mock
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	message "github.com/sygmaprotocol/sygma-core/relayer/message"
	proposal "github.com/sygmaprotocol/sygma-core/relayer/proposal"
	gomock "go.uber.org/mock/gomock"
)

// MockRelayedChain is a mock of RelayedChain interface.
type MockRelayedChain struct {
	ctrl     *gomock.Controller
	recorder *MockRelayedChainMockRecorder
}

// MockRelayedChainMockRecorder is the mock recorder for MockRelayedChain.
type MockRelayedChainMockRecorder struct {
	mock *MockRelayedChain
}

// NewMockRelayedChain creates a new mock instance.
func NewMockRelayedChain(ctrl *gomock.Controller) *MockRelayedChain {
	mock := &MockRelayedChain{ctrl: ctrl}
	mock.recorder = &MockRelayedChainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRelayedChain) EXPECT() *MockRelayedChainMockRecorder {
	return m.recorder
}

// DomainID mocks base method.
func (m *MockRelayedChain) DomainID() uint8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DomainID")
	ret0, _ := ret[0].(uint8)
	return ret0
}

// DomainID indicates an expected call of DomainID.
func (mr *MockRelayedChainMockRecorder) DomainID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainID", reflect.TypeOf((*MockRelayedChain)(nil).DomainID))
}

// PollEvents mocks base method.
func (m *MockRelayedChain) PollEvents(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PollEvents", ctx)
}

// PollEvents indicates an expected call of PollEvents.
func (mr *MockRelayedChainMockRecorder) PollEvents(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollEvents", reflect.TypeOf((*MockRelayedChain)(nil).PollEvents), ctx)
}

// ReceiveMessage mocks base method.
func (m_2 *MockRelayedChain) ReceiveMessage(m *message.Message) (*proposal.Proposal, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "ReceiveMessage", m)
	ret0, _ := ret[0].(*proposal.Proposal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage.
func (mr *MockRelayedChainMockRecorder) ReceiveMessage(m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockRelayedChain)(nil).ReceiveMessage), m)
}

// Write mocks base method.
func (m *MockRelayedChain) Write(proposals []*proposal.Proposal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", proposals)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockRelayedChainMockRecorder) Write(proposals any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockRelayedChain)(nil).Write), proposals)
}
