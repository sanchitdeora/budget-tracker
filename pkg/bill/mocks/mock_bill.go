// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sanchitdeora/budget-tracker/pkg/bill (interfaces: Service)

// Package mock_bill is a generated GoMock package.
package mock_bill

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

// CreateBill mocks base method.
func (m *MockService) CreateBill(arg0 context.Context, arg1 models.Bill) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBill", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBill indicates an expected call of CreateBill.
func (mr *MockServiceMockRecorder) CreateBill(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBill", reflect.TypeOf((*MockService)(nil).CreateBill), arg0, arg1)
}

// CreateBillByUser mocks base method.
func (m *MockService) CreateBillByUser(arg0 context.Context, arg1 models.Bill) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBillByUser", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBillByUser indicates an expected call of CreateBillByUser.
func (mr *MockServiceMockRecorder) CreateBillByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBillByUser", reflect.TypeOf((*MockService)(nil).CreateBillByUser), arg0, arg1)
}

// DeleteBillById mocks base method.
func (m *MockService) DeleteBillById(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBillById", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteBillById indicates an expected call of DeleteBillById.
func (mr *MockServiceMockRecorder) DeleteBillById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBillById", reflect.TypeOf((*MockService)(nil).DeleteBillById), arg0, arg1)
}

// GetBillById mocks base method.
func (m *MockService) GetBillById(arg0 context.Context, arg1 string, arg2 *models.Bill) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillById", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetBillById indicates an expected call of GetBillById.
func (mr *MockServiceMockRecorder) GetBillById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillById", reflect.TypeOf((*MockService)(nil).GetBillById), arg0, arg1, arg2)
}

// GetBills mocks base method.
func (m *MockService) GetBills(arg0 context.Context, arg1 *[]models.Bill) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBills", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetBills indicates an expected call of GetBills.
func (mr *MockServiceMockRecorder) GetBills(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBills", reflect.TypeOf((*MockService)(nil).GetBills), arg0, arg1)
}

// UpdateBillById mocks base method.
func (m *MockService) UpdateBillById(arg0 context.Context, arg1 string, arg2 models.Bill) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBillById", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBillById indicates an expected call of UpdateBillById.
func (mr *MockServiceMockRecorder) UpdateBillById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBillById", reflect.TypeOf((*MockService)(nil).UpdateBillById), arg0, arg1, arg2)
}

// UpdateBillIsPaid mocks base method.
func (m *MockService) UpdateBillIsPaid(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBillIsPaid", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBillIsPaid indicates an expected call of UpdateBillIsPaid.
func (mr *MockServiceMockRecorder) UpdateBillIsPaid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBillIsPaid", reflect.TypeOf((*MockService)(nil).UpdateBillIsPaid), arg0, arg1)
}

// UpdateBillIsUnpaid mocks base method.
func (m *MockService) UpdateBillIsUnpaid(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBillIsUnpaid", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBillIsUnpaid indicates an expected call of UpdateBillIsUnpaid.
func (mr *MockServiceMockRecorder) UpdateBillIsUnpaid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBillIsUnpaid", reflect.TypeOf((*MockService)(nil).UpdateBillIsUnpaid), arg0, arg1)
}
