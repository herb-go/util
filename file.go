package util

import (
	"html"
	"io/ioutil"
	"net/url"
	"os"
	"path"
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
	Watchable() bool
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
func (f File) Watchable() bool {
	return true
}
func (f File) Watcher() FileWatcher {
	return nil
}

type RelativeFile struct {
	Location FileLocation
	Path     string
}

func (f *RelativeFile) AbsolutePath() string {
	switch f.Location {
	case FileLocationConfig:
		return Config(f.Path)
	case FileLocationConstants:
		return Constants(f.Path)
	case FileLocationSystem:
		return System(f.Path)
	case FileLocationResources:
		return Resources(f.Path)
	case FileLocationAppData:
		return AppData(f.Path)
	}
	return Root(f.Path)
}

func (f *RelativeFile) ReadRaw() ([]byte, error) {
	return ioutil.ReadFile(f.AbsolutePath())
}

func (f *RelativeFile) WriteRaw(data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(f.AbsolutePath(), data, perm)
}
func (f *RelativeFile) ID() string {
	u := url.URL{
		Scheme: "relative",
		Host:   string(f.Location),
		Path:   html.EscapeString(f.Path),
	}
	return u.String()
}

func (f *RelativeFile) Watchable() bool {
	return true
}
func (f *RelativeFile) Watcher() FileWatcher {
	return nil
}

func NewRelativeFile() *RelativeFile {
	return &RelativeFile{}
}

func ConfigFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = FileLocationConfig
	return f
}

func ConstantsFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = FileLocationConstants
	return f
}

func RootFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = FileLocationRoot
	return f
}

func SystemFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = FileLocationSystem
	return f
}

func ResourcesFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = FileLocationResources
	return f
}

func AppDataFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = FileLocationAppData
	return f
}
