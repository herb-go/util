package util

import (
	"os"
	"path"
	"path/filepath"
)

var DefaultFolderMode = os.FileMode(0700)
var registeredFolders = [][]string{}

func RegisterDataFolder(folder ...string) string {
	registeredFolders = append(registeredFolders, folder)
	return AppData(folder...)
}

func MustLoadRegisteredFolders() {
	for _, v := range registeredFolders {
		folder := AppData(v...)
		_, err := os.Stat(folder)
		if err != nil {
			if os.IsNotExist(err) {

				err = os.MkdirAll(folder, DefaultFolderMode)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func mustPath(path string, err error) string {
	if err != nil {
		panic(err)
	}
	return path
}

var RootPath string
var ResourcesPath string
var AppDataPath string
var ConfigPath string
var SystemPath string
var ConstantsPath string
var UpdatePaths = func() error {
	if RootPath == "" {
		RootPath = filepath.Join(filepath.Dir(mustPath(os.Executable())), "../")
	}
	if ResourcesPath == "" {
		ResourcesPath = filepath.Join(RootPath, "resources")
	}
	if AppDataPath == "" {
		AppDataPath = filepath.Join(RootPath, "appdata")
	}
	if ConfigPath == "" {
		ConfigPath = filepath.Join(RootPath, "config")
	}
	if SystemPath == "" {
		SystemPath = filepath.Join(RootPath, "system")
	}
	if ConstantsPath == "" {
		ConstantsPath = filepath.Join(RootPath, "system", "constants")
	}
	return nil
}

var MustChRoot = func() {
	Must(os.Chdir(RootPath))
}

func MustGetWD() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path
}

func joinPath(p string, filepath ...string) string {
	return path.Join(p, path.Join(filepath...))
}
func Resources(filepaths ...string) string {
	return joinPath(ResourcesPath, filepaths...)
}
func Config(filepaths ...string) string {
	return joinPath(ConfigPath, filepaths...)
}
func AppData(filepaths ...string) string {
	return joinPath(AppDataPath, filepaths...)
}
func System(filepaths ...string) string {
	return joinPath(SystemPath, filepaths...)
}
func Constants(filepaths ...string) string {
	return joinPath(ConstantsPath, filepaths...)
}
func Root(filepaths ...string) string {
	return joinPath(RootPath, filepaths...)
}
