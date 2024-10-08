// Code generated by MockGen. DO NOT EDIT.
// Source: manager.go
//
// Generated by this command:
//
//	mockgen -source=manager.go -destination=mocks/manager_mock.go
//

// Package mock_config is a generated GoMock package.
package mock_config

import (
	reflect "reflect"

	config "github.com/barisaskaleli/lightweight-bank/config"
	gomock "go.uber.org/mock/gomock"
)

// MockIConfig is a mock of IConfig interface.
type MockIConfig struct {
	ctrl     *gomock.Controller
	recorder *MockIConfigMockRecorder
}

// MockIConfigMockRecorder is the mock recorder for MockIConfig.
type MockIConfigMockRecorder struct {
	mock *MockIConfig
}

// NewMockIConfig creates a new mock instance.
func NewMockIConfig(ctrl *gomock.Controller) *MockIConfig {
	mock := &MockIConfig{ctrl: ctrl}
	mock.recorder = &MockIConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIConfig) EXPECT() *MockIConfigMockRecorder {
	return m.recorder
}

// DB mocks base method.
func (m *MockIConfig) DB() config.DBConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(config.DBConfig)
	return ret0
}

// DB indicates an expected call of DB.
func (mr *MockIConfigMockRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockIConfig)(nil).DB))
}

// Server mocks base method.
func (m *MockIConfig) Server() config.ServerConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Server")
	ret0, _ := ret[0].(config.ServerConfig)
	return ret0
}

// Server indicates an expected call of Server.
func (mr *MockIConfigMockRecorder) Server() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Server", reflect.TypeOf((*MockIConfig)(nil).Server))
}

// Service mocks base method.
func (m *MockIConfig) Service() config.ServiceConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Service")
	ret0, _ := ret[0].(config.ServiceConfig)
	return ret0
}

// Service indicates an expected call of Service.
func (mr *MockIConfigMockRecorder) Service() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Service", reflect.TypeOf((*MockIConfig)(nil).Service))
}
