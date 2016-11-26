package main

import (
	"etcd-downloader/locater"
	"etcd-downloader/provider"
	"flag"

	"github.com/labstack/gommon/log"
)

var endpoint string
var dir string
var ext string
var outpath string
var outfile string

func init() {
	flag.StringVar(&endpoint, "endpoint", "http://localhost:2379", "etcd or consol")
	flag.StringVar(&dir, "dir", "/config", "config file path")
	flag.StringVar(&ext, "ext", "json", "file ext")
	flag.StringVar(&outpath, "opath", "./", "output location")
	flag.StringVar(&outfile, "filename", "config.json", "output filename")
}

func main() {
	flag.Parse()
	b, err := provider.ReadRemoteConfig("etcd", endpoint, dir, ext)
	if err != nil {
		log.Error(err)
		return
	}
	if err := locater.LocateFile(outpath, outfile, b); err != nil {
		log.Error(err)
	}
}
