// Code generated by MockGen. DO NOT EDIT.
// Source: network.go

// Package mock_networks is a generated GoMock package.
package mock_networks

import (
	big "math/big"
	reflect "reflect"

	bindings "github.com/ambrosus/ambrosus-bridge/relay/internal/bindings"
	interfaces "github.com/ambrosus/ambrosus-bridge/relay/internal/bindings/interfaces"
	networks "github.com/ambrosus/ambrosus-bridge/relay/internal/networks"
	ethash "github.com/ambrosus/ambrosus-bridge/relay/pkg/ethash"
	ethclients "github.com/ambrosus/ambrosus-bridge/relay/pkg/ethclients"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
	zerolog "github.com/rs/zerolog"
)

// MockBridge is a mock of Bridge interface.
type MockBridge struct {
	ctrl     *gomock.Controller
	recorder *MockBridgeMockRecorder
}

// MockBridgeMockRecorder is the mock recorder for MockBridge.
type MockBridgeMockRecorder struct {
	mock *MockBridge
}

// NewMockBridge creates a new mock instance.
func NewMockBridge(ctrl *gomock.Controller) *MockBridge {
	mock := &MockBridge{ctrl: ctrl}
	mock.recorder = &MockBridgeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBridge) EXPECT() *MockBridgeMockRecorder {
	return m.recorder
}

// EnsureContractUnpaused mocks base method.
func (m *MockBridge) EnsureContractUnpaused() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnsureContractUnpaused")
}

// EnsureContractUnpaused indicates an expected call of EnsureContractUnpaused.
func (mr *MockBridgeMockRecorder) EnsureContractUnpaused() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureContractUnpaused", reflect.TypeOf((*MockBridge)(nil).EnsureContractUnpaused))
}

// GetClient mocks base method.
func (m *MockBridge) GetClient() ethclients.ClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(ethclients.ClientInterface)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockBridgeMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockBridge)(nil).GetClient))
}

// GetContract mocks base method.
func (m *MockBridge) GetContract() interfaces.BridgeContract {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract")
	ret0, _ := ret[0].(interfaces.BridgeContract)
	return ret0
}

// GetContract indicates an expected call of GetContract.
func (mr *MockBridgeMockRecorder) GetContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockBridge)(nil).GetContract))
}

// GetLogger mocks base method.
func (m *MockBridge) GetLogger() *zerolog.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(*zerolog.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockBridgeMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockBridge)(nil).GetLogger))
}

// GetName mocks base method.
func (m *MockBridge) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockBridgeMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockBridge)(nil).GetName))
}

// GetWsClient mocks base method.
func (m *MockBridge) GetWsClient() ethclients.ClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWsClient")
	ret0, _ := ret[0].(ethclients.ClientInterface)
	return ret0
}

// GetWsClient indicates an expected call of GetWsClient.
func (mr *MockBridgeMockRecorder) GetWsClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWsClient", reflect.TypeOf((*MockBridge)(nil).GetWsClient))
}

// GetWsContract mocks base method.
func (m *MockBridge) GetWsContract() interfaces.BridgeContract {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWsContract")
	ret0, _ := ret[0].(interfaces.BridgeContract)
	return ret0
}

// GetWsContract indicates an expected call of GetWsContract.
func (mr *MockBridgeMockRecorder) GetWsContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWsContract", reflect.TypeOf((*MockBridge)(nil).GetWsContract))
}

// ProcessTx mocks base method.
func (m *MockBridge) ProcessTx(methodName string, txCallback networks.ContractCallFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTx", methodName, txCallback)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessTx indicates an expected call of ProcessTx.
func (mr *MockBridgeMockRecorder) ProcessTx(methodName, txCallback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTx", reflect.TypeOf((*MockBridge)(nil).ProcessTx), methodName, txCallback)
}

// ShouldHavePk mocks base method.
func (m *MockBridge) ShouldHavePk() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShouldHavePk")
}

