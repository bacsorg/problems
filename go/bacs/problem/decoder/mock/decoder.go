// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/bacsorg/problems/go/bacs/problem/decoder (interfaces: HumanDecoder)

package mock_decoder

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of HumanDecoder interface
type MockHumanDecoder struct {
	ctrl     *gomock.Controller
	recorder *_MockHumanDecoderRecorder
}

// Recorder for MockHumanDecoder (not exported)
type _MockHumanDecoderRecorder struct {
	mock *MockHumanDecoder
}

func NewMockHumanDecoder(ctrl *gomock.Controller) *MockHumanDecoder {
	mock := &MockHumanDecoder{ctrl: ctrl}
	mock.recorder = &_MockHumanDecoderRecorder{mock}
	return mock
}

func (_m *MockHumanDecoder) EXPECT() *_MockHumanDecoderRecorder {
	return _m.recorder
}

func (_m *MockHumanDecoder) DecodeBase64ToText(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "DecodeBase64ToText", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHumanDecoderRecorder) DecodeBase64ToText(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DecodeBase64ToText", arg0)
}

func (_m *MockHumanDecoder) DecodeToText(_param0 []byte) (string, error) {
	ret := _m.ctrl.Call(_m, "DecodeToText", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockHumanDecoderRecorder) DecodeToText(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DecodeToText", arg0)
}
