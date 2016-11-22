package provider

import (
	"encoding/json"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var remoteReader RemoteReader

func init() {
	remoteReader = etcdReader{}
}

// GetRemoteReader
func ReadRemoteConfig(provider, endpoint, dir, ext string) ([]byte, error) {
	config, err := remoteReader.Get(provider, endpoint, dir, ext)
	if err != nil {
		return nil, err
	}
	s, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func ReadRemoteSecureConfig(provider, endpoint, dir, key, ext string) ([]byte, error) {
	config, err := remoteReader.GetSecure(provider, endpoint, dir, key, ext)
	if err != nil {
		return nil, err
	}
	s, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

type RemoteReader interface {
	Get(provider, endpoint, dir, ext string) (map[string]interface{}, error)
	GetSecure(provider, endpoint, dir, key, ext string) (map[string]interface{}, error)
}

type etcdReader struct{}

func (r etcdReader) Get(provider, endpoint, dir, ext string) (map[string]interface{}, error) {
	v := viper.New()
	v.AddRemoteProvider(provider, endpoint, dir)
	v.SetConfigType(ext)
	if err := v.ReadRemoteConfig(); err != nil {
		return nil, err
	}
	return v.AllSettings(), nil
}

func (r etcdReader) GetSecure(provider, endpoint, dir, key, ext string) (map[string]interface{}, error) {
	v := viper.New()
	v.AddSecureRemoteProvider(provider, endpoint, dir, key)
	v.SetConfigType(ext)
	if err := v.ReadRemoteConfig(); err != nil {
		return nil, err
	}
	return v.AllSettings(), nil
}
