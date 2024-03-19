// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CanAccess mocks base method.
func (m *MockService) CanAccess(ctx context.Context, token, ip string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanAccess", ctx, token, ip)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanAccess indicates an expected call of CanAccess.
func (mr *MockServiceMockRecorder) CanAccess(ctx, token, ip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanAccess", reflect.TypeOf((*MockService)(nil).CanAccess), ctx, token, ip)
}