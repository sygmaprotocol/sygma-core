// Code generated by MockGen. DO NOT EDIT.
// Source: ./chains/evm/client/client.go
//
// Generated by this command:
//
//	mockgen -destination=./mock/client.go -source=./chains/evm/client/client.go -package mock
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	big "math/big"
	reflect "reflect"

	client "github.com/sygmaprotocol/sygma-core/chains/evm/client"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "go.uber.org/mock/gomock"
)

// MockContractCaller is a mock of ContractCaller interface.
type MockContractCaller struct {
	ctrl     *gomock.Controller
	recorder *MockContractCallerMockRecorder
}

// MockContractCallerMockRecorder is the mock recorder for MockContractCaller.
type MockContractCallerMockRecorder struct {
	mock *MockContractCaller
}

// NewMockContractCaller creates a new mock instance.
func NewMockContractCaller(ctrl *gomock.Controller) *MockContractCaller {
	mock := &MockContractCaller{ctrl: ctrl}
	mock.recorder = &MockContractCallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractCaller) EXPECT() *MockContractCallerMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockContractCaller) CallContract(ctx context.Context, callArgs map[string]any, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", ctx, callArgs, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockContractCallerMockRecorder) CallContract(ctx, callArgs, blockNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockContractCaller)(nil).CallContract), ctx, callArgs, blockNumber)
}

// CodeAt mocks base method.
func (m *MockContractCaller) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeAt", ctx, contract, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CodeAt indicates an expected call of CodeAt.
func (mr *MockContractCallerMockRecorder) CodeAt(ctx, contract, blockNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeAt", reflect.TypeOf((*MockContractCaller)(nil).CodeAt), ctx, contract, blockNumber)
}

// MockTransactionDispatcher is a mock of TransactionDispatcher interface.
type MockTransactionDispatcher struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionDispatcherMockRecorder
}

// MockTransactionDispatcherMockRecorder is the mock recorder for MockTransactionDispatcher.
type MockTransactionDispatcherMockRecorder struct {
	mock *MockTransactionDispatcher
}

// NewMockTransactionDispatcher creates a new mock instance.
func NewMockTransactionDispatcher(ctrl *gomock.Controller) *MockTransactionDispatcher {
	mock := &MockTransactionDispatcher{ctrl: ctrl}
	mock.recorder = &MockTransactionDispatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionDispatcher) EXPECT() *MockTransactionDispatcherMockRecorder {
	return m.recorder
}

// From mocks base method.
func (m *MockTransactionDispatcher) From() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "From")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// From indicates an expected call of From.
func (mr *MockTransactionDispatcherMockRecorder) From() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "From", reflect.TypeOf((*MockTransactionDispatcher)(nil).From))
}

// GetTransactionByHash mocks base method.
func (m *MockTransactionDispatcher) GetTransactionByHash(h common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByHash", h)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTransactionByHash indicates an expected call of GetTransactionByHash.
func (mr *MockTransactionDispatcherMockRecorder) GetTransactionByHash(h any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByHash", reflect.TypeOf((*MockTransactionDispatcher)(nil).GetTransactionByHash), h)
}

// LockNonce mocks base method.
func (m *MockTransactionDispatcher) LockNonce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LockNonce")
}

// LockNonce indicates an expected call of LockNonce.
func (mr *MockTransactionDispatcherMockRecorder) LockNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockNonce", reflect.TypeOf((*MockTransactionDispatcher)(nil).LockNonce))
}

// SignAndSendTransaction mocks base method.
func (m *MockTransactionDispatcher) SignAndSendTransaction(ctx context.Context, tx client.CommonTransaction) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignAndSendTransaction", ctx, tx)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignAndSendTransaction indicates an expected call of SignAndSendTransaction.
func (mr *MockTransactionDispatcherMockRecorder) SignAndSendTransaction(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignAndSendTransaction", reflect.TypeOf((*MockTransactionDispatcher)(nil).SignAndSendTransaction), ctx, tx)
}

// TransactionReceipt mocks base method.
func (m *MockTransactionDispatcher) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", ctx, txHash)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt.
func (mr *MockTransactionDispatcherMockRecorder) TransactionReceipt(ctx, txHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockTransactionDispatcher)(nil).TransactionReceipt), ctx, txHash)
}

