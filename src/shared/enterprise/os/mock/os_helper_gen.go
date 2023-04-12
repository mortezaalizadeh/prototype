// Code generated by MockGen. DO NOT EDIT.
// Source: os_helper.go

// Package mock_os is a generated GoMock package.
package mock_os

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOsHelper is a mock of OsHelper interface.
type MockOsHelper struct {
	ctrl     *gomock.Controller
	recorder *MockOsHelperMockRecorder
}

// MockOsHelperMockRecorder is the mock recorder for MockOsHelper.
type MockOsHelperMockRecorder struct {
	mock *MockOsHelper
}

// NewMockOsHelper creates a new mock instance.
func NewMockOsHelper(ctrl *gomock.Controller) *MockOsHelper {
	mock := &MockOsHelper{ctrl: ctrl}
	mock.recorder = &MockOsHelperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOsHelper) EXPECT() *MockOsHelperMockRecorder {
	return m.recorder
}

// CopyFile mocks base method.
func (m *MockOsHelper) CopyFile(sourcePath, destinationPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyFile", sourcePath, destinationPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyFile indicates an expected call of CopyFile.
func (mr *MockOsHelperMockRecorder) CopyFile(sourcePath, destinationPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyFile", reflect.TypeOf((*MockOsHelper)(nil).CopyFile), sourcePath, destinationPath)
}

// CreateBinaryFile mocks base method.
func (m *MockOsHelper) CreateBinaryFile(path string, content []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBinaryFile", path, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBinaryFile indicates an expected call of CreateBinaryFile.
func (mr *MockOsHelperMockRecorder) CreateBinaryFile(path, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBinaryFile", reflect.TypeOf((*MockOsHelper)(nil).CreateBinaryFile), path, content)
}

// CreateDir mocks base method.
func (m *MockOsHelper) CreateDir(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDir", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDir indicates an expected call of CreateDir.
func (mr *MockOsHelperMockRecorder) CreateDir(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDir", reflect.TypeOf((*MockOsHelper)(nil).CreateDir), path)
}

// CreateTemporaryTextFile mocks base method.
func (m *MockOsHelper) CreateTemporaryTextFile(content string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTemporaryTextFile", content)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTemporaryTextFile indicates an expected call of CreateTemporaryTextFile.
func (mr *MockOsHelperMockRecorder) CreateTemporaryTextFile(content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTemporaryTextFile", reflect.TypeOf((*MockOsHelper)(nil).CreateTemporaryTextFile), content)
}

// DirExist mocks base method.
func (m *MockOsHelper) DirExist(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DirExist", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DirExist indicates an expected call of DirExist.
func (mr *MockOsHelperMockRecorder) DirExist(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DirExist", reflect.TypeOf((*MockOsHelper)(nil).DirExist), path)
}

// FileExist mocks base method.
func (m *MockOsHelper) FileExist(path string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileExist", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// FileExist indicates an expected call of FileExist.
func (mr *MockOsHelperMockRecorder) FileExist(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileExist", reflect.TypeOf((*MockOsHelper)(nil).FileExist), path)
}

// GetEnvironmentVariable mocks base method.
func (m *MockOsHelper) GetEnvironmentVariable(key string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvironmentVariable", key)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEnvironmentVariable indicates an expected call of GetEnvironmentVariable.
func (mr *MockOsHelperMockRecorder) GetEnvironmentVariable(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironmentVariable", reflect.TypeOf((*MockOsHelper)(nil).GetEnvironmentVariable), key)
}

// GetFileAsByteArray mocks base method.
func (m *MockOsHelper) GetFileAsByteArray(path string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileAsByteArray", path)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileAsByteArray indicates an expected call of GetFileAsByteArray.
func (mr *MockOsHelperMockRecorder) GetFileAsByteArray(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileAsByteArray", reflect.TypeOf((*MockOsHelper)(nil).GetFileAsByteArray), path)
}

// GetFileAsString mocks base method.
func (m *MockOsHelper) GetFileAsString(path string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileAsString", path)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileAsString indicates an expected call of GetFileAsString.
func (mr *MockOsHelperMockRecorder) GetFileAsString(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileAsString", reflect.TypeOf((*MockOsHelper)(nil).GetFileAsString), path)
}