// ShouldHavePk indicates an expected call of ShouldHavePk.
func (mr *MockBridgeMockRecorder) ShouldHavePk() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldHavePk", reflect.TypeOf((*MockBridge)(nil).ShouldHavePk))
}

// MockSubmitter is a mock of Submitter interface.
type MockSubmitter struct {
	ctrl     *gomock.Controller
	recorder *MockSubmitterMockRecorder
}

// MockSubmitterMockRecorder is the mock recorder for MockSubmitter.
type MockSubmitterMockRecorder struct {
	mock *MockSubmitter
}

// NewMockSubmitter creates a new mock instance.
func NewMockSubmitter(ctrl *gomock.Controller) *MockSubmitter {
	mock := &MockSubmitter{ctrl: ctrl}
	mock.recorder = &MockSubmitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubmitter) EXPECT() *MockSubmitterMockRecorder {
	return m.recorder
}

// EnsureContractUnpaused mocks base method.
func (m *MockSubmitter) EnsureContractUnpaused() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnsureContractUnpaused")
}

// EnsureContractUnpaused indicates an expected call of EnsureContractUnpaused.
func (mr *MockSubmitterMockRecorder) EnsureContractUnpaused() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureContractUnpaused", reflect.TypeOf((*MockSubmitter)(nil).EnsureContractUnpaused))
}

// GetClient mocks base method.
func (m *MockSubmitter) GetClient() ethclients.ClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(ethclients.ClientInterface)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockSubmitterMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockSubmitter)(nil).GetClient))
}

// GetContract mocks base method.
func (m *MockSubmitter) GetContract() interfaces.BridgeContract {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract")
	ret0, _ := ret[0].(interfaces.BridgeContract)
	return ret0
}

// GetContract indicates an expected call of GetContract.
func (mr *MockSubmitterMockRecorder) GetContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockSubmitter)(nil).GetContract))
}

// GetEventById mocks base method.
func (m *MockSubmitter) GetEventById(eventId *big.Int) (*bindings.BridgeTransfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventById", eventId)
	ret0, _ := ret[0].(*bindings.BridgeTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventById indicates an expected call of GetEventById.
func (mr *MockSubmitterMockRecorder) GetEventById(eventId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventById", reflect.TypeOf((*MockSubmitter)(nil).GetEventById), eventId)
}

// GetLogger mocks base method.
func (m *MockSubmitter) GetLogger() *zerolog.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(*zerolog.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockSubmitterMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockSubmitter)(nil).GetLogger))
}

// GetName mocks base method.
func (m *MockSubmitter) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockSubmitterMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockSubmitter)(nil).GetName))
}

// GetWsClient mocks base method.
func (m *MockSubmitter) GetWsClient() ethclients.ClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWsClient")
	ret0, _ := ret[0].(ethclients.ClientInterface)
	return ret0
}

// GetWsClient indicates an expected call of GetWsClient.
func (mr *MockSubmitterMockRecorder) GetWsClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWsClient", reflect.TypeOf((*MockSubmitter)(nil).GetWsClient))
}

// GetWsContract mocks base method.
func (m *MockSubmitter) GetWsContract() interfaces.BridgeContract {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWsContract")
	ret0, _ := ret[0].(interfaces.BridgeContract)
	return ret0
}

// GetWsContract indicates an expected call of GetWsContract.
func (mr *MockSubmitterMockRecorder) GetWsContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWsContract", reflect.TypeOf((*MockSubmitter)(nil).GetWsContract))
}

