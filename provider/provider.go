package provider

import (
	"bytes"
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var remoteReader RemoteReader

func init() {
	remoteReader = etcdReader{}
}

func marshal(ext string, m map[string]interface{}) (string, error) {
	var s string
	switch ext {
	case "json":
		b, err := json.Marshal(m)
		if err != nil {
			return "", err
		}
		s = string(b)
	case "toml":
		buf := new(bytes.Buffer)
		if err := toml.NewEncoder(buf).Encode(m); err != nil {
			return "", err
		}
		s = buf.String()
	}
	return s, nil
}

// GetRemoteReader
func ReadRemoteConfig(provider, endpoint, dir, ext string) ([]byte, error) {
	config, err := remoteReader.Get(provider, endpoint, dir, ext)
	if err != nil {
		return nil, err
	}
	s, err := marshal(ext, config)
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
	s, err := marshal(ext, config)
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
