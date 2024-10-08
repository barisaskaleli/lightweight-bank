// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source=repository.go -destination=mocks/repository_mock.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	config "github.com/barisaskaleli/lightweight-bank/config"
	model "github.com/barisaskaleli/lightweight-bank/internal/resource/model"
	request "github.com/barisaskaleli/lightweight-bank/internal/resource/request"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockIRepository is a mock of IRepository interface.
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository.
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance.
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// AddTransaction mocks base method.
func (m *MockIRepository) AddTransaction(tx *gorm.DB, sender, receiver model.User, amount, fee float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransaction", tx, sender, receiver, amount, fee)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTransaction indicates an expected call of AddTransaction.
func (mr *MockIRepositoryMockRecorder) AddTransaction(tx, sender, receiver, amount, fee any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransaction", reflect.TypeOf((*MockIRepository)(nil).AddTransaction), tx, sender, receiver, amount, fee)
}

// GetByAccountNumber mocks base method.
func (m *MockIRepository) GetByAccountNumber(accountNumber string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAccountNumber", accountNumber)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAccountNumber indicates an expected call of GetByAccountNumber.
func (mr *MockIRepositoryMockRecorder) GetByAccountNumber(accountNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAccountNumber", reflect.TypeOf((*MockIRepository)(nil).GetByAccountNumber), accountNumber)
}

// GetDatabase mocks base method.
func (m *MockIRepository) GetDatabase() config.IMysqlInstance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabase")
	ret0, _ := ret[0].(config.IMysqlInstance)
	return ret0
}

// GetDatabase indicates an expected call of GetDatabase.
func (mr *MockIRepositoryMockRecorder) GetDatabase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabase", reflect.TypeOf((*MockIRepository)(nil).GetDatabase))
}

// Login mocks base method.
func (m *MockIRepository) Login(request request.LoginRequest) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", request)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockIRepositoryMockRecorder) Login(request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockIRepository)(nil).Login), request)
}

// Register mocks base method.
func (m *MockIRepository) Register(request request.RegisterRequest, accountNumber string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", request, accountNumber)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockIRepositoryMockRecorder) Register(request, accountNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockIRepository)(nil).Register), request, accountNumber)
}

// UpdateBalance mocks base method.
func (m *MockIRepository) UpdateBalance(tx *gorm.DB, accountNumber string, amount float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBalance", tx, accountNumber, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBalance indicates an expected call of UpdateBalance.
func (mr *MockIRepositoryMockRecorder) UpdateBalance(tx, accountNumber, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalance", reflect.TypeOf((*MockIRepository)(nil).UpdateBalance), tx, accountNumber, amount)
}
