// Code generated by MockGen. DO NOT EDIT.
// Source: mainroute.go
//
// Generated by this command:
//
//	mockgen -source=mainroute.go -destination=mock/mainroutemock.go
//

// Package mock_route is a generated GoMock package.
package mock_route

import (
	reflect "reflect"

	route "github.com/victorguarana/vehicle-routing/src/route"
	slc "github.com/victorguarana/vehicle-routing/src/slc"
	gomock "go.uber.org/mock/gomock"
)

// MockIMainRoute is a mock of IMainRoute interface.
type MockIMainRoute struct {
	ctrl     *gomock.Controller
	recorder *MockIMainRouteMockRecorder
}

// MockIMainRouteMockRecorder is the mock recorder for MockIMainRoute.
type MockIMainRouteMockRecorder struct {
	mock *MockIMainRoute
}

// NewMockIMainRoute creates a new mock instance.
func NewMockIMainRoute(ctrl *gomock.Controller) *MockIMainRoute {
	mock := &MockIMainRoute{ctrl: ctrl}
	mock.recorder = &MockIMainRouteMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMainRoute) EXPECT() *MockIMainRouteMockRecorder {
	return m.recorder
}

// Append mocks base method.
func (m *MockIMainRoute) Append(mainStop route.IMainStop) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Append", mainStop)
}

// Append indicates an expected call of Append.
func (mr *MockIMainRouteMockRecorder) Append(mainStop any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockIMainRoute)(nil).Append), mainStop)
}

// InserAt mocks base method.
func (m *MockIMainRoute) InserAt(index int, mainStop route.IMainStop) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InserAt", index, mainStop)
}

// InserAt indicates an expected call of InserAt.
func (mr *MockIMainRouteMockRecorder) InserAt(index, mainStop any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InserAt", reflect.TypeOf((*MockIMainRoute)(nil).InserAt), index, mainStop)
}

// Iterator mocks base method.
func (m *MockIMainRoute) Iterator() slc.Iterator[route.IMainStop] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator")
	ret0, _ := ret[0].(slc.Iterator[route.IMainStop])
	return ret0
}

// Iterator indicates an expected call of Iterator.
func (mr *MockIMainRouteMockRecorder) Iterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockIMainRoute)(nil).Iterator))
}

// Last mocks base method.
func (m *MockIMainRoute) Last() route.IMainStop {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Last")
	ret0, _ := ret[0].(route.IMainStop)
	return ret0
}

// Last indicates an expected call of Last.
func (mr *MockIMainRouteMockRecorder) Last() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockIMainRoute)(nil).Last))
}

// Length mocks base method.
func (m *MockIMainRoute) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockIMainRouteMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockIMainRoute)(nil).Length))
}

// RemoveMainStop mocks base method.
func (m *MockIMainRoute) RemoveMainStop(index int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveMainStop", index)
}

// RemoveMainStop indicates an expected call of RemoveMainStop.
func (mr *MockIMainRouteMockRecorder) RemoveMainStop(index any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMainStop", reflect.TypeOf((*MockIMainRoute)(nil).RemoveMainStop), index)
}
