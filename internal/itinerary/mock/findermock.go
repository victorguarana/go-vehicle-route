// Code generated by MockGen. DO NOT EDIT.
// Source: finder.go
//
// Generated by this command:
//
//	mockgen -source=finder.go -destination=mock/findermock.go
//

// Package mock_itinerary is a generated GoMock package.
package mock_itinerary

import (
	reflect "reflect"

	itinerary "github.com/victorguarana/vehicle-routing/internal/itinerary"
	gomock "go.uber.org/mock/gomock"
)

// MockFinder is a mock of Finder interface.
type MockFinder struct {
	ctrl     *gomock.Controller
	recorder *MockFinderMockRecorder
}

// MockFinderMockRecorder is the mock recorder for MockFinder.
type MockFinderMockRecorder struct {
	mock *MockFinder
}

// NewMockFinder creates a new mock instance.
func NewMockFinder(ctrl *gomock.Controller) *MockFinder {
	mock := &MockFinder{ctrl: ctrl}
	mock.recorder = &MockFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFinder) EXPECT() *MockFinderMockRecorder {
	return m.recorder
}

// FindWorstDroneStop mocks base method.
func (m *MockFinder) FindWorstDroneStop() itinerary.DroneStopCost {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindWorstDroneStop")
	ret0, _ := ret[0].(itinerary.DroneStopCost)
	return ret0
}

// FindWorstDroneStop indicates an expected call of FindWorstDroneStop.
func (mr *MockFinderMockRecorder) FindWorstDroneStop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindWorstDroneStop", reflect.TypeOf((*MockFinder)(nil).FindWorstDroneStop))
}

// FindWorstSwappableCarStopsOrdered mocks base method.
func (m *MockFinder) FindWorstSwappableCarStopsOrdered() []itinerary.CarStopCost {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindWorstSwappableCarStopsOrdered")
	ret0, _ := ret[0].([]itinerary.CarStopCost)
	return ret0
}

// FindWorstSwappableCarStopsOrdered indicates an expected call of FindWorstSwappableCarStopsOrdered.
func (mr *MockFinderMockRecorder) FindWorstSwappableCarStopsOrdered() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindWorstSwappableCarStopsOrdered", reflect.TypeOf((*MockFinder)(nil).FindWorstSwappableCarStopsOrdered))
}