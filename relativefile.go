package util

import (
	"html"
	"io/ioutil"
	"net/url"
	"os"
	"path"

	"github.com/herb-go/herbconfig/source"
)

type RelativeFileLocation string

const RelativeFileLocationRoot = RelativeFileLocation("")
const RelativeFileLocationLogs = RelativeFileLocation("logs")
const RelativeFileLocationAppData = RelativeFileLocation("appdata")
const RelativeFileLocationConfig = RelativeFileLocation("config")
const RelativeFileLocationConstants = RelativeFileLocation("constants")
const RelativeFileLocationSystem = RelativeFileLocation("system")
const RelativeFileLocationResources = RelativeFileLocation("resources")

type RelativeFile struct {
	Location RelativeFileLocation
	Path     string
}

func (f *RelativeFile) AbsolutePath() string {
	switch f.Location {
	case RelativeFileLocationConfig:
		p := Config(f.Path)
		_, err := os.Stat(p)
		if err == nil {
			return p
		}
		return DefaultConfig(f.Path)
	case RelativeFileLocationConstants:
		return Constants(f.Path)
	case RelativeFileLocationSystem:
		return System(f.Path)
	case RelativeFileLocationResources:
		return Resources(f.Path)
	case RelativeFileLocationAppData:
		return AppData(f.Path)
	case RelativeFileLocationLogs:
		return Logs(f.Path)
	}
	return Root(f.Path)
}

func (f *RelativeFile) ReadRaw() ([]byte, error) {
	return ioutil.ReadFile(f.AbsolutePath())
}

func (f *RelativeFile) ID() string {
	u := url.URL{
		Scheme: "relative",
		Host:   string(f.Location),
		Path:   html.EscapeString(f.Path),
	}
	return u.String()
}

func NewRelativeFile() *RelativeFile {
	return &RelativeFile{}
}

func ConfigFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = RelativeFileLocationConfig
	return f
}

func ConstantsFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = RelativeFileLocationConstants
	return f
}

func RootFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = RelativeFileLocationRoot
	return f
}

func SystemFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = RelativeFileLocationSystem
	return f
}

func ResourcesFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = RelativeFileLocationResources
	return f
}

func AppDataFile(filepath ...string) *RelativeFile {
	f := NewRelativeFile()
	f.Path = path.Join(filepath...)
	f.Location = RelativeFileLocationAppData
	return f
}

func registerRelativeFileCreator() {
	source.RegisterCreator("relative", func(id string) (source.Source, error) {
		u, err := url.Parse(id)
		if err != nil {
			return nil, err
		}
		f := NewRelativeFile()
		f.Location = RelativeFileLocation(u.Host)
		f.Path = u.Path
		return f, nil
	})
}

func init() {
	registerRelativeFileCreator()
}
