package tomlconfig

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/herb-go/util"
)

func TestTOMLConfig(t *testing.T) {
	var data = `
		Data="1234"
	`
	type Data struct {
		Data string
	}
	d := &Data{}
	file := util.FileObjectText(data)
	MustLoad(file, d)
	if d.Data != "1234" {
		t.Fatal(d)
	}
	tmpdir, err := ioutil.TempDir("", "herb-go-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)
	util.RootPath = tmpdir
	defer func() {
		util.RootPath = ""
	}()
	util.UpdatePaths()
	file1 := util.RootFile("test.toml")
	MustSave(file1, d)
	d1 := &Data{}
	MustLoad(file1, d1)
	if d1.Data != "1234" {
		t.Fatal(d)
	}
}
