// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
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

// Code generated by MockGen. DO NOT EDIT.
// Source: manager_selector.go

// Package visibility is a generated GoMock package.
package visibility

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	namespace "go.temporal.io/server/common/namespace"
	manager "go.temporal.io/server/common/persistence/visibility/manager"
)

// MockmanagerSelector is a mock of managerSelector interface.
type MockmanagerSelector struct {
	ctrl     *gomock.Controller
	recorder *MockmanagerSelectorMockRecorder
}

// MockmanagerSelectorMockRecorder is the mock recorder for MockmanagerSelector.
type MockmanagerSelectorMockRecorder struct {
	mock *MockmanagerSelector
}

// NewMockmanagerSelector creates a new mock instance.
func NewMockmanagerSelector(ctrl *gomock.Controller) *MockmanagerSelector {
	mock := &MockmanagerSelector{ctrl: ctrl}
	mock.recorder = &MockmanagerSelectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmanagerSelector) EXPECT() *MockmanagerSelectorMockRecorder {
	return m.recorder
}

// readManager mocks base method.
func (m *MockmanagerSelector) readManager(nsName namespace.Name) manager.VisibilityManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "readManager", nsName)
	ret0, _ := ret[0].(manager.VisibilityManager)
	return ret0
}

// readManager indicates an expected call of readManager.
func (mr *MockmanagerSelectorMockRecorder) readManager(nsName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "readManager", reflect.TypeOf((*MockmanagerSelector)(nil).readManager), nsName)
}

// writeManagers mocks base method.
func (m *MockmanagerSelector) writeManagers() ([]manager.VisibilityManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "writeManagers")
	ret0, _ := ret[0].([]manager.VisibilityManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// writeManagers indicates an expected call of writeManagers.
func (mr *MockmanagerSelectorMockRecorder) writeManagers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "writeManagers", reflect.TypeOf((*MockmanagerSelector)(nil).writeManagers))
}