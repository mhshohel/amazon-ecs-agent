// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-agent/agent/eni/statemanager (interfaces: StateManager)

package mock_statemanager

import (
	gomock "github.com/golang/mock/gomock"
	netlink "github.com/vishvananda/netlink"
)

// Mock of StateManager interface
type MockStateManager struct {
	ctrl     *gomock.Controller
	recorder *_MockStateManagerRecorder
}

// Recorder for MockStateManager (not exported)
type _MockStateManagerRecorder struct {
	mock *MockStateManager
}

func NewMockStateManager(ctrl *gomock.Controller) *MockStateManager {
	mock := &MockStateManager{ctrl: ctrl}
	mock.recorder = &_MockStateManagerRecorder{mock}
	return mock
}

func (_m *MockStateManager) EXPECT() *_MockStateManagerRecorder {
	return _m.recorder
}

func (_m *MockStateManager) HandleENIEvent(_param0 string) {
	_m.ctrl.Call(_m, "HandleENIEvent", _param0)
}

func (_mr *_MockStateManagerRecorder) HandleENIEvent(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HandleENIEvent", arg0)
}

func (_m *MockStateManager) Init(_param0 []netlink.Link) {
	_m.ctrl.Call(_m, "Init", _param0)
}

func (_mr *_MockStateManagerRecorder) Init(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Init", arg0)
}

func (_m *MockStateManager) Reconcile(_param0 map[string]string) {
	_m.ctrl.Call(_m, "Reconcile", _param0)
}

func (_mr *_MockStateManagerRecorder) Reconcile(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reconcile", arg0)
}
