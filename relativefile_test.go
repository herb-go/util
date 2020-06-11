package util

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/herb-go/herbconfig/source"
)

func TestOverwriteableConfig(t *testing.T) {
	var data []byte
	defer func() {
		ConfigPath = ""
		DefaultConfigPath = ""
	}()
	tmpdir, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	ConfigPath = tmpdir
	if tmpdir == "" {
		t.Fatal(tmpdir)
	}
	defer func() {
		os.RemoveAll(tmpdir)
	}()
	DefaultConfigPath = filepath.Join(tmpdir, "defaultconfig")
	err = os.Mkdir(DefaultConfigPath, DefaultFolderMode)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(path.Join(tmpdir, "configonly"), []byte("configonly"), DefaultFileMode)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(path.Join(tmpdir, "config"), []byte("config"), DefaultFileMode)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(path.Join(DefaultConfigPath, "config"), []byte("config"), DefaultFileMode)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(path.Join(DefaultConfigPath, "defaultconfigonly"), []byte("defaultconfigonly"), DefaultFileMode)
	if err != nil {
		panic(err)
	}
	configonly, err := source.New("relative://config/configonly")
	if err != nil {
		panic(err)
	}
	data, err = configonly.ReadRaw()
	if err != nil {
		panic(err)
	}
	if string(data) != "configonly" {
		t.Fatal(string(data))
	}
	config, err := source.New("relative://config/config")
	if err != nil {
		panic(err)
	}
	data, err = config.ReadRaw()
	if err != nil {
		panic(err)
	}
	if string(data) != "config" {
		t.Fatal(string(data))
	}
	defaultconfigonly, err := source.New("relative://config/defaultconfigonly")
	if err != nil {
		panic(err)
	}
	data, err = defaultconfigonly.ReadRaw()
	if err != nil {
		panic(err)
	}
	if string(data) != "defaultconfigonly" {
		t.Fatal(string(data))
	}
	notexist, err := source.New("relative://config/notexist")
	if err != nil {
		panic(err)
	}
	_, err = notexist.ReadRaw()
	if !os.IsNotExist(err) {
		panic(err)
	}

}
