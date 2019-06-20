package util

import (
	"html"
	"io/ioutil"
	"net/url"
	"os"
	"path"
)

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

func registerRelativeFileCreator() {
	RegisterFileCreator("relative", func(id string) (FileObject, error) {
		u, err := url.Parse(id)
		if err != nil {
			return nil, err
		}
		f := NewRelativeFile()
		f.Location = FileLocation(u.Host)
		f.Path = u.Path
		return f, nil
	})
}

func init() {
	registerRelativeFileCreator()
}
