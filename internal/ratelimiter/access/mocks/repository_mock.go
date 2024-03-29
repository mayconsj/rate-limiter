// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetAccessCount mocks base method.
func (m *MockRepository) GetAccessCount(ctx context.Context, key string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessCount", ctx, key)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessCount indicates an expected call of GetAccessCount.
func (mr *MockRepositoryMockRecorder) GetAccessCount(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessCount", reflect.TypeOf((*MockRepository)(nil).GetAccessCount), ctx, key)
}

// IncrementAccessCount mocks base method.
func (m *MockRepository) IncrementAccessCount(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementAccessCount", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementAccessCount indicates an expected call of IncrementAccessCount.
func (mr *MockRepositoryMockRecorder) IncrementAccessCount(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementAccessCount", reflect.TypeOf((*MockRepository)(nil).IncrementAccessCount), ctx, key)
}