// ProcessTx mocks base method.
func (m *MockSubmitter) ProcessTx(methodName string, txCallback networks.ContractCallFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTx", methodName, txCallback)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessTx indicates an expected call of ProcessTx.
func (mr *MockSubmitterMockRecorder) ProcessTx(methodName, txCallback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTx", reflect.TypeOf((*MockSubmitter)(nil).ProcessTx), methodName, txCallback)
}

// SendEvent mocks base method.
func (m *MockSubmitter) SendEvent(event *bindings.BridgeTransfer, safetyBlocks uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEvent", event, safetyBlocks)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEvent indicates an expected call of SendEvent.
func (mr *MockSubmitterMockRecorder) SendEvent(event, safetyBlocks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEvent", reflect.TypeOf((*MockSubmitter)(nil).SendEvent), event, safetyBlocks)
}

// ShouldHavePk mocks base method.
func (m *MockSubmitter) ShouldHavePk() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShouldHavePk")
}

// ShouldHavePk indicates an expected call of ShouldHavePk.
func (mr *MockSubmitterMockRecorder) ShouldHavePk() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldHavePk", reflect.TypeOf((*MockSubmitter)(nil).ShouldHavePk))
}

// MockReceiver is a mock of Receiver interface.
type MockReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockReceiverMockRecorder
}

// MockReceiverMockRecorder is the mock recorder for MockReceiver.
type MockReceiverMockRecorder struct {
	mock *MockReceiver
}

// NewMockReceiver creates a new mock instance.
func NewMockReceiver(ctrl *gomock.Controller) *MockReceiver {
	mock := &MockReceiver{ctrl: ctrl}
	mock.recorder = &MockReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceiver) EXPECT() *MockReceiverMockRecorder {
	return m.recorder
}

// EnsureContractUnpaused mocks base method.
func (m *MockReceiver) EnsureContractUnpaused() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnsureContractUnpaused")
}

// EnsureContractUnpaused indicates an expected call of EnsureContractUnpaused.
func (mr *MockReceiverMockRecorder) EnsureContractUnpaused() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureContractUnpaused", reflect.TypeOf((*MockReceiver)(nil).EnsureContractUnpaused))
}

// GetLastReceivedEventId mocks base method.
func (m *MockReceiver) GetLastReceivedEventId() (*big.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastReceivedEventId")
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastReceivedEventId indicates an expected call of GetLastReceivedEventId.
func (mr *MockReceiverMockRecorder) GetLastReceivedEventId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastReceivedEventId", reflect.TypeOf((*MockReceiver)(nil).GetLastReceivedEventId))
}

// GetMinSafetyBlocksNum mocks base method.
func (m *MockReceiver) GetMinSafetyBlocksNum() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinSafetyBlocksNum")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMinSafetyBlocksNum indicates an expected call of GetMinSafetyBlocksNum.
func (mr *MockReceiverMockRecorder) GetMinSafetyBlocksNum() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinSafetyBlocksNum", reflect.TypeOf((*MockReceiver)(nil).GetMinSafetyBlocksNum))
}

// MockBridgeReceiveAura is a mock of BridgeReceiveAura interface.
type MockBridgeReceiveAura struct {
	ctrl     *gomock.Controller
	recorder *MockBridgeReceiveAuraMockRecorder
}

// MockBridgeReceiveAuraMockRecorder is the mock recorder for MockBridgeReceiveAura.
type MockBridgeReceiveAuraMockRecorder struct {
	mock *MockBridgeReceiveAura
}

// NewMockBridgeReceiveAura creates a new mock instance.
func NewMockBridgeReceiveAura(ctrl *gomock.Controller) *MockBridgeReceiveAura {
	mock := &MockBridgeReceiveAura{ctrl: ctrl}
	mock.recorder = &MockBridgeReceiveAuraMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBridgeReceiveAura) EXPECT() *MockBridgeReceiveAuraMockRecorder {
	return m.recorder
}

// EnsureContractUnpaused mocks base method.
func (m *MockBridgeReceiveAura) EnsureContractUnpaused() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnsureContractUnpaused")
}

// EnsureContractUnpaused indicates an expected call of EnsureContractUnpaused.
func (mr *MockBridgeReceiveAuraMockRecorder) EnsureContractUnpaused() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureContractUnpaused", reflect.TypeOf((*MockBridgeReceiveAura)(nil).EnsureContractUnpaused))
}

