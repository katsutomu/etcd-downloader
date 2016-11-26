package provider

import (
	"fmt"
	"gonfigloader/mock"
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

func TestMarshalToml(t *testing.T) {
	var m = map[string]interface{}{
		"database": map[string]string{
			"key1": "val1",
			"key2": "val2",
		},
	}
	_, err := marshal("toml", m)
	if err != nil {
		t.Error("error:", err)
	}
}

func TestMarshalJson(t *testing.T) {
	var m = map[string]interface{}{
		"database": map[string]string{
			"key1": "val1",
			"key2": "val2",
		},
	}
	expect := `{"database":{"key1":"val1","key2":"val2"}}`
	b, err := marshal("json", m)
	actual := string(b)
	if err != nil {
		t.Error("error:", err)
	}
	if expect != actual {
		t.Error(expect + "\n" + actual)
	}
}
