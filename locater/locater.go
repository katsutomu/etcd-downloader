package locater

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// LocateFile
func LocateFile(path string, filename string, b []byte) error {
	if err := os.MkdirAll(path, 0700); err != nil {
		fmt.Println(err)
	}
	return ioutil.WriteFile(strings.Join([]string{path, filename}, "/"), b, os.ModePerm)
}