// GetClient mocks base method.
func (m *MockBridgeReceiveAura) GetClient() ethclients.ClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(ethclients.ClientInterface)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockBridgeReceiveAuraMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockBridgeReceiveAura)(nil).GetClient))
}

// GetContract mocks base method.
func (m *MockBridgeReceiveAura) GetContract() interfaces.BridgeContract {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract")
	ret0, _ := ret[0].(interfaces.BridgeContract)
	return ret0
}

// GetContract indicates an expected call of GetContract.
func (mr *MockBridgeReceiveAuraMockRecorder) GetContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockBridgeReceiveAura)(nil).GetContract))
}

// GetLastProcessedBlockHash mocks base method.
func (m *MockBridgeReceiveAura) GetLastProcessedBlockHash() (*common.Hash, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastProcessedBlockHash")
	ret0, _ := ret[0].(*common.Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastProcessedBlockHash indicates an expected call of GetLastProcessedBlockHash.
func (mr *MockBridgeReceiveAuraMockRecorder) GetLastProcessedBlockHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastProcessedBlockHash", reflect.TypeOf((*MockBridgeReceiveAura)(nil).GetLastProcessedBlockHash))
}

// GetLogger mocks base method.
func (m *MockBridgeReceiveAura) GetLogger() *zerolog.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(*zerolog.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockBridgeReceiveAuraMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockBridgeReceiveAura)(nil).GetLogger))
}

// GetMinSafetyBlocksValidators mocks base method.
func (m *MockBridgeReceiveAura) GetMinSafetyBlocksValidators() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinSafetyBlocksValidators")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMinSafetyBlocksValidators indicates an expected call of GetMinSafetyBlocksValidators.
func (mr *MockBridgeReceiveAuraMockRecorder) GetMinSafetyBlocksValidators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinSafetyBlocksValidators", reflect.TypeOf((*MockBridgeReceiveAura)(nil).GetMinSafetyBlocksValidators))
}

// GetName mocks base method.
func (m *MockBridgeReceiveAura) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockBridgeReceiveAuraMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockBridgeReceiveAura)(nil).GetName))
}

// GetValidatorSet mocks base method.
func (m *MockBridgeReceiveAura) GetValidatorSet() ([]common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorSet")
	ret0, _ := ret[0].([]common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorSet indicates an expected call of GetValidatorSet.
func (mr *MockBridgeReceiveAuraMockRecorder) GetValidatorSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorSet", reflect.TypeOf((*MockBridgeReceiveAura)(nil).GetValidatorSet))
}

// GetWsClient mocks base method.
func (m *MockBridgeReceiveAura) GetWsClient() ethclients.ClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWsClient")
	ret0, _ := ret[0].(ethclients.ClientInterface)
	return ret0
}

// GetWsClient indicates an expected call of GetWsClient.
func (mr *MockBridgeReceiveAuraMockRecorder) GetWsClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWsClient", reflect.TypeOf((*MockBridgeReceiveAura)(nil).GetWsClient))
}

// GetWsContract mocks base method.
func (m *MockBridgeReceiveAura) GetWsContract() interfaces.BridgeContract {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWsContract")
	ret0, _ := ret[0].(interfaces.BridgeContract)
	return ret0
}

// GetWsContract indicates an expected call of GetWsContract.
func (mr *MockBridgeReceiveAuraMockRecorder) GetWsContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWsContract", reflect.TypeOf((*MockBridgeReceiveAura)(nil).GetWsContract))
}

// ProcessTx mocks base method.
func (m *MockBridgeReceiveAura) ProcessTx(methodName string, txCallback networks.ContractCallFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTx", methodName, txCallback)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessTx indicates an expected call of ProcessTx.
func (mr *MockBridgeReceiveAuraMockRecorder) ProcessTx(methodName, txCallback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTx", reflect.TypeOf((*MockBridgeReceiveAura)(nil).ProcessTx), methodName, txCallback)
}

// ShouldHavePk mocks base method.
func (m *MockBridgeReceiveAura) ShouldHavePk() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShouldHavePk")
}

