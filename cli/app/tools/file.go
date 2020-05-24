package tools

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/herb-go/util"
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

func CopyIfNotExist(src string, target ...string) error {
	result, err := FileExists(target...)
	if err != nil {
		return err
	}
	if result {
		return nil
	}
	bs, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(target...), bs, util.DefaultFileMode)
}
