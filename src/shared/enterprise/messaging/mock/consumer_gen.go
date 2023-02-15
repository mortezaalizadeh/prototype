// Code generated by MockGen. DO NOT EDIT.
// Source: consumer.go

// Package mock_messaging is a generated GoMock package.
package mock_messaging

import (
	context "context"
	reflect "reflect"

	messaging "github.com/Connectilly/connectilly/src/shared/enterprise/messaging"
	gomock "github.com/golang/mock/gomock"
)

// MockMessageConsumer is a mock of MessageConsumer interface.
type MockMessageConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockMessageConsumerMockRecorder
}

// MockMessageConsumerMockRecorder is the mock recorder for MockMessageConsumer.
type MockMessageConsumerMockRecorder struct {
	mock *MockMessageConsumer
}

// NewMockMessageConsumer creates a new mock instance.
func NewMockMessageConsumer(ctrl *gomock.Controller) *MockMessageConsumer {
	mock := &MockMessageConsumer{ctrl: ctrl}
	mock.recorder = &MockMessageConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageConsumer) EXPECT() *MockMessageConsumerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMessageConsumer) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockMessageConsumerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMessageConsumer)(nil).Close))
}

// Connect mocks base method.
func (m *MockMessageConsumer) Connect(topic string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", topic)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockMessageConsumerMockRecorder) Connect(topic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockMessageConsumer)(nil).Connect), topic)
}

// Consume mocks base method.
func (m *MockMessageConsumer) Consume(ctx context.Context) (messaging.Message, messaging.MessageProcessedCallback, messaging.MessageFailedCallback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", ctx)
	ret0, _ := ret[0].(messaging.Message)
	ret1, _ := ret[1].(messaging.MessageProcessedCallback)
	ret2, _ := ret[2].(messaging.MessageFailedCallback)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Consume indicates an expected call of Consume.
func (mr *MockMessageConsumerMockRecorder) Consume(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockMessageConsumer)(nil).Consume), ctx)
}
