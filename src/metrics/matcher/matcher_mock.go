// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/metrics/matcher (interfaces: Matcher)

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

// Package matcher is a generated GoMock package.
package matcher

import (
	"reflect"

	"github.com/m3db/m3/src/metrics/aggregation"
	"github.com/m3db/m3/src/metrics/metric"
	"github.com/m3db/m3/src/metrics/metric/id"
	"github.com/m3db/m3/src/metrics/rules"
	"github.com/m3db/m3/src/metrics/rules/view"

	"github.com/golang/mock/gomock"
)

// MockMatcher is a mock of Matcher interface.
type MockMatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMatcherMockRecorder
}

// MockMatcherMockRecorder is the mock recorder for MockMatcher.
type MockMatcherMockRecorder struct {
	mock *MockMatcher
}

// NewMockMatcher creates a new mock instance.
func NewMockMatcher(ctrl *gomock.Controller) *MockMatcher {
	mock := &MockMatcher{ctrl: ctrl}
	mock.recorder = &MockMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMatcher) EXPECT() *MockMatcherMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMatcher) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockMatcherMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMatcher)(nil).Close))
}

// ForwardMatch mocks base method.
func (m *MockMatcher) ForwardMatch(arg0 id.ID, arg1, arg2 int64, arg3 rules.MatchOptions) (rules.MatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForwardMatch", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(rules.MatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForwardMatch indicates an expected call of ForwardMatch.
func (mr *MockMatcherMockRecorder) ForwardMatch(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForwardMatch", reflect.TypeOf((*MockMatcher)(nil).ForwardMatch), arg0, arg1, arg2, arg3)
}

// LatestRollupRules mocks base method.
func (m *MockMatcher) LatestRollupRules(arg0 []byte, arg1 int64) ([]view.RollupRule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestRollupRules", arg0, arg1)
	ret0, _ := ret[0].([]view.RollupRule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestRollupRules indicates an expected call of LatestRollupRules.
func (mr *MockMatcherMockRecorder) LatestRollupRules(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestRollupRules", reflect.TypeOf((*MockMatcher)(nil).LatestRollupRules), arg0, arg1)
}

// ReverseMatch mocks base method.
func (m *MockMatcher) ReverseMatch(arg0 id.ID, arg1, arg2 int64, arg3 metric.Type, arg4 aggregation.Type, arg5 bool, arg6 aggregation.TypesOptions) (rules.MatchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReverseMatch", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(rules.MatchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReverseMatch indicates an expected call of ReverseMatch.
func (mr *MockMatcherMockRecorder) ReverseMatch(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReverseMatch", reflect.TypeOf((*MockMatcher)(nil).ReverseMatch), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}
