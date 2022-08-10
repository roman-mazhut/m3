// Code generated by MockGen. DO NOT EDIT.
// Source: ../../namespace/kvadmin/types.go

// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package kvadmin is a generated GoMock package.
package kvadmin

import (
	"reflect"

	"github.com/m3db/m3/src/dbnode/generated/proto/namespace"

	"github.com/golang/mock/gomock"
)

// MockNamespaceMetadataAdminService is a mock of NamespaceMetadataAdminService interface.
type MockNamespaceMetadataAdminService struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceMetadataAdminServiceMockRecorder
}

// MockNamespaceMetadataAdminServiceMockRecorder is the mock recorder for MockNamespaceMetadataAdminService.
type MockNamespaceMetadataAdminServiceMockRecorder struct {
	mock *MockNamespaceMetadataAdminService
}

// NewMockNamespaceMetadataAdminService creates a new mock instance.
func NewMockNamespaceMetadataAdminService(ctrl *gomock.Controller) *MockNamespaceMetadataAdminService {
	mock := &MockNamespaceMetadataAdminService{ctrl: ctrl}
	mock.recorder = &MockNamespaceMetadataAdminServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNamespaceMetadataAdminService) EXPECT() *MockNamespaceMetadataAdminServiceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockNamespaceMetadataAdminService) Add(name string, options *namespace.NamespaceOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", name, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockNamespaceMetadataAdminServiceMockRecorder) Add(name, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockNamespaceMetadataAdminService)(nil).Add), name, options)
}

// Delete mocks base method.
func (m *MockNamespaceMetadataAdminService) Delete(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockNamespaceMetadataAdminServiceMockRecorder) Delete(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNamespaceMetadataAdminService)(nil).Delete), name)
}

// DeploySchema mocks base method.
func (m *MockNamespaceMetadataAdminService) DeploySchema(name, protoFileName, msgName string, protos map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeploySchema", name, protoFileName, msgName, protos)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeploySchema indicates an expected call of DeploySchema.
func (mr *MockNamespaceMetadataAdminServiceMockRecorder) DeploySchema(name, protoFileName, msgName, protos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeploySchema", reflect.TypeOf((*MockNamespaceMetadataAdminService)(nil).DeploySchema), name, protoFileName, msgName, protos)
}

// Get mocks base method.
func (m *MockNamespaceMetadataAdminService) Get(name string) (*namespace.NamespaceOptions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(*namespace.NamespaceOptions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockNamespaceMetadataAdminServiceMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNamespaceMetadataAdminService)(nil).Get), name)
}

// GetAll mocks base method.
func (m *MockNamespaceMetadataAdminService) GetAll() (*namespace.Registry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*namespace.Registry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockNamespaceMetadataAdminServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockNamespaceMetadataAdminService)(nil).GetAll))
}

// ResetSchema mocks base method.
func (m *MockNamespaceMetadataAdminService) ResetSchema(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetSchema", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetSchema indicates an expected call of ResetSchema.
func (mr *MockNamespaceMetadataAdminServiceMockRecorder) ResetSchema(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetSchema", reflect.TypeOf((*MockNamespaceMetadataAdminService)(nil).ResetSchema), name)
}

// Set mocks base method.
func (m *MockNamespaceMetadataAdminService) Set(name string, options *namespace.NamespaceOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", name, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockNamespaceMetadataAdminServiceMockRecorder) Set(name, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockNamespaceMetadataAdminService)(nil).Set), name, options)
}
