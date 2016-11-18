package provider

import (
	"bytes"
	"io"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var remoteReader RemoteReader

func init() {
	remoteReader = etcdReader{}
}

// GetRemoteReader
func GetRemoteReader(provider, endpoint, dir, ext string) (io.Reader, error) {
	remoteReader.getAll(provider, endpoint, dir, ext)
	return bytes.NewReader([]byte("Hello\nbyte.Reader\n")), nil
}

type RemoteReader interface {
	getAll(provider, endpoint, dir, ext string) (map[string]interface{}, error)
}

type etcdReader struct{}

func (r etcdReader) getAll(provider, endpoint, dir, ext string) (map[string]interface{}, error) {
	v := viper.New()
	v.AddRemoteProvider(provider, endpoint, dir)
	v.SetConfigType(ext)
	if err := v.ReadRemoteConfig(); err != nil {
		return nil, err
	}
	return v.AllSettings(), nil
}
