package config

import (
	"io/ioutil"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/herb-go/util"
)

func TestLoader(t *testing.T) {
	defer func() {
		CleanLoaders()
	}()
	CleanLoaders()
	var jdata = `//comment
	{"Data":"12345"}`
	RegisterLoader(util.FileObjectText(jdata), func(file util.FileObject) {

	})
	LoadAll()
}

func TestNamedLoader(t *testing.T) {
	defer func() {
		CleanLoaders()
	}()
	CleanLoaders()
	var jdata = `//comment
	{"Data":"12345"}`
	RegisterLoader(util.FileObjectText(jdata), func(file util.FileObject) {

	})
	LoadAll(util.FileObjectText(jdata))
}

func TestWatcher(t *testing.T) {
	defer func() {
		CleanLoaders()
	}()
	CleanLoaders()
	var data string
	tmpdir, err := ioutil.TempDir("", "herb-go-test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.RemoveAll(tmpdir)
	}()
	util.RootPath = tmpdir
	util.UpdatePaths()
	file := util.RootFile("test.toml")
	err = util.WriteFile(file, []byte("test"), 0700)
	if err != nil {
		t.Fatal(err)
	}
	wg := sync.WaitGroup{}
	RegisterLoaderAndWatch(file, func(file util.FileObject) {
		d, err := util.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		data = string(d)
		wg.Done()
	})
	wg.Add(1)
	LoadAll()
	wg.Add(1)
	err = util.WriteFile(file, []byte("test1"), 0700)
	if err != nil {
		t.Fatal(err)
	}
	wg.Wait()
	if data != "test1" {
		t.Fatal(data)
	}

	Watcher.Close()

	time.Sleep(time.Millisecond)
	defer func() {
		err := Watcher.Start()
		if err != nil {
			t.Fatal(err)
		}
	}()
	err = util.WriteFile(file, []byte("test2"), 0700)
	time.Sleep(time.Millisecond)

}
