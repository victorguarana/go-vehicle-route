// Code generated by MockGen. DO NOT EDIT.
// Source: src/routes/dronestop.go
//
// Generated by this command:
//
//	mockgen -source=src/routes/dronestop.go -destination=src/routes/mocks/dronestopmock.go -package=mockroutes
//

// Package mockroutes is a generated GoMock package.
package mockroutes

import (
	reflect "reflect"

	gps "github.com/victorguarana/go-vehicle-route/src/gps"
	routes "github.com/victorguarana/go-vehicle-route/src/routes"
	vehicles "github.com/victorguarana/go-vehicle-route/src/vehicles"
	gomock "go.uber.org/mock/gomock"
)

// MockIDroneStop is a mock of IDroneStop interface.
type MockIDroneStop struct {
	ctrl     *gomock.Controller
	recorder *MockIDroneStopMockRecorder
}

// MockIDroneStopMockRecorder is the mock recorder for MockIDroneStop.
type MockIDroneStopMockRecorder struct {
	mock *MockIDroneStop
}

// NewMockIDroneStop creates a new mock instance.
func NewMockIDroneStop(ctrl *gomock.Controller) *MockIDroneStop {
	mock := &MockIDroneStop{ctrl: ctrl}
	mock.recorder = &MockIDroneStopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDroneStop) EXPECT() *MockIDroneStopMockRecorder {
	return m.recorder
}

// Drone mocks base method.
func (m *MockIDroneStop) Drone() vehicles.IDrone {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Drone")
	ret0, _ := ret[0].(vehicles.IDrone)
	return ret0
}

// Drone indicates an expected call of Drone.
func (mr *MockIDroneStopMockRecorder) Drone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Drone", reflect.TypeOf((*MockIDroneStop)(nil).Drone))
}

// Flight mocks base method.
func (m *MockIDroneStop) Flight() routes.IFlight {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flight")
	ret0, _ := ret[0].(routes.IFlight)
	return ret0
}

// Flight indicates an expected call of Flight.
func (mr *MockIDroneStopMockRecorder) Flight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flight", reflect.TypeOf((*MockIDroneStop)(nil).Flight))
}

// Point mocks base method.
func (m *MockIDroneStop) Point() *gps.Point {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Point")
	ret0, _ := ret[0].(*gps.Point)
	return ret0
}

// Point indicates an expected call of Point.
func (mr *MockIDroneStopMockRecorder) Point() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Point", reflect.TypeOf((*MockIDroneStop)(nil).Point))
}
