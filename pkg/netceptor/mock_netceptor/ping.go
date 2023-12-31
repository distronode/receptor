// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/netceptor/ping.go

// Package mock_netceptor is a generated GoMock package.
package mock_netceptor

import (
	context "context"
	reflect "reflect"
	time "time"

	netceptor "github.com/distronode/receptor/pkg/netceptor"
	gomock "github.com/golang/mock/gomock"
)

// MockNetcForPing is a mock of NetcForPing interface.
type MockNetcForPing struct {
	ctrl     *gomock.Controller
	recorder *MockNetcForPingMockRecorder
}

// MockNetcForPingMockRecorder is the mock recorder for MockNetcForPing.
type MockNetcForPingMockRecorder struct {
	mock *MockNetcForPing
}

// NewMockNetcForPing creates a new mock instance.
func NewMockNetcForPing(ctrl *gomock.Controller) *MockNetcForPing {
	mock := &MockNetcForPing{ctrl: ctrl}
	mock.recorder = &MockNetcForPingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetcForPing) EXPECT() *MockNetcForPingMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockNetcForPing) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNetcForPingMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNetcForPing)(nil).Context))
}

// ListenPacket mocks base method.
func (m *MockNetcForPing) ListenPacket(service string) (netceptor.PacketConner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenPacket", service)
	ret0, _ := ret[0].(netceptor.PacketConner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenPacket indicates an expected call of ListenPacket.
func (mr *MockNetcForPingMockRecorder) ListenPacket(service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenPacket", reflect.TypeOf((*MockNetcForPing)(nil).ListenPacket), service)
}

// NewAddr mocks base method.
func (m *MockNetcForPing) NewAddr(target, service string) netceptor.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAddr", target, service)
	ret0, _ := ret[0].(netceptor.Addr)
	return ret0
}

// NewAddr indicates an expected call of NewAddr.
func (mr *MockNetcForPingMockRecorder) NewAddr(target, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAddr", reflect.TypeOf((*MockNetcForPing)(nil).NewAddr), target, service)
}

// NodeID mocks base method.
func (m *MockNetcForPing) NodeID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeID")
	ret0, _ := ret[0].(string)
	return ret0
}

// NodeID indicates an expected call of NodeID.
func (mr *MockNetcForPingMockRecorder) NodeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeID", reflect.TypeOf((*MockNetcForPing)(nil).NodeID))
}

// MockNetcForTraceroute is a mock of NetcForTraceroute interface.
type MockNetcForTraceroute struct {
	ctrl     *gomock.Controller
	recorder *MockNetcForTracerouteMockRecorder
}

// MockNetcForTracerouteMockRecorder is the mock recorder for MockNetcForTraceroute.
type MockNetcForTracerouteMockRecorder struct {
	mock *MockNetcForTraceroute
}

// NewMockNetcForTraceroute creates a new mock instance.
func NewMockNetcForTraceroute(ctrl *gomock.Controller) *MockNetcForTraceroute {
	mock := &MockNetcForTraceroute{ctrl: ctrl}
	mock.recorder = &MockNetcForTracerouteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetcForTraceroute) EXPECT() *MockNetcForTracerouteMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockNetcForTraceroute) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockNetcForTracerouteMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockNetcForTraceroute)(nil).Context))
}

// MaxForwardingHops mocks base method.
func (m *MockNetcForTraceroute) MaxForwardingHops() byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxForwardingHops")
	ret0, _ := ret[0].(byte)
	return ret0
}

// MaxForwardingHops indicates an expected call of MaxForwardingHops.
func (mr *MockNetcForTracerouteMockRecorder) MaxForwardingHops() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxForwardingHops", reflect.TypeOf((*MockNetcForTraceroute)(nil).MaxForwardingHops))
}

// Ping mocks base method.
func (m *MockNetcForTraceroute) Ping(ctx context.Context, target string, hopsToLive byte) (time.Duration, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx, target, hopsToLive)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Ping indicates an expected call of Ping.
func (mr *MockNetcForTracerouteMockRecorder) Ping(ctx, target, hopsToLive interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockNetcForTraceroute)(nil).Ping), ctx, target, hopsToLive)
}
