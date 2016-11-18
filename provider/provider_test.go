package provider

import (
	"bytes"
	"testing"
)

type DummyReader struct {
}

func (d DummyReader) getAll(provider, endpoint, dir, ext string) (map[string]interface{}, error) {
	return make(map[string]interface{}), nil
}
func TestGetRemoteReader(t *testing.T) {
	remoteReader = DummyReader{}
	r, err := GetRemoteReader("etcd", "hogehoge.com", "endpoint", "json")
	_, ok := r.(*bytes.Reader)
	if err != nil {
		t.Errorf("エラーが返却:%s", err)
	}
	if !ok {
		t.Errorf("byte.Readerに変換できない")
	}
}
