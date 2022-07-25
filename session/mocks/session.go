// Code generated by MockGen. DO NOT EDIT.
// Source: session/session.go

// Package mock_session is a generated GoMock package.
package mocks

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	nats "github.com/nats-io/nats.go"
	networkentity "github.com/topfreegames/pitaya/v2/networkentity"
	session "github.com/topfreegames/pitaya/v2/session"
)

// MockSessionPool is a mock of SessionPool interface.
type MockSessionPool struct {
	ctrl     *gomock.Controller
	recorder *MockSessionPoolMockRecorder
}

// MockSessionPoolMockRecorder is the mock recorder for MockSessionPool.
type MockSessionPoolMockRecorder struct {
	mock *MockSessionPool
}

// NewMockSessionPool creates a new mock instance.
func NewMockSessionPool(ctrl *gomock.Controller) *MockSessionPool {
	mock := &MockSessionPool{ctrl: ctrl}
	mock.recorder = &MockSessionPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionPool) EXPECT() *MockSessionPoolMockRecorder {
	return m.recorder
}

// CloseAll mocks base method.
func (m *MockSessionPool) CloseAll() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CloseAll")
}

// CloseAll indicates an expected call of CloseAll.
func (mr *MockSessionPoolMockRecorder) CloseAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAll", reflect.TypeOf((*MockSessionPool)(nil).CloseAll))
}

// GetSessionByID mocks base method.
func (m *MockSessionPool) GetSessionByID(id int64) session.Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionByID", id)
	ret0, _ := ret[0].(session.Session)
	return ret0
}

// GetSessionByID indicates an expected call of GetSessionByID.
func (mr *MockSessionPoolMockRecorder) GetSessionByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByID", reflect.TypeOf((*MockSessionPool)(nil).GetSessionByID), id)
}

// GetSessionByUID mocks base method.
func (m *MockSessionPool) GetSessionByUID(uid string) session.Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionByUID", uid)
	ret0, _ := ret[0].(session.Session)
	return ret0
}

// GetSessionByUID indicates an expected call of GetSessionByUID.
func (mr *MockSessionPoolMockRecorder) GetSessionByUID(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByUID", reflect.TypeOf((*MockSessionPool)(nil).GetSessionByUID), uid)
}

// GetSessionCloseCallbacks mocks base method.
func (m *MockSessionPool) GetSessionCloseCallbacks() []func(session.Session) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionCloseCallbacks")
	ret0, _ := ret[0].([]func(session.Session))
	return ret0
}

// GetSessionCloseCallbacks indicates an expected call of GetSessionCloseCallbacks.
func (mr *MockSessionPoolMockRecorder) GetSessionCloseCallbacks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionCloseCallbacks", reflect.TypeOf((*MockSessionPool)(nil).GetSessionCloseCallbacks))
}

// GetSessionCount mocks base method.
func (m *MockSessionPool) GetSessionCount() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionCount")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetSessionCount indicates an expected call of GetSessionCount.
func (mr *MockSessionPoolMockRecorder) GetSessionCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionCount", reflect.TypeOf((*MockSessionPool)(nil).GetSessionCount))
}

// NewSession mocks base method.
func (m *MockSessionPool) NewSession(entity networkentity.NetworkEntity, frontend bool, UID ...string) session.Session {
	m.ctrl.T.Helper()
	varargs := []interface{}{entity, frontend}
	for _, a := range UID {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewSession", varargs...)
	ret0, _ := ret[0].(session.Session)
	return ret0
}

// NewSession indicates an expected call of NewSession.
func (mr *MockSessionPoolMockRecorder) NewSession(entity, frontend interface{}, UID ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{entity, frontend}, UID...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSession", reflect.TypeOf((*MockSessionPool)(nil).NewSession), varargs...)
}

// OnAfterSessionBind mocks base method.
func (m *MockSessionPool) OnAfterSessionBind(f func(context.Context, session.Session) error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnAfterSessionBind", f)
}

// OnAfterSessionBind indicates an expected call of OnAfterSessionBind.
func (mr *MockSessionPoolMockRecorder) OnAfterSessionBind(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnAfterSessionBind", reflect.TypeOf((*MockSessionPool)(nil).OnAfterSessionBind), f)
}

// OnSessionBind mocks base method.
func (m *MockSessionPool) OnSessionBind(f func(context.Context, session.Session) error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnSessionBind", f)
}

// OnSessionBind indicates an expected call of OnSessionBind.
func (mr *MockSessionPoolMockRecorder) OnSessionBind(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnSessionBind", reflect.TypeOf((*MockSessionPool)(nil).OnSessionBind), f)
}

// OnSessionClose mocks base method.
func (m *MockSessionPool) OnSessionClose(f func(session.Session)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnSessionClose", f)
}

// OnSessionClose indicates an expected call of OnSessionClose.
func (mr *MockSessionPoolMockRecorder) OnSessionClose(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnSessionClose", reflect.TypeOf((*MockSessionPool)(nil).OnSessionClose), f)
}

// MockSession is a mock of Session interface.
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession.
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance.
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// Bind mocks base method.
func (m *MockSession) Bind(ctx context.Context, uid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", ctx, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind.
func (mr *MockSessionMockRecorder) Bind(ctx, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockSession)(nil).Bind), ctx, uid)
}

