// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	entity "payment_service/internal/entity"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPaymentRepository is a mock of PaymentRepository interface.
type MockPaymentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentRepositoryMockRecorder
	isgomock struct{}
}

// MockPaymentRepositoryMockRecorder is the mock recorder for MockPaymentRepository.
type MockPaymentRepositoryMockRecorder struct {
	mock *MockPaymentRepository
}

// NewMockPaymentRepository creates a new mock instance.
func NewMockPaymentRepository(ctrl *gomock.Controller) *MockPaymentRepository {
	mock := &MockPaymentRepository{ctrl: ctrl}
	mock.recorder = &MockPaymentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentRepository) EXPECT() *MockPaymentRepositoryMockRecorder {
	return m.recorder
}

// CheckAndDeductBalance mocks base method.
func (m *MockPaymentRepository) CheckAndDeductBalance(ctx context.Context, cardNum string, amount float64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAndDeductBalance", ctx, cardNum, amount)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAndDeductBalance indicates an expected call of CheckAndDeductBalance.
func (mr *MockPaymentRepositoryMockRecorder) CheckAndDeductBalance(ctx, cardNum, amount any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAndDeductBalance", reflect.TypeOf((*MockPaymentRepository)(nil).CheckAndDeductBalance), ctx, cardNum, amount)
}

// Close mocks base method.
func (m *MockPaymentRepository) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPaymentRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPaymentRepository)(nil).Close))
}

// CreatePayment mocks base method.
func (m *MockPaymentRepository) CreatePayment(ctx context.Context, payment *entity.Payment) (*entity.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePayment", ctx, payment)
	ret0, _ := ret[0].(*entity.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePayment indicates an expected call of CreatePayment.
func (mr *MockPaymentRepositoryMockRecorder) CreatePayment(ctx, payment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockPaymentRepository)(nil).CreatePayment), ctx, payment)
}

// DeletePayment mocks base method.
func (m *MockPaymentRepository) DeletePayment(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePayment", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePayment indicates an expected call of DeletePayment.
func (mr *MockPaymentRepositoryMockRecorder) DeletePayment(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePayment", reflect.TypeOf((*MockPaymentRepository)(nil).DeletePayment), ctx, id)
}

// GetPayment mocks base method.
func (m *MockPaymentRepository) GetPayment(ctx context.Context, id int64) (*entity.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayment", ctx, id)
	ret0, _ := ret[0].(*entity.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayment indicates an expected call of GetPayment.
func (mr *MockPaymentRepositoryMockRecorder) GetPayment(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayment", reflect.TypeOf((*MockPaymentRepository)(nil).GetPayment), ctx, id)
}

// UpdatePayment mocks base method.
func (m *MockPaymentRepository) UpdatePayment(ctx context.Context, payment *entity.Payment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePayment", ctx, payment)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePayment indicates an expected call of UpdatePayment.
func (mr *MockPaymentRepositoryMockRecorder) UpdatePayment(ctx, payment any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePayment", reflect.TypeOf((*MockPaymentRepository)(nil).UpdatePayment), ctx, payment)
}

// ValidateCard mocks base method.
func (m *MockPaymentRepository) ValidateCard(ctx context.Context, cardNum, cvv, expDate, name string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCard", ctx, cardNum, cvv, expDate, name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateCard indicates an expected call of ValidateCard.
func (mr *MockPaymentRepositoryMockRecorder) ValidateCard(ctx, cardNum, cvv, expDate, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCard", reflect.TypeOf((*MockPaymentRepository)(nil).ValidateCard), ctx, cardNum, cvv, expDate, name)
}
