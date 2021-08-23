// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extension-provider-azure/pkg/azure/client (interfaces: DNSZone,DNSRecordSet,Factory)

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	client "github.com/gardener/gardener-extension-provider-azure/pkg/azure/client"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockDNSZone is a mock of DNSZone interface.
type MockDNSZone struct {
	ctrl     *gomock.Controller
	recorder *MockDNSZoneMockRecorder
}

// MockDNSZoneMockRecorder is the mock recorder for MockDNSZone.
type MockDNSZoneMockRecorder struct {
	mock *MockDNSZone
}

// NewMockDNSZone creates a new mock instance.
func NewMockDNSZone(ctrl *gomock.Controller) *MockDNSZone {
	mock := &MockDNSZone{ctrl: ctrl}
	mock.recorder = &MockDNSZoneMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDNSZone) EXPECT() *MockDNSZoneMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockDNSZone) GetAll(arg0 context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockDNSZoneMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDNSZone)(nil).GetAll), arg0)
}

// MockDNSRecordSet is a mock of DNSRecordSet interface.
type MockDNSRecordSet struct {
	ctrl     *gomock.Controller
	recorder *MockDNSRecordSetMockRecorder
}

// MockDNSRecordSetMockRecorder is the mock recorder for MockDNSRecordSet.
type MockDNSRecordSetMockRecorder struct {
	mock *MockDNSRecordSet
}

// NewMockDNSRecordSet creates a new mock instance.
func NewMockDNSRecordSet(ctrl *gomock.Controller) *MockDNSRecordSet {
	mock := &MockDNSRecordSet{ctrl: ctrl}
	mock.recorder = &MockDNSRecordSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDNSRecordSet) EXPECT() *MockDNSRecordSetMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockDNSRecordSet) CreateOrUpdate(arg0 context.Context, arg1, arg2, arg3 string, arg4 []string, arg5 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockDNSRecordSetMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockDNSRecordSet)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Delete mocks base method.
func (m *MockDNSRecordSet) Delete(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockDNSRecordSetMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDNSRecordSet)(nil).Delete), arg0, arg1, arg2, arg3)
}

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// DNSRecordSet mocks base method.
func (m *MockFactory) DNSRecordSet(arg0 context.Context, arg1 v1.SecretReference) (client.DNSRecordSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DNSRecordSet", arg0, arg1)
	ret0, _ := ret[0].(client.DNSRecordSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DNSRecordSet indicates an expected call of DNSRecordSet.
func (mr *MockFactoryMockRecorder) DNSRecordSet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DNSRecordSet", reflect.TypeOf((*MockFactory)(nil).DNSRecordSet), arg0, arg1)
}

// DNSZone mocks base method.
func (m *MockFactory) DNSZone(arg0 context.Context, arg1 v1.SecretReference) (client.DNSZone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DNSZone", arg0, arg1)
	ret0, _ := ret[0].(client.DNSZone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DNSZone indicates an expected call of DNSZone.
func (mr *MockFactoryMockRecorder) DNSZone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DNSZone", reflect.TypeOf((*MockFactory)(nil).DNSZone), arg0, arg1)
}

// Group mocks base method.
func (m *MockFactory) Group(arg0 context.Context, arg1 v1.SecretReference) (client.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group", arg0, arg1)
	ret0, _ := ret[0].(client.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Group indicates an expected call of Group.
func (mr *MockFactoryMockRecorder) Group(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockFactory)(nil).Group), arg0, arg1)
}

// Storage mocks base method.
func (m *MockFactory) Storage(arg0 context.Context, arg1 v1.SecretReference) (client.Storage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Storage", arg0, arg1)
	ret0, _ := ret[0].(client.Storage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Storage indicates an expected call of Storage.
func (mr *MockFactoryMockRecorder) Storage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockFactory)(nil).Storage), arg0, arg1)
}

// StorageAccount mocks base method.
func (m *MockFactory) StorageAccount(arg0 context.Context, arg1 v1.SecretReference) (client.StorageAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageAccount", arg0, arg1)
	ret0, _ := ret[0].(client.StorageAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageAccount indicates an expected call of StorageAccount.
func (mr *MockFactoryMockRecorder) StorageAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageAccount", reflect.TypeOf((*MockFactory)(nil).StorageAccount), arg0, arg1)
}

// Vmss mocks base method.
func (m *MockFactory) Vmss(arg0 context.Context, arg1 v1.SecretReference) (client.Vmss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vmss", arg0, arg1)
	ret0, _ := ret[0].(client.Vmss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Vmss indicates an expected call of Vmss.
func (mr *MockFactoryMockRecorder) Vmss(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vmss", reflect.TypeOf((*MockFactory)(nil).Vmss), arg0, arg1)
}