package util

import (
	"html"
	"io/ioutil"
	"net/url"
	"os"
)

type FileLocation string

const FileLocationRoot = FileLocation("")
const FileLocationAppData = FileLocation("appdata")
const FileLocationConfig = FileLocation("config")
const FileLocationConstants = FileLocation("constants")
const FileLocationSystem = FileLocation("system")
const FileLocationResources = FileLocation("resources")

func IsSameFile(src FileObject, dst FileObject) bool {
	return src.ID() == dst.ID()
}

type FileWatcher func(callback func()) (unwatcher func())

type FileObject interface {
	ReadRaw() ([]byte, error)
	WriteRaw([]byte, os.FileMode) error
	AbsolutePath() string
	ID() string
	Watcher() FileWatcher
}

func ReadFile(file FileObject) ([]byte, error) {
	return file.ReadRaw()
}

func WriteFile(file FileObject, data []byte, mode os.FileMode) error {
	return file.WriteRaw(data, mode)
}

type File string

func (f File) ReadRaw() ([]byte, error) {
	return ioutil.ReadFile(f.AbsolutePath())
}

func (f File) WriteRaw(data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(f.AbsolutePath(), data, perm)
}

func (f File) AbsolutePath() string {
	return string(f)
}
func (f File) ID() string {
	u := url.URL{
		Scheme: "file",
		Host:   "local",
		Path:   html.EscapeString(string(f)),
	}
	return u.String()
}
func (f File) Watcher() FileWatcher {
	return nil
}

func registerLocalFileCreator() {
	RegisterFileCreator("file", func(id string) (FileObject, error) {
		u, err := url.Parse(id)
		if err != nil {
			return nil, err
		}
		return File(u.Path), nil
	})
}

func init() {
	registerLocalFileCreator()
}
