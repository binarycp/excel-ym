package utils

import (
	"syscall"
)

func GetRoot(path string) string {
	fullPath, err := syscall.FullPath(path)
	Err(err)
	return fullPath
}
