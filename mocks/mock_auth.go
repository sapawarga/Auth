// Code generated by MockGen. DO NOT EDIT.
// Source: repository/repository_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/sapawarga/auth/model"
)

// MockAuthI is a mock of AuthI interface.
type MockAuthI struct {
	ctrl     *gomock.Controller
	recorder *MockAuthIMockRecorder
}

// MockAuthIMockRecorder is the mock recorder for MockAuthI.
type MockAuthIMockRecorder struct {
	mock *MockAuthI
}

// NewMockAuthI creates a new mock instance.
func NewMockAuthI(ctrl *gomock.Controller) *MockAuthI {
	mock := &MockAuthI{ctrl: ctrl}
	mock.recorder = &MockAuthIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthI) EXPECT() *MockAuthIMockRecorder {
	return m.recorder
}

// GetActorCurrentLoginByUsername mocks base method.
func (m *MockAuthI) GetActorCurrentLoginByUsername(ctx context.Context, username string) (*model.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorCurrentLoginByUsername", ctx, username)
	ret0, _ := ret[0].(*model.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorCurrentLoginByUsername indicates an expected call of GetActorCurrentLoginByUsername.
func (mr *MockAuthIMockRecorder) GetActorCurrentLoginByUsername(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorCurrentLoginByUsername", reflect.TypeOf((*MockAuthI)(nil).GetActorCurrentLoginByUsername), ctx, username)
}

// GetActorDetailByUsername mocks base method.
func (m *MockAuthI) GetActorDetailByUsername(ctx context.Context, username string) (*model.UserDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorDetailByUsername", ctx, username)
	ret0, _ := ret[0].(*model.UserDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorDetailByUsername indicates an expected call of GetActorDetailByUsername.
func (mr *MockAuthIMockRecorder) GetActorDetailByUsername(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorDetailByUsername", reflect.TypeOf((*MockAuthI)(nil).GetActorDetailByUsername), ctx, username)
}

// MockJWToken is a mock of JWToken interface.
type MockJWToken struct {
	ctrl     *gomock.Controller
	recorder *MockJWTokenMockRecorder
}

// MockJWTokenMockRecorder is the mock recorder for MockJWToken.
type MockJWTokenMockRecorder struct {
	mock *MockJWToken
}

// NewMockJWToken creates a new mock instance.
func NewMockJWToken(ctrl *gomock.Controller) *MockJWToken {
	mock := &MockJWToken{ctrl: ctrl}
	mock.recorder = &MockJWTokenMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJWToken) EXPECT() *MockJWTokenMockRecorder {
	return m.recorder
}

// ParsingToken mocks base method.
func (m *MockJWToken) ParsingToken(ctx context.Context, token string) (*model.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParsingToken", ctx, token)
	ret0, _ := ret[0].(*model.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParsingToken indicates an expected call of ParsingToken.
func (mr *MockJWTokenMockRecorder) ParsingToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParsingToken", reflect.TypeOf((*MockJWToken)(nil).ParsingToken), ctx, token)
}