// ShouldHavePk indicates an expected call of ShouldHavePk.
func (mr *MockBridgeReceiveAuraMockRecorder) ShouldHavePk() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldHavePk", reflect.TypeOf((*MockBridgeReceiveAura)(nil).ShouldHavePk))
}

// SubmitTransferAura mocks base method.
func (m *MockBridgeReceiveAura) SubmitTransferAura(arg0 *bindings.CheckAuraAuraProof) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitTransferAura", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitTransferAura indicates an expected call of SubmitTransferAura.
func (mr *MockBridgeReceiveAuraMockRecorder) SubmitTransferAura(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransferAura", reflect.TypeOf((*MockBridgeReceiveAura)(nil).SubmitTransferAura), arg0)
}

// MockBridgeReceiveEthash is a mock of BridgeReceiveEthash interface.
type MockBridgeReceiveEthash struct {
	ctrl     *gomock.Controller
	recorder *MockBridgeReceiveEthashMockRecorder
}

// MockBridgeReceiveEthashMockRecorder is the mock recorder for MockBridgeReceiveEthash.
type MockBridgeReceiveEthashMockRecorder struct {
	mock *MockBridgeReceiveEthash
}

// NewMockBridgeReceiveEthash creates a new mock instance.
func NewMockBridgeReceiveEthash(ctrl *gomock.Controller) *MockBridgeReceiveEthash {
	mock := &MockBridgeReceiveEthash{ctrl: ctrl}
	mock.recorder = &MockBridgeReceiveEthashMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBridgeReceiveEthash) EXPECT() *MockBridgeReceiveEthashMockRecorder {
	return m.recorder
}

// EnsureContractUnpaused mocks base method.
func (m *MockBridgeReceiveEthash) EnsureContractUnpaused() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnsureContractUnpaused")
}

// EnsureContractUnpaused indicates an expected call of EnsureContractUnpaused.
func (mr *MockBridgeReceiveEthashMockRecorder) EnsureContractUnpaused() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureContractUnpaused", reflect.TypeOf((*MockBridgeReceiveEthash)(nil).EnsureContractUnpaused))
}

// GetClient mocks base method.
func (m *MockBridgeReceiveEthash) GetClient() ethclients.ClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(ethclients.ClientInterface)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockBridgeReceiveEthashMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockBridgeReceiveEthash)(nil).GetClient))
}

// GetContract mocks base method.
func (m *MockBridgeReceiveEthash) GetContract() interfaces.BridgeContract {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract")
	ret0, _ := ret[0].(interfaces.BridgeContract)
	return ret0
}

// GetContract indicates an expected call of GetContract.
func (mr *MockBridgeReceiveEthashMockRecorder) GetContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockBridgeReceiveEthash)(nil).GetContract))
}

// GetLogger mocks base method.
func (m *MockBridgeReceiveEthash) GetLogger() *zerolog.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(*zerolog.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockBridgeReceiveEthashMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockBridgeReceiveEthash)(nil).GetLogger))
}

// GetName mocks base method.
func (m *MockBridgeReceiveEthash) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockBridgeReceiveEthashMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockBridgeReceiveEthash)(nil).GetName))
}

// GetWsClient mocks base method.
func (m *MockBridgeReceiveEthash) GetWsClient() ethclients.ClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWsClient")
	ret0, _ := ret[0].(ethclients.ClientInterface)
	return ret0
}

// GetWsClient indicates an expected call of GetWsClient.
func (mr *MockBridgeReceiveEthashMockRecorder) GetWsClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWsClient", reflect.TypeOf((*MockBridgeReceiveEthash)(nil).GetWsClient))
}

