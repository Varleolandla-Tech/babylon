// Code generated by MockGen. DO NOT EDIT.
// Source: x/zoneconcierge/types/expected_keepers.go

// Package types is a generated GoMock package.
package types

import (
	context "context"
	reflect "reflect"

	types "github.com/babylonchain/babylon/types"
	types0 "github.com/babylonchain/babylon/x/btccheckpoint/types"
	types1 "github.com/babylonchain/babylon/x/btclightclient/types"
	types2 "github.com/babylonchain/babylon/x/checkpointing/types"
	types3 "github.com/babylonchain/babylon/x/epoching/types"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	types4 "github.com/cosmos/cosmos-sdk/types"
	types5 "github.com/cosmos/ibc-go/modules/capability/types"
	types6 "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	types7 "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	types8 "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	exported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetModuleAccount mocks base method.
func (m *MockAccountKeeper) GetModuleAccount(ctx context.Context, name string) types4.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, name)
	ret0, _ := ret[0].(types4.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccount(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccount), ctx, name)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(name string) types4.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", name)
	ret0, _ := ret[0].(types4.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), name)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BlockedAddr mocks base method.
func (m *MockBankKeeper) BlockedAddr(addr types4.AccAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockedAddr", addr)
	ret0, _ := ret[0].(bool)
	return ret0
}

// BlockedAddr indicates an expected call of BlockedAddr.
func (mr *MockBankKeeperMockRecorder) BlockedAddr(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockedAddr", reflect.TypeOf((*MockBankKeeper)(nil).BlockedAddr), addr)
}

// BurnCoins mocks base method.
func (m *MockBankKeeper) BurnCoins(ctx context.Context, moduleName string, amt types4.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnCoins indicates an expected call of BurnCoins.
func (mr *MockBankKeeperMockRecorder) BurnCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnCoins", reflect.TypeOf((*MockBankKeeper)(nil).BurnCoins), ctx, moduleName, amt)
}

// MintCoins mocks base method.
func (m *MockBankKeeper) MintCoins(ctx context.Context, moduleName string, amt types4.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintCoins indicates an expected call of MintCoins.
func (mr *MockBankKeeperMockRecorder) MintCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCoins", reflect.TypeOf((*MockBankKeeper)(nil).MintCoins), ctx, moduleName, amt)
}

