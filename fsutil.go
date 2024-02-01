package fsutil

import (
	"os"
	"path/filepath"
)

func MustWrite(path string, content []byte) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return os.WriteFile(path, content, os.ModePerm)
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