// UnlockNonce mocks base method.
func (m *MockTransactionDispatcher) UnlockNonce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnlockNonce")
}

// UnlockNonce indicates an expected call of UnlockNonce.
func (mr *MockTransactionDispatcherMockRecorder) UnlockNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockNonce", reflect.TypeOf((*MockTransactionDispatcher)(nil).UnlockNonce))
}

// UnsafeIncreaseNonce mocks base method.
func (m *MockTransactionDispatcher) UnsafeIncreaseNonce() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsafeIncreaseNonce")
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsafeIncreaseNonce indicates an expected call of UnsafeIncreaseNonce.
func (mr *MockTransactionDispatcherMockRecorder) UnsafeIncreaseNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsafeIncreaseNonce", reflect.TypeOf((*MockTransactionDispatcher)(nil).UnsafeIncreaseNonce))
}

// UnsafeNonce mocks base method.
func (m *MockTransactionDispatcher) UnsafeNonce() (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsafeNonce")
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnsafeNonce indicates an expected call of UnsafeNonce.
func (mr *MockTransactionDispatcherMockRecorder) UnsafeNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsafeNonce", reflect.TypeOf((*MockTransactionDispatcher)(nil).UnsafeNonce))
}

// WaitAndReturnTxReceipt mocks base method.
func (m *MockTransactionDispatcher) WaitAndReturnTxReceipt(h common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitAndReturnTxReceipt", h)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitAndReturnTxReceipt indicates an expected call of WaitAndReturnTxReceipt.
func (mr *MockTransactionDispatcherMockRecorder) WaitAndReturnTxReceipt(h any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitAndReturnTxReceipt", reflect.TypeOf((*MockTransactionDispatcher)(nil).WaitAndReturnTxReceipt), h)
}

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CallContract mocks base method.
func (m *MockClient) CallContract(ctx context.Context, callArgs map[string]any, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CallContract", ctx, callArgs, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CallContract indicates an expected call of CallContract.
func (mr *MockClientMockRecorder) CallContract(ctx, callArgs, blockNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CallContract", reflect.TypeOf((*MockClient)(nil).CallContract), ctx, callArgs, blockNumber)
}

// CodeAt mocks base method.
func (m *MockClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeAt", ctx, contract, blockNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CodeAt indicates an expected call of CodeAt.
func (mr *MockClientMockRecorder) CodeAt(ctx, contract, blockNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeAt", reflect.TypeOf((*MockClient)(nil).CodeAt), ctx, contract, blockNumber)
}

// From mocks base method.
func (m *MockClient) From() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "From")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// From indicates an expected call of From.
func (mr *MockClientMockRecorder) From() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "From", reflect.TypeOf((*MockClient)(nil).From))
}

// GetTransactionByHash mocks base method.
func (m *MockClient) GetTransactionByHash(h common.Hash) (*types.Transaction, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionByHash", h)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTransactionByHash indicates an expected call of GetTransactionByHash.
func (mr *MockClientMockRecorder) GetTransactionByHash(h any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionByHash", reflect.TypeOf((*MockClient)(nil).GetTransactionByHash), h)
}

// LockNonce mocks base method.
func (m *MockClient) LockNonce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LockNonce")
}

// LockNonce indicates an expected call of LockNonce.
func (mr *MockClientMockRecorder) LockNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockNonce", reflect.TypeOf((*MockClient)(nil).LockNonce))
}

// SignAndSendTransaction mocks base method.
func (m *MockClient) SignAndSendTransaction(ctx context.Context, tx client.CommonTransaction) (common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignAndSendTransaction", ctx, tx)
	ret0, _ := ret[0].(common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignAndSendTransaction indicates an expected call of SignAndSendTransaction.
func (mr *MockClientMockRecorder) SignAndSendTransaction(ctx, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignAndSendTransaction", reflect.TypeOf((*MockClient)(nil).SignAndSendTransaction), ctx, tx)
}

// TransactionReceipt mocks base method.
func (m *MockClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", ctx, txHash)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt.
func (mr *MockClientMockRecorder) TransactionReceipt(ctx, txHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockClient)(nil).TransactionReceipt), ctx, txHash)
}

