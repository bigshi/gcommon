package gutil

import (
	"fmt"
	"os"
	"time"
)

func ExistPath(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func ListenTmp(port int) {
	dir := fmt.Sprintf("/tmp/golang.%d.%d", port, time.Now().UnixNano())
	if !ExistPath(dir) {
		os.MkdirAll(dir, os.FileMode(0777))
	}
}
