// Code generated by MockGen. DO NOT EDIT.
// Source: internal/data/contracts/auth.go
//
// Generated by this command:
//
//	mockgen -source=internal/data/contracts/auth.go -destination=internal/infra/mocks/auth_repository.go -package=mockinfra
//
// Package mockinfra is a generated GoMock package.
package mockinfra

import (
	context "context"
	reflect "reflect"

	contracts "bitbucket.org/elevatt/sirius/internal/data/contracts"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthRepositoryContract is a mock of AuthRepositoryContract interface.
type MockAuthRepositoryContract struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryContractMockRecorder
}

// MockAuthRepositoryContractMockRecorder is the mock recorder for MockAuthRepositoryContract.
type MockAuthRepositoryContractMockRecorder struct {
	mock *MockAuthRepositoryContract
}

// NewMockAuthRepositoryContract creates a new mock instance.
func NewMockAuthRepositoryContract(ctrl *gomock.Controller) *MockAuthRepositoryContract {
	mock := &MockAuthRepositoryContract{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryContractMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepositoryContract) EXPECT() *MockAuthRepositoryContractMockRecorder {
	return m.recorder
}

// FirstAdmin mocks base method.
func (m *MockAuthRepositoryContract) FirstAdmin(ctx context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirstAdmin", ctx)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FirstAdmin indicates an expected call of FirstAdmin.
func (mr *MockAuthRepositoryContractMockRecorder) FirstAdmin(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirstAdmin", reflect.TypeOf((*MockAuthRepositoryContract)(nil).FirstAdmin), ctx)
}

// GetUserByID mocks base method.
func (m *MockAuthRepositoryContract) GetUserByID(ctx context.Context, id string) (*contracts.UserContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*contracts.UserContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockAuthRepositoryContractMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockAuthRepositoryContract)(nil).GetUserByID), ctx, id)
}

// SignInADMIN mocks base method.
func (m *MockAuthRepositoryContract) SignInADMIN(ctx context.Context, email, password string) (*contracts.UserContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignInADMIN", ctx, email, password)
	ret0, _ := ret[0].(*contracts.UserContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInADMIN indicates an expected call of SignInADMIN.
func (mr *MockAuthRepositoryContractMockRecorder) SignInADMIN(ctx, email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInADMIN", reflect.TypeOf((*MockAuthRepositoryContract)(nil).SignInADMIN), ctx, email, password)
}

// SignInCLIENT mocks base method.
func (m *MockAuthRepositoryContract) SignInCLIENT(ctx context.Context, email, password string) (*contracts.UserContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignInCLIENT", ctx, email, password)
	ret0, _ := ret[0].(*contracts.UserContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInCLIENT indicates an expected call of SignInCLIENT.
func (mr *MockAuthRepositoryContractMockRecorder) SignInCLIENT(ctx, email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInCLIENT", reflect.TypeOf((*MockAuthRepositoryContract)(nil).SignInCLIENT), ctx, email, password)
}

// SignInWORKER mocks base method.
func (m *MockAuthRepositoryContract) SignInWORKER(ctx context.Context, email, password string) (*contracts.UserContract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignInWORKER", ctx, email, password)
	ret0, _ := ret[0].(*contracts.UserContract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInWORKER indicates an expected call of SignInWORKER.
func (mr *MockAuthRepositoryContractMockRecorder) SignInWORKER(ctx, email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInWORKER", reflect.TypeOf((*MockAuthRepositoryContract)(nil).SignInWORKER), ctx, email, password)
}

// SignUpCLIENT mocks base method.
func (m *MockAuthRepositoryContract) SignUpCLIENT(ctx context.Context, input contracts.UserContract) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUpCLIENT", ctx, input)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUpCLIENT indicates an expected call of SignUpCLIENT.
func (mr *MockAuthRepositoryContractMockRecorder) SignUpCLIENT(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUpCLIENT", reflect.TypeOf((*MockAuthRepositoryContract)(nil).SignUpCLIENT), ctx, input)
}
