package util

import (
	"os"
	"path/filepath"
)

var DefaultFolderMode = os.FileMode(0770)
var registeredFolders = [][]string{}

func RegisterDataFolder(folder ...string) string {
	registeredFolders = append(registeredFolders, folder)
	return AppData(folder...)
}
func MakeLoggerFolderIfNotExist() {
	_, err := os.Stat(LogsPath)
	if err != nil {
		if os.IsNotExist(err) {

			err = os.MkdirAll(LogsPath, DefaultFolderMode)
			if err != nil {
				panic(err)
			}
		}
	}
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
func IsPathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

var AutoDetectCWD = true
var RootPath string
var ResourcesPath string
var AppDataPath string
var LogsPath string
var ConfigPath string
var SystemPath string
var ConstantsPath string
var DefaultConfigPath string
var UpdatePaths = func() error {
	if RootPath == "" {
		var execute string
		execute = filepath.Dir(mustPath(os.Executable()))
		if AutoDetectCWD && ConfigPath == "" && !IsPathExists(filepath.Join(execute, "../", "config")) {
			execute = MustGetWD()
		}
		RootPath = filepath.Join(execute, "../")
	}
	if ResourcesPath == "" {
		ResourcesPath = filepath.Join(RootPath, "resources")
	}
	if AppDataPath == "" {
		AppDataPath = filepath.Join(RootPath, "appdata")
	}
	if LogsPath == "" {
		LogsPath = filepath.Join(RootPath, "logs")
	}
	if ConfigPath == "" {
		ConfigPath = filepath.Join(RootPath, "config")
	}
	if SystemPath == "" {
		SystemPath = filepath.Join(RootPath, "system")
	}
	if ConstantsPath == "" {
		ConstantsPath = filepath.Join(SystemPath, "constants")
	}
	if DefaultConfigPath == "" {
		DefaultConfigPath = filepath.Join(SystemPath, "defaultconfig")
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

func joinPath(p string, filepaths ...string) string {
	return filepath.Join(p, filepath.Join(filepaths...))
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
func Logs(filepaths ...string) string {
	return joinPath(LogsPath, filepaths...)
}
func System(filepaths ...string) string {
	return joinPath(SystemPath, filepaths...)
}
func Constants(filepaths ...string) string {
	return joinPath(ConstantsPath, filepaths...)
}
func DefaultConfig(filepaths ...string) string {
	return joinPath(DefaultConfigPath, filepaths...)
}
func Root(filepaths ...string) string {
	return joinPath(RootPath, filepaths...)
}
