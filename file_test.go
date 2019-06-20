package util

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFile(t *testing.T) {
	file, err := ioutil.TempFile("", "herb-go-test")
	if err != nil {
		t.Fatal(err)
	}
	_, err = file.WriteString("testcontent")
	if err != nil {
		t.Fatal(err)
	}
	name := file.Name()
	file.Close()
	defer os.Remove(name)
	file1 := File(name)
	data, err := ReadFile(file1)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "testcontent" {
		t.Fatal(string(data))
	}

	if !file1.Watchable() {
		t.Fatal(file1.Watchable())
	}
	if file1.Watcher() != nil {
		t.Fatal(file1.Watcher())
	}
	err = WriteFile(file1, []byte("testcontentupdated"), 0700)
	if err != nil {
		t.Fatal(err)
	}
	data, err = ReadFile(file1)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "testcontentupdated" {
		t.Fatal(string(data))
	}
	file2 := File(name + ".notexists")

	if IsSameFile(file1, file2) {
		t.Fatal(file2.ID())
	}
}

func TestRelativeFile(t *testing.T) {
	defer func() {
		RootPath = ""
		ResourcesPath = ""
		AppDataPath = ""
		ConfigPath = ""
		SystemPath = ""
		ConstantsPath = ""
	}()
	root, err := ioutil.TempDir("", "herb-go-test")
	if err != nil {
		t.Fatal(err)
	}
	RootPath = root
	UpdatePaths()
	file := RootFile("folder", "file")
	if filepath.Dir(file.AbsolutePath()) != Root("folder") {
		t.Fatal(file)
	}
	file = ResourcesFile("folder", "file")
	if filepath.Dir(file.AbsolutePath()) != Resources("folder") {
		t.Fatal(file)
	}
	file = AppDataFile("folder", "file")
	if filepath.Dir(file.AbsolutePath()) != AppData("folder") {
		t.Fatal(file)
	}
	file = ConfigFile("folder", "file")
	if filepath.Dir(file.AbsolutePath()) != Config("folder") {
		t.Fatal(file)
	}
	file = SystemFile("folder", "file")
	if filepath.Dir(file.AbsolutePath()) != System("folder") {
		t.Fatal(file)
	}
	file = ConstantsFile("folder", "file")
	if filepath.Dir(file.AbsolutePath()) != Constants("folder") {
		t.Fatal(file)
	}
	file = RootFile("file")
	err = WriteFile(file, []byte("testcontent"), 0700)
	if err != nil {
		t.Fatal(err)
	}
	if !file.Watchable() {
		t.Fatal(file.Watchable())
	}
	if file.Watcher() != nil {
		t.Fatal(file.Watcher())
	}

	data, err := ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "testcontent" {
		t.Fatal(string(data))
	}
	file2 := ConfigFile("file")
	if IsSameFile(file, file2) {
		t.Fatal(file2.ID())
	}
}