// SendCoins mocks base method.
func (m *MockBankKeeper) SendCoins(ctx context.Context, fromAddr, toAddr types4.AccAddress, amt types4.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoins", ctx, fromAddr, toAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoins indicates an expected call of SendCoins.
func (mr *MockBankKeeperMockRecorder) SendCoins(ctx, fromAddr, toAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoins", reflect.TypeOf((*MockBankKeeper)(nil).SendCoins), ctx, fromAddr, toAddr, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr types4.AccAddress, recipientModule string, amt types4.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr types4.AccAddress, amt types4.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// MockICS4Wrapper is a mock of ICS4Wrapper interface.
type MockICS4Wrapper struct {
	ctrl     *gomock.Controller
	recorder *MockICS4WrapperMockRecorder
}

// MockICS4WrapperMockRecorder is the mock recorder for MockICS4Wrapper.
type MockICS4WrapperMockRecorder struct {
	mock *MockICS4Wrapper
}

// NewMockICS4Wrapper creates a new mock instance.
func NewMockICS4Wrapper(ctrl *gomock.Controller) *MockICS4Wrapper {
	mock := &MockICS4Wrapper{ctrl: ctrl}
	mock.recorder = &MockICS4WrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICS4Wrapper) EXPECT() *MockICS4WrapperMockRecorder {
	return m.recorder
}

// SendPacket mocks base method.
func (m *MockICS4Wrapper) SendPacket(ctx types4.Context, channelCap *types5.Capability, sourcePort, sourceChannel string, timeoutHeight types6.Height, timeoutTimestamp uint64, data []byte) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendPacket", ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendPacket indicates an expected call of SendPacket.
func (mr *MockICS4WrapperMockRecorder) SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPacket", reflect.TypeOf((*MockICS4Wrapper)(nil).SendPacket), ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}

// MockChannelKeeper is a mock of ChannelKeeper interface.
type MockChannelKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockChannelKeeperMockRecorder
}

// MockChannelKeeperMockRecorder is the mock recorder for MockChannelKeeper.
type MockChannelKeeperMockRecorder struct {
	mock *MockChannelKeeper
}

// NewMockChannelKeeper creates a new mock instance.
func NewMockChannelKeeper(ctrl *gomock.Controller) *MockChannelKeeper {
	mock := &MockChannelKeeper{ctrl: ctrl}
	mock.recorder = &MockChannelKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelKeeper) EXPECT() *MockChannelKeeperMockRecorder {
	return m.recorder
}

// GetAllChannels mocks base method.
func (m *MockChannelKeeper) GetAllChannels(ctx types4.Context) []types8.IdentifiedChannel {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllChannels", ctx)
	ret0, _ := ret[0].([]types8.IdentifiedChannel)
	return ret0
}

// GetAllChannels indicates an expected call of GetAllChannels.
func (mr *MockChannelKeeperMockRecorder) GetAllChannels(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllChannels", reflect.TypeOf((*MockChannelKeeper)(nil).GetAllChannels), ctx)
}

// GetChannel mocks base method.
func (m *MockChannelKeeper) GetChannel(ctx types4.Context, srcPort, srcChan string) (types8.Channel, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannel", ctx, srcPort, srcChan)
	ret0, _ := ret[0].(types8.Channel)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetChannel indicates an expected call of GetChannel.
func (mr *MockChannelKeeperMockRecorder) GetChannel(ctx, srcPort, srcChan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannel", reflect.TypeOf((*MockChannelKeeper)(nil).GetChannel), ctx, srcPort, srcChan)
}

// GetChannelClientState mocks base method.
func (m *MockChannelKeeper) GetChannelClientState(ctx types4.Context, portID, channelID string) (string, exported.ClientState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannelClientState", ctx, portID, channelID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(exported.ClientState)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetChannelClientState indicates an expected call of GetChannelClientState.
func (mr *MockChannelKeeperMockRecorder) GetChannelClientState(ctx, portID, channelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannelClientState", reflect.TypeOf((*MockChannelKeeper)(nil).GetChannelClientState), ctx, portID, channelID)
}

// GetNextSequenceSend mocks base method.
func (m *MockChannelKeeper) GetNextSequenceSend(ctx types4.Context, portID, channelID string) (uint64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextSequenceSend", ctx, portID, channelID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetNextSequenceSend indicates an expected call of GetNextSequenceSend.
func (mr *MockChannelKeeperMockRecorder) GetNextSequenceSend(ctx, portID, channelID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextSequenceSend", reflect.TypeOf((*MockChannelKeeper)(nil).GetNextSequenceSend), ctx, portID, channelID)
}

// MockClientKeeper is a mock of ClientKeeper interface.
type MockClientKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockClientKeeperMockRecorder
}

// MockClientKeeperMockRecorder is the mock recorder for MockClientKeeper.
type MockClientKeeperMockRecorder struct {
	mock *MockClientKeeper
}

// NewMockClientKeeper creates a new mock instance.
func NewMockClientKeeper(ctrl *gomock.Controller) *MockClientKeeper {
	mock := &MockClientKeeper{ctrl: ctrl}
	mock.recorder = &MockClientKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientKeeper) EXPECT() *MockClientKeeperMockRecorder {
	return m.recorder
}

// GetClientState mocks base method.
func (m *MockClientKeeper) GetClientState(ctx types4.Context, clientID string) (exported.ClientState, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientState", ctx, clientID)
	ret0, _ := ret[0].(exported.ClientState)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetClientState indicates an expected call of GetClientState.
func (mr *MockClientKeeperMockRecorder) GetClientState(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientState", reflect.TypeOf((*MockClientKeeper)(nil).GetClientState), ctx, clientID)
}

// SetClientState mocks base method.
func (m *MockClientKeeper) SetClientState(ctx types4.Context, clientID string, clientState exported.ClientState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClientState", ctx, clientID, clientState)
}

// SetClientState indicates an expected call of SetClientState.
func (mr *MockClientKeeperMockRecorder) SetClientState(ctx, clientID, clientState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClientState", reflect.TypeOf((*MockClientKeeper)(nil).SetClientState), ctx, clientID, clientState)
}

// MockConnectionKeeper is a mock of ConnectionKeeper interface.
type MockConnectionKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionKeeperMockRecorder
}

// MockConnectionKeeperMockRecorder is the mock recorder for MockConnectionKeeper.
type MockConnectionKeeperMockRecorder struct {
	mock *MockConnectionKeeper
}

// NewMockConnectionKeeper creates a new mock instance.
func NewMockConnectionKeeper(ctrl *gomock.Controller) *MockConnectionKeeper {
	mock := &MockConnectionKeeper{ctrl: ctrl}
	mock.recorder = &MockConnectionKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionKeeper) EXPECT() *MockConnectionKeeperMockRecorder {
	return m.recorder
}

// GetConnection mocks base method.
func (m *MockConnectionKeeper) GetConnection(ctx types4.Context, connectionID string) (types7.ConnectionEnd, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnection", ctx, connectionID)
	ret0, _ := ret[0].(types7.ConnectionEnd)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetConnection indicates an expected call of GetConnection.
func (mr *MockConnectionKeeperMockRecorder) GetConnection(ctx, connectionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnection", reflect.TypeOf((*MockConnectionKeeper)(nil).GetConnection), ctx, connectionID)
}

// MockPortKeeper is a mock of PortKeeper interface.
type MockPortKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockPortKeeperMockRecorder
}

// MockPortKeeperMockRecorder is the mock recorder for MockPortKeeper.
type MockPortKeeperMockRecorder struct {
	mock *MockPortKeeper
}

// NewMockPortKeeper creates a new mock instance.
func NewMockPortKeeper(ctrl *gomock.Controller) *MockPortKeeper {
	mock := &MockPortKeeper{ctrl: ctrl}
	mock.recorder = &MockPortKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPortKeeper) EXPECT() *MockPortKeeperMockRecorder {
	return m.recorder
}

// BindPort mocks base method.
func (m *MockPortKeeper) BindPort(ctx types4.Context, portID string) *types5.Capability {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindPort", ctx, portID)
	ret0, _ := ret[0].(*types5.Capability)
	return ret0
}

// BindPort indicates an expected call of BindPort.
func (mr *MockPortKeeperMockRecorder) BindPort(ctx, portID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindPort", reflect.TypeOf((*MockPortKeeper)(nil).BindPort), ctx, portID)
}

// MockScopedKeeper is a mock of ScopedKeeper interface.
type MockScopedKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockScopedKeeperMockRecorder
}

// MockScopedKeeperMockRecorder is the mock recorder for MockScopedKeeper.
type MockScopedKeeperMockRecorder struct {
	mock *MockScopedKeeper
}

// NewMockScopedKeeper creates a new mock instance.
func NewMockScopedKeeper(ctrl *gomock.Controller) *MockScopedKeeper {
	mock := &MockScopedKeeper{ctrl: ctrl}
	mock.recorder = &MockScopedKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScopedKeeper) EXPECT() *MockScopedKeeperMockRecorder {
	return m.recorder
}

// AuthenticateCapability mocks base method.
func (m *MockScopedKeeper) AuthenticateCapability(ctx types4.Context, cap *types5.Capability, name string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticateCapability", ctx, cap, name)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthenticateCapability indicates an expected call of AuthenticateCapability.
func (mr *MockScopedKeeperMockRecorder) AuthenticateCapability(ctx, cap, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticateCapability", reflect.TypeOf((*MockScopedKeeper)(nil).AuthenticateCapability), ctx, cap, name)
}

// ClaimCapability mocks base method.
func (m *MockScopedKeeper) ClaimCapability(ctx types4.Context, cap *types5.Capability, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClaimCapability", ctx, cap, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClaimCapability indicates an expected call of ClaimCapability.
func (mr *MockScopedKeeperMockRecorder) ClaimCapability(ctx, cap, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClaimCapability", reflect.TypeOf((*MockScopedKeeper)(nil).ClaimCapability), ctx, cap, name)
}

// GetCapability mocks base method.
func (m *MockScopedKeeper) GetCapability(ctx types4.Context, name string) (*types5.Capability, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCapability", ctx, name)
	ret0, _ := ret[0].(*types5.Capability)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetCapability indicates an expected call of GetCapability.
func (mr *MockScopedKeeperMockRecorder) GetCapability(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCapability", reflect.TypeOf((*MockScopedKeeper)(nil).GetCapability), ctx, name)
}

// LookupModules mocks base method.
func (m *MockScopedKeeper) LookupModules(ctx types4.Context, name string) ([]string, *types5.Capability, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupModules", ctx, name)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(*types5.Capability)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LookupModules indicates an expected call of LookupModules.
func (mr *MockScopedKeeperMockRecorder) LookupModules(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupModules", reflect.TypeOf((*MockScopedKeeper)(nil).LookupModules), ctx, name)
}

// MockBTCLightClientKeeper is a mock of BTCLightClientKeeper interface.
type MockBTCLightClientKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBTCLightClientKeeperMockRecorder
}

// MockBTCLightClientKeeperMockRecorder is the mock recorder for MockBTCLightClientKeeper.
type MockBTCLightClientKeeperMockRecorder struct {
	mock *MockBTCLightClientKeeper
}

// NewMockBTCLightClientKeeper creates a new mock instance.
func NewMockBTCLightClientKeeper(ctrl *gomock.Controller) *MockBTCLightClientKeeper {
	mock := &MockBTCLightClientKeeper{ctrl: ctrl}
	mock.recorder = &MockBTCLightClientKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBTCLightClientKeeper) EXPECT() *MockBTCLightClientKeeperMockRecorder {
	return m.recorder
}

// GetHeaderByHash mocks base method.
func (m *MockBTCLightClientKeeper) GetHeaderByHash(ctx context.Context, hash *types.BTCHeaderHashBytes) *types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaderByHash", ctx, hash)
	ret0, _ := ret[0].(*types1.BTCHeaderInfo)
	return ret0
}

// GetHeaderByHash indicates an expected call of GetHeaderByHash.
func (mr *MockBTCLightClientKeeperMockRecorder) GetHeaderByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaderByHash", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetHeaderByHash), ctx, hash)
}

// GetMainChainFrom mocks base method.
func (m *MockBTCLightClientKeeper) GetMainChainFrom(ctx context.Context, startHeight uint64) []*types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMainChainFrom", ctx, startHeight)
	ret0, _ := ret[0].([]*types1.BTCHeaderInfo)
	return ret0
}

// GetMainChainFrom indicates an expected call of GetMainChainFrom.
func (mr *MockBTCLightClientKeeperMockRecorder) GetMainChainFrom(ctx, startHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMainChainFrom", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetMainChainFrom), ctx, startHeight)
}

// GetMainChainUpTo mocks base method.
func (m *MockBTCLightClientKeeper) GetMainChainUpTo(ctx context.Context, depth uint64) []*types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMainChainUpTo", ctx, depth)
	ret0, _ := ret[0].([]*types1.BTCHeaderInfo)
	return ret0
}

// GetMainChainUpTo indicates an expected call of GetMainChainUpTo.
func (mr *MockBTCLightClientKeeperMockRecorder) GetMainChainUpTo(ctx, depth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMainChainUpTo", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetMainChainUpTo), ctx, depth)
}

// GetTipInfo mocks base method.
func (m *MockBTCLightClientKeeper) GetTipInfo(ctx context.Context) *types1.BTCHeaderInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTipInfo", ctx)
	ret0, _ := ret[0].(*types1.BTCHeaderInfo)
	return ret0
}

