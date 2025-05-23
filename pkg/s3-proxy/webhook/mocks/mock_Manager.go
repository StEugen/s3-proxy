// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/webhook (interfaces: Manager)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/mock_Manager.go -package=mocks github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/webhook Manager
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	webhook "github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/webhook"
	gomock "go.uber.org/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
	isgomock struct{}
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Load mocks base method.
func (m *MockManager) Load() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockManagerMockRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockManager)(nil).Load))
}

// ManageDELETEHooks mocks base method.
func (m *MockManager) ManageDELETEHooks(ctx context.Context, targetKey, requestPath string, s3Metadata *webhook.S3Metadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ManageDELETEHooks", ctx, targetKey, requestPath, s3Metadata)
}

// ManageDELETEHooks indicates an expected call of ManageDELETEHooks.
func (mr *MockManagerMockRecorder) ManageDELETEHooks(ctx, targetKey, requestPath, s3Metadata any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManageDELETEHooks", reflect.TypeOf((*MockManager)(nil).ManageDELETEHooks), ctx, targetKey, requestPath, s3Metadata)
}

// ManageGETHooks mocks base method.
func (m *MockManager) ManageGETHooks(ctx context.Context, targetKey, requestPath string, inputMetadata *webhook.GetInputMetadata, s3Metadata *webhook.S3Metadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ManageGETHooks", ctx, targetKey, requestPath, inputMetadata, s3Metadata)
}

// ManageGETHooks indicates an expected call of ManageGETHooks.
func (mr *MockManagerMockRecorder) ManageGETHooks(ctx, targetKey, requestPath, inputMetadata, s3Metadata any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManageGETHooks", reflect.TypeOf((*MockManager)(nil).ManageGETHooks), ctx, targetKey, requestPath, inputMetadata, s3Metadata)
}

// ManageHEADHooks mocks base method.
func (m *MockManager) ManageHEADHooks(ctx context.Context, targetKey, requestPath string, inputMetadata *webhook.HeadInputMetadata, s3Metadata *webhook.S3Metadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ManageHEADHooks", ctx, targetKey, requestPath, inputMetadata, s3Metadata)
}

// ManageHEADHooks indicates an expected call of ManageHEADHooks.
func (mr *MockManagerMockRecorder) ManageHEADHooks(ctx, targetKey, requestPath, inputMetadata, s3Metadata any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManageHEADHooks", reflect.TypeOf((*MockManager)(nil).ManageHEADHooks), ctx, targetKey, requestPath, inputMetadata, s3Metadata)
}

// ManagePUTHooks mocks base method.
func (m *MockManager) ManagePUTHooks(ctx context.Context, targetKey, requestPath string, inputMetadata *webhook.PutInputMetadata, s3Metadata *webhook.S3Metadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ManagePUTHooks", ctx, targetKey, requestPath, inputMetadata, s3Metadata)
}

// ManagePUTHooks indicates an expected call of ManagePUTHooks.
func (mr *MockManagerMockRecorder) ManagePUTHooks(ctx, targetKey, requestPath, inputMetadata, s3Metadata any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManagePUTHooks", reflect.TypeOf((*MockManager)(nil).ManagePUTHooks), ctx, targetKey, requestPath, inputMetadata, s3Metadata)
}