// Clear mocks base method.
func (m *MockSession) Clear() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Clear")
}

// Clear indicates an expected call of Clear.
func (mr *MockSessionMockRecorder) Clear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockSession)(nil).Clear))
}

// Close mocks base method.
func (m *MockSession) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockSessionMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSession)(nil).Close))
}

// Float32 mocks base method.
func (m *MockSession) Float32(key string) float32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Float32", key)
	ret0, _ := ret[0].(float32)
	return ret0
}

// Float32 indicates an expected call of Float32.
func (mr *MockSessionMockRecorder) Float32(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float32", reflect.TypeOf((*MockSession)(nil).Float32), key)
}

// Float64 mocks base method.
func (m *MockSession) Float64(key string) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Float64", key)
	ret0, _ := ret[0].(float64)
	return ret0
}

// Float64 indicates an expected call of Float64.
func (mr *MockSessionMockRecorder) Float64(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float64", reflect.TypeOf((*MockSession)(nil).Float64), key)
}

// Get mocks base method.
func (m *MockSession) Get(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockSessionMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSession)(nil).Get), key)
}

// GetData mocks base method.
func (m *MockSession) GetData() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// GetData indicates an expected call of GetData.
func (mr *MockSessionMockRecorder) GetData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockSession)(nil).GetData))
}

// GetDataEncoded mocks base method.
func (m *MockSession) GetDataEncoded() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataEncoded")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetDataEncoded indicates an expected call of GetDataEncoded.
func (mr *MockSessionMockRecorder) GetDataEncoded() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataEncoded", reflect.TypeOf((*MockSession)(nil).GetDataEncoded))
}

// GetHandshakeData mocks base method.
func (m *MockSession) GetHandshakeData() *session.HandshakeData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandshakeData")
	ret0, _ := ret[0].(*session.HandshakeData)
	return ret0
}

// GetHandshakeData indicates an expected call of GetHandshakeData.
func (mr *MockSessionMockRecorder) GetHandshakeData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandshakeData", reflect.TypeOf((*MockSession)(nil).GetHandshakeData))
}

// GetIsFrontend mocks base method.
func (m *MockSession) GetIsFrontend() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIsFrontend")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetIsFrontend indicates an expected call of GetIsFrontend.
func (mr *MockSessionMockRecorder) GetIsFrontend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIsFrontend", reflect.TypeOf((*MockSession)(nil).GetIsFrontend))
}

// GetOnCloseCallbacks mocks base method.
func (m *MockSession) GetOnCloseCallbacks() []func() {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOnCloseCallbacks")
	ret0, _ := ret[0].([]func())
	return ret0
}

// GetOnCloseCallbacks indicates an expected call of GetOnCloseCallbacks.
func (mr *MockSessionMockRecorder) GetOnCloseCallbacks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOnCloseCallbacks", reflect.TypeOf((*MockSession)(nil).GetOnCloseCallbacks))
}

// GetSubscriptions mocks base method.
func (m *MockSession) GetSubscriptions() []*nats.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptions")
	ret0, _ := ret[0].([]*nats.Subscription)
	return ret0
}

// GetSubscriptions indicates an expected call of GetSubscriptions.
func (mr *MockSessionMockRecorder) GetSubscriptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptions", reflect.TypeOf((*MockSession)(nil).GetSubscriptions))
}

// HasKey mocks base method.
func (m *MockSession) HasKey(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasKey", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasKey indicates an expected call of HasKey.
func (mr *MockSessionMockRecorder) HasKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasKey", reflect.TypeOf((*MockSession)(nil).HasKey), key)
}

