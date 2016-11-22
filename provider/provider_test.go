package provider

import (
	"etcd-downloader/mock"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetRemoteReader(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_provider.NewMockRemoteReader(ctrl)
	m := make(map[string]interface{})
	m["hoge"] = 1
	mock.EXPECT().Get("etcd", "hogehoge.com", "endpoint", "json").Return(m, nil)
	remoteReader = mock

	b, err := ReadRemoteConfig("etcd", "hogehoge.com", "endpoint", "json")
	if err != nil {
		t.Errorf("エラーが返却:%s", err)
		return
	}
	if `{"hoge":1}` != string(b) {
		t.Errorf("byte.Readerに変換できない\n%s", string(b))
	}
}
func TestShouldGetRemoteReaderReturnErrorOnGetAllError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock_provider.NewMockRemoteReader(ctrl)
	mock.EXPECT().Get("etcd", "hogehoge.com", "endpoint", "json").Return(nil, fmt.Errorf("test"))
	remoteReader = mock
	_, err := ReadRemoteConfig("etcd", "hogehoge.com", "endpoint", "json")
	if err == nil {
		t.Errorf("エラーが返却されなかった:%s", err)
	}
}
