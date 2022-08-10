// Code generated by MockGen. DO NOT EDIT.
// Source: ../../postings/types.go

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

// Package postings is a generated GoMock package.
package postings

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockList is a mock of List interface.
type MockList struct {
	ctrl     *gomock.Controller
	recorder *MockListMockRecorder
}

// MockListMockRecorder is the mock recorder for MockList.
type MockListMockRecorder struct {
	mock *MockList
}

// NewMockList creates a new mock instance.
func NewMockList(ctrl *gomock.Controller) *MockList {
	mock := &MockList{ctrl: ctrl}
	mock.recorder = &MockListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockList) EXPECT() *MockListMockRecorder {
	return m.recorder
}

// CloneAsMutable mocks base method.
func (m *MockList) CloneAsMutable() MutableList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloneAsMutable")
	ret0, _ := ret[0].(MutableList)
	return ret0
}

// CloneAsMutable indicates an expected call of CloneAsMutable.
func (mr *MockListMockRecorder) CloneAsMutable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneAsMutable", reflect.TypeOf((*MockList)(nil).CloneAsMutable))
}

// Contains mocks base method.
func (m *MockList) Contains(id ID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contains", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Contains indicates an expected call of Contains.
func (mr *MockListMockRecorder) Contains(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contains", reflect.TypeOf((*MockList)(nil).Contains), id)
}

// Difference mocks base method.
func (m *MockList) Difference(other List) (List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", other)
	ret0, _ := ret[0].(List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Difference indicates an expected call of Difference.
func (mr *MockListMockRecorder) Difference(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockList)(nil).Difference), other)
}

// Equal mocks base method.
func (m *MockList) Equal(other List) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", other)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockListMockRecorder) Equal(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockList)(nil).Equal), other)
}

// Intersect mocks base method.
func (m *MockList) Intersect(other List) (List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersect", other)
	ret0, _ := ret[0].(List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Intersect indicates an expected call of Intersect.
func (mr *MockListMockRecorder) Intersect(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersect", reflect.TypeOf((*MockList)(nil).Intersect), other)
}

// IsEmpty mocks base method.
func (m *MockList) IsEmpty() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmpty indicates an expected call of IsEmpty.
func (mr *MockListMockRecorder) IsEmpty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmpty", reflect.TypeOf((*MockList)(nil).IsEmpty))
}

// Iterator mocks base method.
func (m *MockList) Iterator() Iterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator")
	ret0, _ := ret[0].(Iterator)
	return ret0
}

// Iterator indicates an expected call of Iterator.
func (mr *MockListMockRecorder) Iterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockList)(nil).Iterator))
}

// Len mocks base method.
func (m *MockList) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len.
func (mr *MockListMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockList)(nil).Len))
}

// Max mocks base method.
func (m *MockList) Max() (ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Max")
	ret0, _ := ret[0].(ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Max indicates an expected call of Max.
func (mr *MockListMockRecorder) Max() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Max", reflect.TypeOf((*MockList)(nil).Max))
}

// MockMutableList is a mock of MutableList interface.
type MockMutableList struct {
	ctrl     *gomock.Controller
	recorder *MockMutableListMockRecorder
}

// MockMutableListMockRecorder is the mock recorder for MockMutableList.
type MockMutableListMockRecorder struct {
	mock *MockMutableList
}

// NewMockMutableList creates a new mock instance.
func NewMockMutableList(ctrl *gomock.Controller) *MockMutableList {
	mock := &MockMutableList{ctrl: ctrl}
	mock.recorder = &MockMutableListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMutableList) EXPECT() *MockMutableListMockRecorder {
	return m.recorder
}

// AddIterator mocks base method.
func (m *MockMutableList) AddIterator(iter Iterator) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIterator", iter)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddIterator indicates an expected call of AddIterator.
func (mr *MockMutableListMockRecorder) AddIterator(iter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIterator", reflect.TypeOf((*MockMutableList)(nil).AddIterator), iter)
}

// AddRange mocks base method.
func (m *MockMutableList) AddRange(min, max ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRange", min, max)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRange indicates an expected call of AddRange.
func (mr *MockMutableListMockRecorder) AddRange(min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRange", reflect.TypeOf((*MockMutableList)(nil).AddRange), min, max)
}

// CloneAsMutable mocks base method.
func (m *MockMutableList) CloneAsMutable() MutableList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloneAsMutable")
	ret0, _ := ret[0].(MutableList)
	return ret0
}

// CloneAsMutable indicates an expected call of CloneAsMutable.
func (mr *MockMutableListMockRecorder) CloneAsMutable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneAsMutable", reflect.TypeOf((*MockMutableList)(nil).CloneAsMutable))
}

