// Code generated by protoc-gen-go-client-mock. DO NOT EDIT.
// versions:
// - protoc-gen-go-client-mock v0.0.1

package myapp

import (
	"github.com/golang/mock/gomock"
)

var _ Client = (*MockClient)(nil)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock               *MockClient
	mockProjectService *MockProjectServiceClient
	mockUserService    *MockUserServiceClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{
		mock:               mock,
		mockProjectService: NewMockProjectServiceClient(ctrl),
		mockUserService:    NewMockUserServiceClient(ctrl),
	}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

func (m *MockClient) ProjectService() ProjectServiceClient {
	m.ctrl.T.Helper()
	return m.recorder.mockProjectService
}

func (mr *MockClientMockRecorder) ProjectService() *MockProjectServiceClientMockRecorder {
	mr.mock.ctrl.T.Helper()
	return mr.mockProjectService.EXPECT()
}

func (m *MockClient) UserService() UserServiceClient {
	m.ctrl.T.Helper()
	return m.recorder.mockUserService
}

func (mr *MockClientMockRecorder) UserService() *MockUserServiceClientMockRecorder {
	mr.mock.ctrl.T.Helper()
	return mr.mockUserService.EXPECT()
}
