// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sanchitdeora/budget-tracker/pkg/goal (interfaces: Service)

// Package mock_goal is a generated GoMock package.
package mock_goal

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/sanchitdeora/budget-tracker/models"
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

// CreateGoalById mocks base method.
func (m *MockService) CreateGoalById(arg0 context.Context, arg1 models.Goal) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGoalById", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGoalById indicates an expected call of CreateGoalById.
func (mr *MockServiceMockRecorder) CreateGoalById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGoalById", reflect.TypeOf((*MockService)(nil).CreateGoalById), arg0, arg1)
}

// DeleteGoalById mocks base method.
func (m *MockService) DeleteGoalById(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGoalById", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGoalById indicates an expected call of DeleteGoalById.
func (mr *MockServiceMockRecorder) DeleteGoalById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGoalById", reflect.TypeOf((*MockService)(nil).DeleteGoalById), arg0, arg1)
}

// GetGoalById mocks base method.
func (m *MockService) GetGoalById(arg0 context.Context, arg1 string) (*models.Goal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGoalById", arg0, arg1)
	ret0, _ := ret[0].(*models.Goal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGoalById indicates an expected call of GetGoalById.
func (mr *MockServiceMockRecorder) GetGoalById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGoalById", reflect.TypeOf((*MockService)(nil).GetGoalById), arg0, arg1)
}

// GetGoals mocks base method.
func (m *MockService) GetGoals(arg0 context.Context, arg1 *[]models.Goal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGoals", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetGoals indicates an expected call of GetGoals.
func (mr *MockServiceMockRecorder) GetGoals(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGoals", reflect.TypeOf((*MockService)(nil).GetGoals), arg0, arg1)
}

// RemoveBudgetIdFromGoal mocks base method.
func (m *MockService) RemoveBudgetIdFromGoal(arg0 context.Context, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBudgetIdFromGoal", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveBudgetIdFromGoal indicates an expected call of RemoveBudgetIdFromGoal.
func (mr *MockServiceMockRecorder) RemoveBudgetIdFromGoal(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBudgetIdFromGoal", reflect.TypeOf((*MockService)(nil).RemoveBudgetIdFromGoal), arg0, arg1, arg2)
}

// UpdateBudgetIdsList mocks base method.
func (m *MockService) UpdateBudgetIdsList(arg0 context.Context, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBudgetIdsList", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBudgetIdsList indicates an expected call of UpdateBudgetIdsList.
func (mr *MockServiceMockRecorder) UpdateBudgetIdsList(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBudgetIdsList", reflect.TypeOf((*MockService)(nil).UpdateBudgetIdsList), arg0, arg1, arg2)
}

// UpdateGoalById mocks base method.
func (m *MockService) UpdateGoalById(arg0 context.Context, arg1 string, arg2 models.Goal) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGoalById", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGoalById indicates an expected call of UpdateGoalById.
func (mr *MockServiceMockRecorder) UpdateGoalById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGoalById", reflect.TypeOf((*MockService)(nil).UpdateGoalById), arg0, arg1, arg2)
}
