package gonfigloader

import (
	"gonfigloader/locater"
	"gonfigloader/provider"

	"github.com/labstack/gommon/log"
)

func Download(endpoint, dir, ext, outpath, outfile string) {
	b, err := provider.ReadRemoteConfig("etcd", endpoint, dir, ext)
	if err != nil {
		log.Error(err)
		return
	}
	if err := locater.LocateFile(outpath, outfile, b); err != nil {
		log.Error(err)
	}
}
