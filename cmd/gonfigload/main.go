package main

import (
	"flag"

	"github.com/katsutomu/gonfigloader"
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
	gonfigloader.Download(endpoint, dir, ext, outpath, outfile)
}
