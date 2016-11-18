// Automatically generated by MockGen. DO NOT EDIT!
// Source: etcd-downloader/provider (interfaces: RemoteReader)

package mock_provider

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of RemoteReader interface
type MockRemoteReader struct {
	ctrl     *gomock.Controller
	recorder *_MockRemoteReaderRecorder
}

// Recorder for MockRemoteReader (not exported)
type _MockRemoteReaderRecorder struct {
	mock *MockRemoteReader
}

func NewMockRemoteReader(ctrl *gomock.Controller) *MockRemoteReader {
	mock := &MockRemoteReader{ctrl: ctrl}
	mock.recorder = &_MockRemoteReaderRecorder{mock}
	return mock
}

func (_m *MockRemoteReader) EXPECT() *_MockRemoteReaderRecorder {
	return _m.recorder
}

func (_m *MockRemoteReader) GetAll(_param0 string, _param1 string, _param2 string, _param3 string) (map[string]interface{}, error) {
	ret := _m.ctrl.Call(_m, "GetAll", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRemoteReaderRecorder) GetAll(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAll", arg0, arg1, arg2, arg3)
}