// GetTipInfo indicates an expected call of GetTipInfo.
func (mr *MockBTCLightClientKeeperMockRecorder) GetTipInfo(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTipInfo", reflect.TypeOf((*MockBTCLightClientKeeper)(nil).GetTipInfo), ctx)
}

// MockBtcCheckpointKeeper is a mock of BtcCheckpointKeeper interface.
type MockBtcCheckpointKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBtcCheckpointKeeperMockRecorder
}

// MockBtcCheckpointKeeperMockRecorder is the mock recorder for MockBtcCheckpointKeeper.
type MockBtcCheckpointKeeperMockRecorder struct {
	mock *MockBtcCheckpointKeeper
}

// NewMockBtcCheckpointKeeper creates a new mock instance.
func NewMockBtcCheckpointKeeper(ctrl *gomock.Controller) *MockBtcCheckpointKeeper {
	mock := &MockBtcCheckpointKeeper{ctrl: ctrl}
	mock.recorder = &MockBtcCheckpointKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBtcCheckpointKeeper) EXPECT() *MockBtcCheckpointKeeperMockRecorder {
	return m.recorder
}

// GetBestSubmission mocks base method.
func (m *MockBtcCheckpointKeeper) GetBestSubmission(ctx context.Context, e uint64) (types0.BtcStatus, *types0.SubmissionKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBestSubmission", ctx, e)
	ret0, _ := ret[0].(types0.BtcStatus)
	ret1, _ := ret[1].(*types0.SubmissionKey)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBestSubmission indicates an expected call of GetBestSubmission.
func (mr *MockBtcCheckpointKeeperMockRecorder) GetBestSubmission(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBestSubmission", reflect.TypeOf((*MockBtcCheckpointKeeper)(nil).GetBestSubmission), ctx, e)
}

// GetEpochBestSubmissionBtcInfo mocks base method.
func (m *MockBtcCheckpointKeeper) GetEpochBestSubmissionBtcInfo(ctx context.Context, ed *types0.EpochData) *types0.SubmissionBtcInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpochBestSubmissionBtcInfo", ctx, ed)
	ret0, _ := ret[0].(*types0.SubmissionBtcInfo)
	return ret0
}