// GetWsContract mocks base method.
func (m *MockBridgeReceiveEthash) GetWsContract() interfaces.BridgeContract {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWsContract")
	ret0, _ := ret[0].(interfaces.BridgeContract)
	return ret0
}

// GetWsContract indicates an expected call of GetWsContract.
func (mr *MockBridgeReceiveEthashMockRecorder) GetWsContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWsContract", reflect.TypeOf((*MockBridgeReceiveEthash)(nil).GetWsContract))
}

// IsEpochSet mocks base method.
func (m *MockBridgeReceiveEthash) IsEpochSet(epoch uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEpochSet", epoch)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsEpochSet indicates an expected call of IsEpochSet.
func (mr *MockBridgeReceiveEthashMockRecorder) IsEpochSet(epoch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEpochSet", reflect.TypeOf((*MockBridgeReceiveEthash)(nil).IsEpochSet), epoch)
}

// ProcessTx mocks base method.
func (m *MockBridgeReceiveEthash) ProcessTx(methodName string, txCallback networks.ContractCallFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTx", methodName, txCallback)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessTx indicates an expected call of ProcessTx.
func (mr *MockBridgeReceiveEthashMockRecorder) ProcessTx(methodName, txCallback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTx", reflect.TypeOf((*MockBridgeReceiveEthash)(nil).ProcessTx), methodName, txCallback)
}

// ShouldHavePk mocks base method.
func (m *MockBridgeReceiveEthash) ShouldHavePk() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShouldHavePk")
}

// ShouldHavePk indicates an expected call of ShouldHavePk.
func (mr *MockBridgeReceiveEthashMockRecorder) ShouldHavePk() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldHavePk", reflect.TypeOf((*MockBridgeReceiveEthash)(nil).ShouldHavePk))
}

// SubmitEpochData mocks base method.
func (m *MockBridgeReceiveEthash) SubmitEpochData(arg0 *ethash.EpochData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitEpochData", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitEpochData indicates an expected call of SubmitEpochData.
func (mr *MockBridgeReceiveEthashMockRecorder) SubmitEpochData(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitEpochData", reflect.TypeOf((*MockBridgeReceiveEthash)(nil).SubmitEpochData), arg0)
}

// SubmitTransferPoW mocks base method.
func (m *MockBridgeReceiveEthash) SubmitTransferPoW(arg0 *bindings.CheckPoWPoWProof) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitTransferPoW", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitTransferPoW indicates an expected call of SubmitTransferPoW.
func (mr *MockBridgeReceiveEthashMockRecorder) SubmitTransferPoW(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransferPoW", reflect.TypeOf((*MockBridgeReceiveEthash)(nil).SubmitTransferPoW), arg0)
}

// MockBridgeReceivePoSA is a mock of BridgeReceivePoSA interface.
type MockBridgeReceivePoSA struct {
	ctrl     *gomock.Controller
	recorder *MockBridgeReceivePoSAMockRecorder
}

// MockBridgeReceivePoSAMockRecorder is the mock recorder for MockBridgeReceivePoSA.
type MockBridgeReceivePoSAMockRecorder struct {
	mock *MockBridgeReceivePoSA
}

// NewMockBridgeReceivePoSA creates a new mock instance.
func NewMockBridgeReceivePoSA(ctrl *gomock.Controller) *MockBridgeReceivePoSA {
	mock := &MockBridgeReceivePoSA{ctrl: ctrl}
	mock.recorder = &MockBridgeReceivePoSAMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBridgeReceivePoSA) EXPECT() *MockBridgeReceivePoSAMockRecorder {
	return m.recorder
}

// EnsureContractUnpaused mocks base method.
func (m *MockBridgeReceivePoSA) EnsureContractUnpaused() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnsureContractUnpaused")
}

// EnsureContractUnpaused indicates an expected call of EnsureContractUnpaused.
func (mr *MockBridgeReceivePoSAMockRecorder) EnsureContractUnpaused() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureContractUnpaused", reflect.TypeOf((*MockBridgeReceivePoSA)(nil).EnsureContractUnpaused))
}

