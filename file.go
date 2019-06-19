package util

import (
	"io/ioutil"
	"net/url"
	"os"
	"path"
)

type FileLocation string

const FileLocationRoot = FileLocation("")
const FileLocationConfig = FileLocation("config")
const FileLocationConstants = FileLocation("constants")
const FileLocationSystem = FileLocation("system")
const FileLocationResoures = FileLocation("resoures")

func IsSameFile(src FileObject, dst FileObject) bool {
	return src.ID() == dst.ID()
}

type FileWatcher func(callback func()) (unwatcher func())

type FileObject interface {
	ReadRaw() ([]byte, error)
	WriteRaw([]byte, os.FileMode) error
	URI() string
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
	return ioutil.ReadFile(f.URI())
}

func (f File) WriteRaw(data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(f.URI(), data, perm)
}

func (f File) URI() string {
	return string(f)
}
func (f File) ID() string {
	return path.Join("file://local", url.PathEscape((string(f))))
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

func (f *RelativeFile) URI() string {
	switch f.Location {
	case FileLocationConfig:
		return Config(f.Path)
	case FileLocationConstants:
		return Constants(f.Path)
	case FileLocationSystem:
		return System(f.Path)
	case FileLocationResoures:
		return Resource(f.Path)

	}
	return Root(f.Path)
}

func (f *RelativeFile) ReadRaw() ([]byte, error) {
	return ioutil.ReadFile(f.URI())
}

func (f *RelativeFile) WriteRaw(data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(f.URI(), data, perm)
}
func (f *RelativeFile) ID() string {
	return path.Join("relative://", string(f.Location), url.PathEscape(f.Path))
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

func ResouresFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = FileLocationResoures
	return f
}
