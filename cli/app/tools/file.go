package tools

import (
	"os"
	"path/filepath"
)

func FileExists(pathlist ...string) (bool, error) {
	path := filepath.Join(pathlist...)
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsFolder(pathlist ...string) (bool, error) {
	path := filepath.Join(pathlist...)
	mode, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return mode.IsDir(), nil
}