// Contains mocks base method.
func (m *MockMutableList) Contains(id ID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contains", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Contains indicates an expected call of Contains.
func (mr *MockMutableListMockRecorder) Contains(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contains", reflect.TypeOf((*MockMutableList)(nil).Contains), id)
}

// Difference mocks base method.
func (m *MockMutableList) Difference(other List) (List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Difference", other)
	ret0, _ := ret[0].(List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Difference indicates an expected call of Difference.
func (mr *MockMutableListMockRecorder) Difference(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Difference", reflect.TypeOf((*MockMutableList)(nil).Difference), other)
}

// Equal mocks base method.
func (m *MockMutableList) Equal(other List) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Equal", other)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Equal indicates an expected call of Equal.
func (mr *MockMutableListMockRecorder) Equal(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Equal", reflect.TypeOf((*MockMutableList)(nil).Equal), other)
}

// Insert mocks base method.
func (m *MockMutableList) Insert(i ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", i)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockMutableListMockRecorder) Insert(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockMutableList)(nil).Insert), i)
}

// Intersect mocks base method.
func (m *MockMutableList) Intersect(other List) (List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Intersect", other)
	ret0, _ := ret[0].(List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Intersect indicates an expected call of Intersect.
func (mr *MockMutableListMockRecorder) Intersect(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Intersect", reflect.TypeOf((*MockMutableList)(nil).Intersect), other)
}

// IsEmpty mocks base method.
func (m *MockMutableList) IsEmpty() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmpty indicates an expected call of IsEmpty.
func (mr *MockMutableListMockRecorder) IsEmpty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmpty", reflect.TypeOf((*MockMutableList)(nil).IsEmpty))
}

// Iterator mocks base method.
func (m *MockMutableList) Iterator() Iterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator")
	ret0, _ := ret[0].(Iterator)
	return ret0
}

// Iterator indicates an expected call of Iterator.
func (mr *MockMutableListMockRecorder) Iterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockMutableList)(nil).Iterator))
}

// Len mocks base method.
func (m *MockMutableList) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len.
func (mr *MockMutableListMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockMutableList)(nil).Len))
}

// Max mocks base method.
func (m *MockMutableList) Max() (ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Max")
	ret0, _ := ret[0].(ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Max indicates an expected call of Max.
func (mr *MockMutableListMockRecorder) Max() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Max", reflect.TypeOf((*MockMutableList)(nil).Max))
}

// RemoveRange mocks base method.
func (m *MockMutableList) RemoveRange(min, max ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRange", min, max)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRange indicates an expected call of RemoveRange.
func (mr *MockMutableListMockRecorder) RemoveRange(min, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRange", reflect.TypeOf((*MockMutableList)(nil).RemoveRange), min, max)
}

// Reset mocks base method.
func (m *MockMutableList) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockMutableListMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockMutableList)(nil).Reset))
}

// UnionInPlace mocks base method.
func (m *MockMutableList) UnionInPlace(other List) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnionInPlace", other)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnionInPlace indicates an expected call of UnionInPlace.
func (mr *MockMutableListMockRecorder) UnionInPlace(other interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnionInPlace", reflect.TypeOf((*MockMutableList)(nil).UnionInPlace), other)
}

// UnionManyInPlace mocks base method.
func (m *MockMutableList) UnionManyInPlace(others []List) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnionManyInPlace", others)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnionManyInPlace indicates an expected call of UnionManyInPlace.
func (mr *MockMutableListMockRecorder) UnionManyInPlace(others interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnionManyInPlace", reflect.TypeOf((*MockMutableList)(nil).UnionManyInPlace), others)
}

// MockIterator is a mock of Iterator interface.
type MockIterator struct {
	ctrl     *gomock.Controller
	recorder *MockIteratorMockRecorder
}

// MockIteratorMockRecorder is the mock recorder for MockIterator.
type MockIteratorMockRecorder struct {
	mock *MockIterator
}

// NewMockIterator creates a new mock instance.
func NewMockIterator(ctrl *gomock.Controller) *MockIterator {
	mock := &MockIterator{ctrl: ctrl}
	mock.recorder = &MockIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIterator) EXPECT() *MockIteratorMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIterator) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIteratorMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIterator)(nil).Close))
}

// Current mocks base method.
func (m *MockIterator) Current() ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(ID)
	return ret0
}

// Current indicates an expected call of Current.
func (mr *MockIteratorMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockIterator)(nil).Current))
}

// Err mocks base method.
func (m *MockIterator) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockIteratorMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockIterator)(nil).Err))
}

// Next mocks base method.
func (m *MockIterator) Next() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIterator)(nil).Next))
}

// MockPool is a mock of Pool interface.
type MockPool struct {
	ctrl     *gomock.Controller
	recorder *MockPoolMockRecorder
}

// MockPoolMockRecorder is the mock recorder for MockPool.
type MockPoolMockRecorder struct {
	mock *MockPool
}

// NewMockPool creates a new mock instance.
func NewMockPool(ctrl *gomock.Controller) *MockPool {
	mock := &MockPool{ctrl: ctrl}
	mock.recorder = &MockPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPool) EXPECT() *MockPoolMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockPool) Get() MutableList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(MutableList)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockPoolMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPool)(nil).Get))
}

// Put mocks base method.
func (m *MockPool) Put(pl MutableList) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", pl)
}

// Put indicates an expected call of Put.
func (mr *MockPoolMockRecorder) Put(pl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockPool)(nil).Put), pl)
}
