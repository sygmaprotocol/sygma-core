// Code generated by MockGen. DO NOT EDIT.
// Source: chains/evm/transactor/signAndSend/signAndSend.go
//
// Generated by this command:
//
//	mockgen -source=chains/evm/transactor/signAndSend/signAndSend.go -destination=./mock/signAndSend.go -package mock
//
// Package mock is a generated GoMock package.
package mock

import (
	big "math/big"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockGasPricer is a mock of GasPricer interface.
type MockGasPricer struct {
	ctrl     *gomock.Controller
	recorder *MockGasPricerMockRecorder
}

// MockGasPricerMockRecorder is the mock recorder for MockGasPricer.
type MockGasPricerMockRecorder struct {
	mock *MockGasPricer
}

// NewMockGasPricer creates a new mock instance.
func NewMockGasPricer(ctrl *gomock.Controller) *MockGasPricer {
	mock := &MockGasPricer{ctrl: ctrl}
	mock.recorder = &MockGasPricerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGasPricer) EXPECT() *MockGasPricerMockRecorder {
	return m.recorder
}

// GasPrice mocks base method.
func (m *MockGasPricer) GasPrice(priority *uint8) ([]*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GasPrice", priority)
	ret0, _ := ret[0].([]*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GasPrice indicates an expected call of GasPrice.
func (mr *MockGasPricerMockRecorder) GasPrice(priority any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GasPrice", reflect.TypeOf((*MockGasPricer)(nil).GasPrice), priority)
}
