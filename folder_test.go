package util

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestFolder(t *testing.T) {
	defer func() {
		RootPath = ""
		ResourcesPath = ""
		AppDataPath = ""
		ConfigPath = ""
		SystemPath = ""
		ConstantsPath = ""
	}()
	executable, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	wd := filepath.Join(filepath.Dir(executable), "../")
	UpdatePaths()
	if RootPath != wd {
		t.Fatal(RootPath)
	}
	if ResourcesPath != path.Join(RootPath, "resources") {
		t.Fatal(ResourcesPath)
	}
	if AppDataPath != path.Join(RootPath, "appdata") {
		t.Fatal(AppDataPath)
	}
	if ConfigPath != path.Join(RootPath, "config") {
		t.Fatal(ConfigPath)
	}
	if SystemPath != path.Join(RootPath, "system") {
		t.Fatal(SystemPath)
	}
	if ConstantsPath != path.Join(RootPath, "system", "constants") {
		t.Fatal(ConstantsPath)
	}
	MustChRoot()
	cwd := MustGetWD()
	if cwd != wd {
		t.Fatal(cwd)
	}
	RootPath = "/fake/RootPath"
	ResourcesPath = "/fake/ResourcesPath"
	AppDataPath = "/fake/AppDataPath"
	ConfigPath = "/fake/ConfigPath"
	SystemPath = "/fake/SystemPath"
	ConstantsPath = "/fake/ConstantsPath"

	UpdatePaths()
	if RootPath != "/fake/RootPath" {
		t.Fatal(RootPath)
	}
	if ResourcesPath != "/fake/ResourcesPath" {
		t.Fatal(ResourcesPath)
	}
	if AppDataPath != "/fake/AppDataPath" {
		t.Fatal(AppDataPath)
	}
	if ConfigPath != "/fake/ConfigPath" {
		t.Fatal(ConfigPath)
	}
	if SystemPath != "/fake/SystemPath" {
		t.Fatal(SystemPath)
	}
	if ConstantsPath != "/fake/ConstantsPath" {
		t.Fatal(ConstantsPath)
	}
	p := Root("folder", "file")
	if p != path.Join(RootPath, "folder", "file") {
		t.Fatal(p)
	}
	p = Resources("folder", "file")
	if p != path.Join(ResourcesPath, "folder", "file") {
		t.Fatal(p)
	}
	p = AppData("folder", "file")
	if p != path.Join(AppDataPath, "folder", "file") {
		t.Fatal(p)
	}
	p = Config("folder", "file")
	if p != path.Join(ConfigPath, "folder", "file") {
		t.Fatal(p)
	}
	p = System("folder", "file")
	if p != path.Join(SystemPath, "folder", "file") {
		t.Fatal(p)
	}
	p = Constants("folder", "file")
	if p != path.Join(ConstantsPath, "folder", "file") {
		t.Fatal(p)
	}
}

func TestRegisteredDataFolder(t *testing.T) {
	defer func() {
		AppDataPath = ""
	}()
	folder, err := ioutil.TempDir("", "herb-go-test")
	AppDataPath = folder
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.RemoveAll(folder)
	}()
	folder1 := RegisterDataFolder("test", "folder1")
	folder2 := RegisterDataFolder("test2", "folder1")
	_, err = os.Stat(folder1)
	if !os.IsNotExist(err) {
		t.Fatal(err)
	}
	_, err = os.Stat(folder2)
	if !os.IsNotExist(err) {
		t.Fatal(err)
	}
	MustLoadRegisteredFolders()
	_, err = os.Stat(folder1)
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.Stat(folder2)
	if err != nil {
		t.Fatal(err)
	}

}