// GetClient mocks base method.
func (m *MockBridgeReceivePoSA) GetClient() ethclients.ClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(ethclients.ClientInterface)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockBridgeReceivePoSAMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockBridgeReceivePoSA)(nil).GetClient))
}

// GetContract mocks base method.
func (m *MockBridgeReceivePoSA) GetContract() interfaces.BridgeContract {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract")
	ret0, _ := ret[0].(interfaces.BridgeContract)
	return ret0
}

// GetContract indicates an expected call of GetContract.
func (mr *MockBridgeReceivePoSAMockRecorder) GetContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockBridgeReceivePoSA)(nil).GetContract))
}

// GetCurrentEpoch mocks base method.
func (m *MockBridgeReceivePoSA) GetCurrentEpoch() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentEpoch")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentEpoch indicates an expected call of GetCurrentEpoch.
func (mr *MockBridgeReceivePoSAMockRecorder) GetCurrentEpoch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentEpoch", reflect.TypeOf((*MockBridgeReceivePoSA)(nil).GetCurrentEpoch))
}

// GetLogger mocks base method.
func (m *MockBridgeReceivePoSA) GetLogger() *zerolog.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(*zerolog.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockBridgeReceivePoSAMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockBridgeReceivePoSA)(nil).GetLogger))
}

// GetName mocks base method.
func (m *MockBridgeReceivePoSA) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockBridgeReceivePoSAMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockBridgeReceivePoSA)(nil).GetName))
}

// GetWsClient mocks base method.
func (m *MockBridgeReceivePoSA) GetWsClient() ethclients.ClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWsClient")
	ret0, _ := ret[0].(ethclients.ClientInterface)
	return ret0
}

// GetWsClient indicates an expected call of GetWsClient.
func (mr *MockBridgeReceivePoSAMockRecorder) GetWsClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWsClient", reflect.TypeOf((*MockBridgeReceivePoSA)(nil).GetWsClient))
}

// GetWsContract mocks base method.
func (m *MockBridgeReceivePoSA) GetWsContract() interfaces.BridgeContract {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWsContract")
	ret0, _ := ret[0].(interfaces.BridgeContract)
	return ret0
}

// GetWsContract indicates an expected call of GetWsContract.
func (mr *MockBridgeReceivePoSAMockRecorder) GetWsContract() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWsContract", reflect.TypeOf((*MockBridgeReceivePoSA)(nil).GetWsContract))
}

// ProcessTx mocks base method.
func (m *MockBridgeReceivePoSA) ProcessTx(methodName string, txCallback networks.ContractCallFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTx", methodName, txCallback)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessTx indicates an expected call of ProcessTx.
func (mr *MockBridgeReceivePoSAMockRecorder) ProcessTx(methodName, txCallback interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTx", reflect.TypeOf((*MockBridgeReceivePoSA)(nil).ProcessTx), methodName, txCallback)
}

// ShouldHavePk mocks base method.
func (m *MockBridgeReceivePoSA) ShouldHavePk() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShouldHavePk")
}

// ShouldHavePk indicates an expected call of ShouldHavePk.
func (mr *MockBridgeReceivePoSAMockRecorder) ShouldHavePk() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldHavePk", reflect.TypeOf((*MockBridgeReceivePoSA)(nil).ShouldHavePk))
}

// SubmitTransferPoSA mocks base method.
func (m *MockBridgeReceivePoSA) SubmitTransferPoSA(proof *bindings.CheckPoSAPoSAProof) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitTransferPoSA", proof)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitTransferPoSA indicates an expected call of SubmitTransferPoSA.
func (mr *MockBridgeReceivePoSAMockRecorder) SubmitTransferPoSA(proof interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransferPoSA", reflect.TypeOf((*MockBridgeReceivePoSA)(nil).SubmitTransferPoSA), proof)
}