// UnlockNonce mocks base method.
func (m *MockClient) UnlockNonce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UnlockNonce")
}

// UnlockNonce indicates an expected call of UnlockNonce.
func (mr *MockClientMockRecorder) UnlockNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockNonce", reflect.TypeOf((*MockClient)(nil).UnlockNonce))
}

// UnsafeIncreaseNonce mocks base method.
func (m *MockClient) UnsafeIncreaseNonce() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsafeIncreaseNonce")
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsafeIncreaseNonce indicates an expected call of UnsafeIncreaseNonce.
func (mr *MockClientMockRecorder) UnsafeIncreaseNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsafeIncreaseNonce", reflect.TypeOf((*MockClient)(nil).UnsafeIncreaseNonce))
}

// UnsafeNonce mocks base method.
func (m *MockClient) UnsafeNonce() (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsafeNonce")
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnsafeNonce indicates an expected call of UnsafeNonce.
func (mr *MockClientMockRecorder) UnsafeNonce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsafeNonce", reflect.TypeOf((*MockClient)(nil).UnsafeNonce))
}

// WaitAndReturnTxReceipt mocks base method.
func (m *MockClient) WaitAndReturnTxReceipt(h common.Hash) (*types.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitAndReturnTxReceipt", h)
	ret0, _ := ret[0].(*types.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitAndReturnTxReceipt indicates an expected call of WaitAndReturnTxReceipt.
func (mr *MockClientMockRecorder) WaitAndReturnTxReceipt(h any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitAndReturnTxReceipt", reflect.TypeOf((*MockClient)(nil).WaitAndReturnTxReceipt), h)
}

// MockSigner is a mock of Signer interface.
type MockSigner struct {
	ctrl     *gomock.Controller
	recorder *MockSignerMockRecorder
}

// MockSignerMockRecorder is the mock recorder for MockSigner.
type MockSignerMockRecorder struct {
	mock *MockSigner
}

// NewMockSigner creates a new mock instance.
func NewMockSigner(ctrl *gomock.Controller) *MockSigner {
	mock := &MockSigner{ctrl: ctrl}
	mock.recorder = &MockSignerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSigner) EXPECT() *MockSignerMockRecorder {
	return m.recorder
}

// CommonAddress mocks base method.
func (m *MockSigner) CommonAddress() common.Address {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommonAddress")
	ret0, _ := ret[0].(common.Address)
	return ret0
}

// CommonAddress indicates an expected call of CommonAddress.
func (mr *MockSignerMockRecorder) CommonAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommonAddress", reflect.TypeOf((*MockSigner)(nil).CommonAddress))
}

// Sign mocks base method.
func (m *MockSigner) Sign(digestHash []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", digestHash)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockSignerMockRecorder) Sign(digestHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockSigner)(nil).Sign), digestHash)
}

// MockCommonTransaction is a mock of CommonTransaction interface.
type MockCommonTransaction struct {
	ctrl     *gomock.Controller
	recorder *MockCommonTransactionMockRecorder
}

// MockCommonTransactionMockRecorder is the mock recorder for MockCommonTransaction.
type MockCommonTransactionMockRecorder struct {
	mock *MockCommonTransaction
}

// NewMockCommonTransaction creates a new mock instance.
func NewMockCommonTransaction(ctrl *gomock.Controller) *MockCommonTransaction {
	mock := &MockCommonTransaction{ctrl: ctrl}
	mock.recorder = &MockCommonTransactionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommonTransaction) EXPECT() *MockCommonTransactionMockRecorder {
	return m.recorder
}

// Hash mocks base method.
func (m *MockCommonTransaction) Hash() common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash")
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// Hash indicates an expected call of Hash.
func (mr *MockCommonTransactionMockRecorder) Hash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockCommonTransaction)(nil).Hash))
}

// RawWithSignature mocks base method.
func (m *MockCommonTransaction) RawWithSignature(signer client.Signer, domainID *big.Int) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawWithSignature", signer, domainID)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RawWithSignature indicates an expected call of RawWithSignature.
func (mr *MockCommonTransactionMockRecorder) RawWithSignature(signer, domainID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawWithSignature", reflect.TypeOf((*MockCommonTransaction)(nil).RawWithSignature), signer, domainID)
}
