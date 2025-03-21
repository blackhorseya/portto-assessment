// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -destination=./mock_repository.go -package=entity -source=repository.go
//

// Package entity is a generated GoMock package.
package entity

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCoinRepository is a mock of CoinRepository interface.
type MockCoinRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCoinRepositoryMockRecorder
	isgomock struct{}
}

// MockCoinRepositoryMockRecorder is the mock recorder for MockCoinRepository.
type MockCoinRepositoryMockRecorder struct {
	mock *MockCoinRepository
}

// NewMockCoinRepository creates a new mock instance.
func NewMockCoinRepository(ctrl *gomock.Controller) *MockCoinRepository {
	mock := &MockCoinRepository{ctrl: ctrl}
	mock.recorder = &MockCoinRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoinRepository) EXPECT() *MockCoinRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCoinRepository) Create(c context.Context, coin *Coin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", c, coin)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCoinRepositoryMockRecorder) Create(c, coin any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCoinRepository)(nil).Create), c, coin)
}

// Delete mocks base method.
func (m *MockCoinRepository) Delete(c context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", c, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCoinRepositoryMockRecorder) Delete(c, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCoinRepository)(nil).Delete), c, id)
}

// GetByID mocks base method.
func (m *MockCoinRepository) GetByID(c context.Context, id uint) (*Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", c, id)
	ret0, _ := ret[0].(*Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockCoinRepositoryMockRecorder) GetByID(c, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockCoinRepository)(nil).GetByID), c, id)
}

// GetByName mocks base method.
func (m *MockCoinRepository) GetByName(c context.Context, name string) (*Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", c, name)
	ret0, _ := ret[0].(*Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockCoinRepositoryMockRecorder) GetByName(c, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockCoinRepository)(nil).GetByName), c, name)
}

// List mocks base method.
func (m *MockCoinRepository) List(c context.Context, cond ListCondition) ([]*Coin, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", c, cond)
	ret0, _ := ret[0].([]*Coin)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockCoinRepositoryMockRecorder) List(c, cond any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCoinRepository)(nil).List), c, cond)
}

// Poke mocks base method.
func (m *MockCoinRepository) Poke(c context.Context, id uint, score int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Poke", c, id, score)
	ret0, _ := ret[0].(error)
	return ret0
}

// Poke indicates an expected call of Poke.
func (mr *MockCoinRepositoryMockRecorder) Poke(c, id, score any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Poke", reflect.TypeOf((*MockCoinRepository)(nil).Poke), c, id, score)
}

// UpdateDescription mocks base method.
func (m *MockCoinRepository) UpdateDescription(c context.Context, id uint, description string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDescription", c, id, description)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDescription indicates an expected call of UpdateDescription.
func (mr *MockCoinRepositoryMockRecorder) UpdateDescription(c, id, description any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDescription", reflect.TypeOf((*MockCoinRepository)(nil).UpdateDescription), c, id, description)
}