// ID mocks base method.
func (m *MockSession) ID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockSessionMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockSession)(nil).ID))
}

// Int mocks base method.
func (m *MockSession) Int(key string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int", key)
	ret0, _ := ret[0].(int)
	return ret0
}

// Int indicates an expected call of Int.
func (mr *MockSessionMockRecorder) Int(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockSession)(nil).Int), key)
}

// Int16 mocks base method.
func (m *MockSession) Int16(key string) int16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int16", key)
	ret0, _ := ret[0].(int16)
	return ret0
}

// Int16 indicates an expected call of Int16.
func (mr *MockSessionMockRecorder) Int16(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int16", reflect.TypeOf((*MockSession)(nil).Int16), key)
}

// Int32 mocks base method.
func (m *MockSession) Int32(key string) int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int32", key)
	ret0, _ := ret[0].(int32)
	return ret0
}

// Int32 indicates an expected call of Int32.
func (mr *MockSessionMockRecorder) Int32(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int32", reflect.TypeOf((*MockSession)(nil).Int32), key)
}

// Int64 mocks base method.
func (m *MockSession) Int64(key string) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int64", key)
	ret0, _ := ret[0].(int64)
	return ret0
}

// Int64 indicates an expected call of Int64.
func (mr *MockSessionMockRecorder) Int64(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64", reflect.TypeOf((*MockSession)(nil).Int64), key)
}

// Int8 mocks base method.
func (m *MockSession) Int8(key string) int8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int8", key)
	ret0, _ := ret[0].(int8)
	return ret0
}

// Int8 indicates an expected call of Int8.
func (mr *MockSessionMockRecorder) Int8(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int8", reflect.TypeOf((*MockSession)(nil).Int8), key)
}

// Kick mocks base method.
func (m *MockSession) Kick(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kick", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Kick indicates an expected call of Kick.
func (mr *MockSessionMockRecorder) Kick(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kick", reflect.TypeOf((*MockSession)(nil).Kick), ctx)
}

// OnClose mocks base method.
func (m *MockSession) OnClose(c func()) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnClose", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnClose indicates an expected call of OnClose.
func (mr *MockSessionMockRecorder) OnClose(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnClose", reflect.TypeOf((*MockSession)(nil).OnClose), c)
}

// Push mocks base method.
func (m *MockSession) Push(route string, v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Push", route, v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Push indicates an expected call of Push.
func (mr *MockSessionMockRecorder) Push(route, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Push", reflect.TypeOf((*MockSession)(nil).Push), route, v)
}

// PushToFront mocks base method.
func (m *MockSession) PushToFront(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PushToFront", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// PushToFront indicates an expected call of PushToFront.
func (mr *MockSessionMockRecorder) PushToFront(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushToFront", reflect.TypeOf((*MockSession)(nil).PushToFront), ctx)
}

// RemoteAddr mocks base method.
func (m *MockSession) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockSessionMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockSession)(nil).RemoteAddr))
}

// Remove mocks base method.
func (m *MockSession) Remove(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockSessionMockRecorder) Remove(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockSession)(nil).Remove), key)
}

// ResponseMID mocks base method.
func (m *MockSession) ResponseMID(ctx context.Context, mid uint, v interface{}, err ...bool) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, mid, v}
	for _, a := range err {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResponseMID", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResponseMID indicates an expected call of ResponseMID.
func (mr *MockSessionMockRecorder) ResponseMID(ctx, mid, v interface{}, err ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, mid, v}, err...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseMID", reflect.TypeOf((*MockSession)(nil).ResponseMID), varargs...)
}

// Set mocks base method.
func (m *MockSession) Set(key string, value interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockSessionMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockSession)(nil).Set), key, value)
}

