package tomlconfig

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/herb-go/herbconfig/configuration"
	"github.com/herb-go/util"
)

func TestTOMLConfig(t *testing.T) {
	var data = `
	#comment
		Data="1234"
	`
	type Data struct {
		Data string
	}
	d := &Data{}
	file := configuration.Text(data)
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

}