// GetEpochBestSubmissionBtcInfo indicates an expected call of GetEpochBestSubmissionBtcInfo.
func (mr *MockBtcCheckpointKeeperMockRecorder) GetEpochBestSubmissionBtcInfo(ctx, ed interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpochBestSubmissionBtcInfo", reflect.TypeOf((*MockBtcCheckpointKeeper)(nil).GetEpochBestSubmissionBtcInfo), ctx, ed)
}

// GetEpochData mocks base method.
func (m *MockBtcCheckpointKeeper) GetEpochData(ctx context.Context, e uint64) *types0.EpochData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpochData", ctx, e)
	ret0, _ := ret[0].(*types0.EpochData)
	return ret0
}

// GetEpochData indicates an expected call of GetEpochData.
func (mr *MockBtcCheckpointKeeperMockRecorder) GetEpochData(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpochData", reflect.TypeOf((*MockBtcCheckpointKeeper)(nil).GetEpochData), ctx, e)
}

// GetParams mocks base method.
func (m *MockBtcCheckpointKeeper) GetParams(ctx context.Context) types0.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types0.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockBtcCheckpointKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockBtcCheckpointKeeper)(nil).GetParams), ctx)
}

// GetSubmissionData mocks base method.
func (m *MockBtcCheckpointKeeper) GetSubmissionData(ctx context.Context, sk types0.SubmissionKey) *types0.SubmissionData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubmissionData", ctx, sk)
	ret0, _ := ret[0].(*types0.SubmissionData)
	return ret0
}