// SetData mocks base method.
func (m *MockSession) SetData(data map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetData", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetData indicates an expected call of SetData.
func (mr *MockSessionMockRecorder) SetData(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetData", reflect.TypeOf((*MockSession)(nil).SetData), data)
}

// SetDataEncoded mocks base method.
func (m *MockSession) SetDataEncoded(encodedData []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDataEncoded", encodedData)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDataEncoded indicates an expected call of SetDataEncoded.
func (mr *MockSessionMockRecorder) SetDataEncoded(encodedData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDataEncoded", reflect.TypeOf((*MockSession)(nil).SetDataEncoded), encodedData)
}

// SetFrontendData mocks base method.
func (m *MockSession) SetFrontendData(frontendID string, frontendSessionID int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFrontendData", frontendID, frontendSessionID)
}

// SetFrontendData indicates an expected call of SetFrontendData.
func (mr *MockSessionMockRecorder) SetFrontendData(frontendID, frontendSessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFrontendData", reflect.TypeOf((*MockSession)(nil).SetFrontendData), frontendID, frontendSessionID)
}

// SetHandshakeData mocks base method.
func (m *MockSession) SetHandshakeData(data *session.HandshakeData) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHandshakeData", data)
}

// SetHandshakeData indicates an expected call of SetHandshakeData.
func (mr *MockSessionMockRecorder) SetHandshakeData(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHandshakeData", reflect.TypeOf((*MockSession)(nil).SetHandshakeData), data)
}

// SetIsFrontend mocks base method.
func (m *MockSession) SetIsFrontend(isFrontend bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetIsFrontend", isFrontend)
}

// SetIsFrontend indicates an expected call of SetIsFrontend.
func (mr *MockSessionMockRecorder) SetIsFrontend(isFrontend interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIsFrontend", reflect.TypeOf((*MockSession)(nil).SetIsFrontend), isFrontend)
}

// SetOnCloseCallbacks mocks base method.
func (m *MockSession) SetOnCloseCallbacks(callbacks []func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOnCloseCallbacks", callbacks)
}

// SetOnCloseCallbacks indicates an expected call of SetOnCloseCallbacks.
func (mr *MockSessionMockRecorder) SetOnCloseCallbacks(callbacks interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOnCloseCallbacks", reflect.TypeOf((*MockSession)(nil).SetOnCloseCallbacks), callbacks)
}

// SetSubscriptions mocks base method.
func (m *MockSession) SetSubscriptions(subscriptions []*nats.Subscription) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSubscriptions", subscriptions)
}

// SetSubscriptions indicates an expected call of SetSubscriptions.
func (mr *MockSessionMockRecorder) SetSubscriptions(subscriptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSubscriptions", reflect.TypeOf((*MockSession)(nil).SetSubscriptions), subscriptions)
}

// String mocks base method.
func (m *MockSession) String(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockSessionMockRecorder) String(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockSession)(nil).String), key)
}

// UID mocks base method.
func (m *MockSession) UID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UID")
	ret0, _ := ret[0].(string)
	return ret0
}

// UID indicates an expected call of UID.
func (mr *MockSessionMockRecorder) UID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UID", reflect.TypeOf((*MockSession)(nil).UID))
}

// Uint mocks base method.
func (m *MockSession) Uint(key string) uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uint", key)
	ret0, _ := ret[0].(uint)
	return ret0
}

// Uint indicates an expected call of Uint.
func (mr *MockSessionMockRecorder) Uint(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint", reflect.TypeOf((*MockSession)(nil).Uint), key)
}

// Uint16 mocks base method.
func (m *MockSession) Uint16(key string) uint16 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uint16", key)
	ret0, _ := ret[0].(uint16)
	return ret0
}

// Uint16 indicates an expected call of Uint16.
func (mr *MockSessionMockRecorder) Uint16(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint16", reflect.TypeOf((*MockSession)(nil).Uint16), key)
}

// Uint32 mocks base method.
func (m *MockSession) Uint32(key string) uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uint32", key)
	ret0, _ := ret[0].(uint32)
	return ret0
}

// Uint32 indicates an expected call of Uint32.
func (mr *MockSessionMockRecorder) Uint32(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint32", reflect.TypeOf((*MockSession)(nil).Uint32), key)
}

// Uint64 mocks base method.
func (m *MockSession) Uint64(key string) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uint64", key)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Uint64 indicates an expected call of Uint64.
func (mr *MockSessionMockRecorder) Uint64(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint64", reflect.TypeOf((*MockSession)(nil).Uint64), key)
}

// Uint8 mocks base method.
func (m *MockSession) Uint8(key string) uint8 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uint8", key)
	ret0, _ := ret[0].(uint8)
	return ret0
}

// Uint8 indicates an expected call of Uint8.
func (mr *MockSessionMockRecorder) Uint8(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint8", reflect.TypeOf((*MockSession)(nil).Uint8), key)
}

// Value mocks base method.
func (m *MockSession) Value(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockSessionMockRecorder) Value(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockSession)(nil).Value), key)
}
