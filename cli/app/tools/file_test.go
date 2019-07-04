package tools

import (
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestFile(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "herb-go-test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.RemoveAll(tmpdir)
	}()
	result, err := FileExists(tmpdir, "notexists")
	if err != nil {
		t.Fatal(err)
	}
	if result {
		t.Fatal(result)
	}
	ioutil.WriteFile(path.Join(tmpdir, "file"), []byte("123"), 0700)
	result, err = FileExists(tmpdir, "file")
	if err != nil {
		t.Fatal(err)
	}
	if !result {
		t.Fatal(result)
	}
	result, err = IsFolder(tmpdir)
	if err != nil {
		t.Fatal(err)
	}
	if !result {
		t.Fatal(result)
	}
	result, err = IsFolder(tmpdir, "file")
	if err != nil {
		t.Fatal(err)
	}
	if result {
		t.Fatal(result)
	}
	result, err = IsFolder(tmpdir, "notexists")
	if err != nil {
		t.Fatal(err)
	}
	if result {
		t.Fatal(result)
	}

}