// GetSubmissionData indicates an expected call of GetSubmissionData.
func (mr *MockBtcCheckpointKeeperMockRecorder) GetSubmissionData(ctx, sk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubmissionData", reflect.TypeOf((*MockBtcCheckpointKeeper)(nil).GetSubmissionData), ctx, sk)
}

// MockCheckpointingKeeper is a mock of CheckpointingKeeper interface.
type MockCheckpointingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockCheckpointingKeeperMockRecorder
}

// MockCheckpointingKeeperMockRecorder is the mock recorder for MockCheckpointingKeeper.
type MockCheckpointingKeeperMockRecorder struct {
	mock *MockCheckpointingKeeper
}

// NewMockCheckpointingKeeper creates a new mock instance.
func NewMockCheckpointingKeeper(ctrl *gomock.Controller) *MockCheckpointingKeeper {
	mock := &MockCheckpointingKeeper{ctrl: ctrl}
	mock.recorder = &MockCheckpointingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCheckpointingKeeper) EXPECT() *MockCheckpointingKeeperMockRecorder {
	return m.recorder
}

// GetBLSPubKeySet mocks base method.
func (m *MockCheckpointingKeeper) GetBLSPubKeySet(ctx context.Context, epochNumber uint64) ([]*types2.ValidatorWithBlsKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBLSPubKeySet", ctx, epochNumber)
	ret0, _ := ret[0].([]*types2.ValidatorWithBlsKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBLSPubKeySet indicates an expected call of GetBLSPubKeySet.
func (mr *MockCheckpointingKeeperMockRecorder) GetBLSPubKeySet(ctx, epochNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBLSPubKeySet", reflect.TypeOf((*MockCheckpointingKeeper)(nil).GetBLSPubKeySet), ctx, epochNumber)
}

// GetFinalizedEpoch mocks base method.
func (m *MockCheckpointingKeeper) GetFinalizedEpoch(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFinalizedEpoch", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFinalizedEpoch indicates an expected call of GetFinalizedEpoch.
func (mr *MockCheckpointingKeeperMockRecorder) GetFinalizedEpoch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFinalizedEpoch", reflect.TypeOf((*MockCheckpointingKeeper)(nil).GetFinalizedEpoch), ctx)
}

// GetRawCheckpoint mocks base method.
func (m *MockCheckpointingKeeper) GetRawCheckpoint(ctx context.Context, epochNumber uint64) (*types2.RawCheckpointWithMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawCheckpoint", ctx, epochNumber)
	ret0, _ := ret[0].(*types2.RawCheckpointWithMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawCheckpoint indicates an expected call of GetRawCheckpoint.
func (mr *MockCheckpointingKeeperMockRecorder) GetRawCheckpoint(ctx, epochNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawCheckpoint", reflect.TypeOf((*MockCheckpointingKeeper)(nil).GetRawCheckpoint), ctx, epochNumber)
}

// MockEpochingKeeper is a mock of EpochingKeeper interface.
type MockEpochingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockEpochingKeeperMockRecorder
}

// MockEpochingKeeperMockRecorder is the mock recorder for MockEpochingKeeper.
type MockEpochingKeeperMockRecorder struct {
	mock *MockEpochingKeeper
}

// NewMockEpochingKeeper creates a new mock instance.
func NewMockEpochingKeeper(ctrl *gomock.Controller) *MockEpochingKeeper {
	mock := &MockEpochingKeeper{ctrl: ctrl}
	mock.recorder = &MockEpochingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEpochingKeeper) EXPECT() *MockEpochingKeeperMockRecorder {
	return m.recorder
}

// GetEpoch mocks base method.
func (m *MockEpochingKeeper) GetEpoch(ctx context.Context) *types3.Epoch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEpoch", ctx)
	ret0, _ := ret[0].(*types3.Epoch)
	return ret0
}

// GetEpoch indicates an expected call of GetEpoch.
func (mr *MockEpochingKeeperMockRecorder) GetEpoch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEpoch", reflect.TypeOf((*MockEpochingKeeper)(nil).GetEpoch), ctx)
}

// GetHistoricalEpoch mocks base method.
func (m *MockEpochingKeeper) GetHistoricalEpoch(ctx context.Context, epochNumber uint64) (*types3.Epoch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoricalEpoch", ctx, epochNumber)
	ret0, _ := ret[0].(*types3.Epoch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistoricalEpoch indicates an expected call of GetHistoricalEpoch.
func (mr *MockEpochingKeeperMockRecorder) GetHistoricalEpoch(ctx, epochNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoricalEpoch", reflect.TypeOf((*MockEpochingKeeper)(nil).GetHistoricalEpoch), ctx, epochNumber)
}

// MockCometClient is a mock of CometClient interface.
type MockCometClient struct {
	ctrl     *gomock.Controller
	recorder *MockCometClientMockRecorder
}

// MockCometClientMockRecorder is the mock recorder for MockCometClient.
type MockCometClientMockRecorder struct {
	mock *MockCometClient
}

// NewMockCometClient creates a new mock instance.
func NewMockCometClient(ctrl *gomock.Controller) *MockCometClient {
	mock := &MockCometClient{ctrl: ctrl}
	mock.recorder = &MockCometClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCometClient) EXPECT() *MockCometClientMockRecorder {
	return m.recorder
}

// Tx mocks base method.
func (m *MockCometClient) Tx(ctx context.Context, hash []byte, prove bool) (*coretypes.ResultTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx", ctx, hash, prove)
	ret0, _ := ret[0].(*coretypes.ResultTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tx indicates an expected call of Tx.
func (mr *MockCometClientMockRecorder) Tx(ctx, hash, prove interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockCometClient)(nil).Tx), ctx, hash, prove)
}
