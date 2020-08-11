// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/libopenstorage/openstorage/api (interfaces: OpenStorageFilesystemCheckServer,OpenStorageFilesystemCheckClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/libopenstorage/openstorage/api"
	grpc "google.golang.org/grpc"
)

// MockOpenStorageFilesystemCheckServer is a mock of OpenStorageFilesystemCheckServer interface.
type MockOpenStorageFilesystemCheckServer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageFilesystemCheckServerMockRecorder
}

// MockOpenStorageFilesystemCheckServerMockRecorder is the mock recorder for MockOpenStorageFilesystemCheckServer.
type MockOpenStorageFilesystemCheckServerMockRecorder struct {
	mock *MockOpenStorageFilesystemCheckServer
}

// NewMockOpenStorageFilesystemCheckServer creates a new mock instance.
func NewMockOpenStorageFilesystemCheckServer(ctrl *gomock.Controller) *MockOpenStorageFilesystemCheckServer {
	mock := &MockOpenStorageFilesystemCheckServer{ctrl: ctrl}
	mock.recorder = &MockOpenStorageFilesystemCheckServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStorageFilesystemCheckServer) EXPECT() *MockOpenStorageFilesystemCheckServerMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockOpenStorageFilesystemCheckServer) Start(arg0 context.Context, arg1 *api.SdkFilesystemCheckStartRequest) (*api.SdkFilesystemCheckStartResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckStartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start.
func (mr *MockOpenStorageFilesystemCheckServerMockRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockOpenStorageFilesystemCheckServer)(nil).Start), arg0, arg1)
}

// Status mocks base method.
func (m *MockOpenStorageFilesystemCheckServer) Status(arg0 context.Context, arg1 *api.SdkFilesystemCheckStatusRequest) (*api.SdkFilesystemCheckStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockOpenStorageFilesystemCheckServerMockRecorder) Status(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockOpenStorageFilesystemCheckServer)(nil).Status), arg0, arg1)
}

// Stop mocks base method.
func (m *MockOpenStorageFilesystemCheckServer) Stop(arg0 context.Context, arg1 *api.SdkFilesystemCheckStopRequest) (*api.SdkFilesystemCheckStopResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckStopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stop indicates an expected call of Stop.
func (mr *MockOpenStorageFilesystemCheckServerMockRecorder) Stop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockOpenStorageFilesystemCheckServer)(nil).Stop), arg0, arg1)
}

// MockOpenStorageFilesystemCheckClient is a mock of OpenStorageFilesystemCheckClient interface.
type MockOpenStorageFilesystemCheckClient struct {
	ctrl     *gomock.Controller
	recorder *MockOpenStorageFilesystemCheckClientMockRecorder
}

// MockOpenStorageFilesystemCheckClientMockRecorder is the mock recorder for MockOpenStorageFilesystemCheckClient.
type MockOpenStorageFilesystemCheckClientMockRecorder struct {
	mock *MockOpenStorageFilesystemCheckClient
}

// NewMockOpenStorageFilesystemCheckClient creates a new mock instance.
func NewMockOpenStorageFilesystemCheckClient(ctrl *gomock.Controller) *MockOpenStorageFilesystemCheckClient {
	mock := &MockOpenStorageFilesystemCheckClient{ctrl: ctrl}
	mock.recorder = &MockOpenStorageFilesystemCheckClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenStorageFilesystemCheckClient) EXPECT() *MockOpenStorageFilesystemCheckClientMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockOpenStorageFilesystemCheckClient) Start(arg0 context.Context, arg1 *api.SdkFilesystemCheckStartRequest, arg2 ...grpc.CallOption) (*api.SdkFilesystemCheckStartResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Start", varargs...)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckStartResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Start indicates an expected call of Start.
func (mr *MockOpenStorageFilesystemCheckClientMockRecorder) Start(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockOpenStorageFilesystemCheckClient)(nil).Start), varargs...)
}

// Status mocks base method.
func (m *MockOpenStorageFilesystemCheckClient) Status(arg0 context.Context, arg1 *api.SdkFilesystemCheckStatusRequest, arg2 ...grpc.CallOption) (*api.SdkFilesystemCheckStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Status", varargs...)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockOpenStorageFilesystemCheckClientMockRecorder) Status(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockOpenStorageFilesystemCheckClient)(nil).Status), varargs...)
}

// Stop mocks base method.
func (m *MockOpenStorageFilesystemCheckClient) Stop(arg0 context.Context, arg1 *api.SdkFilesystemCheckStopRequest, arg2 ...grpc.CallOption) (*api.SdkFilesystemCheckStopResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Stop", varargs...)
	ret0, _ := ret[0].(*api.SdkFilesystemCheckStopResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stop indicates an expected call of Stop.
func (mr *MockOpenStorageFilesystemCheckClientMockRecorder) Stop(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockOpenStorageFilesystemCheckClient)(nil).Stop), varargs...)
}