// Code generated by MockGen. DO NOT EDIT.
// Source: client_event_handler_generator_service.go

// Package mock_generators is a generated GoMock package.
package mock_generators

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockClientEventHandlerGeneratorService is a mock of ClientEventHandlerGeneratorService interface.
type MockClientEventHandlerGeneratorService struct {
	ctrl     *gomock.Controller
	recorder *MockClientEventHandlerGeneratorServiceMockRecorder
}

// MockClientEventHandlerGeneratorServiceMockRecorder is the mock recorder for MockClientEventHandlerGeneratorService.
type MockClientEventHandlerGeneratorServiceMockRecorder struct {
	mock *MockClientEventHandlerGeneratorService
}

// NewMockClientEventHandlerGeneratorService creates a new mock instance.
func NewMockClientEventHandlerGeneratorService(ctrl *gomock.Controller) *MockClientEventHandlerGeneratorService {
	mock := &MockClientEventHandlerGeneratorService{ctrl: ctrl}
	mock.recorder = &MockClientEventHandlerGeneratorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientEventHandlerGeneratorService) EXPECT() *MockClientEventHandlerGeneratorServiceMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockClientEventHandlerGeneratorService) Generate(packageName, eventType, outputPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", packageName, eventType, outputPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// Generate indicates an expected call of Generate.
func (mr *MockClientEventHandlerGeneratorServiceMockRecorder) Generate(packageName, eventType, outputPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockClientEventHandlerGeneratorService)(nil).Generate), packageName, eventType, outputPath)
}