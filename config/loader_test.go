package config

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/herb-go/herbconfig/configuration"

	"github.com/herb-go/util"
)

func TestLoader(t *testing.T) {
	util.UpdatePaths()
	defer func() {
		CleanLoaders()
	}()
	CleanLoaders()
	var jdata = `//comment
	{"Data":"12345"}`
	RegisterLoader(configuration.Text(jdata), func(file configuration.Configuration) {

	})
	LoadAll()
}

func TestNamedLoader(t *testing.T) {
	util.UpdatePaths()
	defer func() {
		CleanLoaders()
	}()
	CleanLoaders()
	var jdata = `//comment
	{"Data":"12345"}`
	RegisterLoader(configuration.Text(jdata), func(file configuration.Configuration) {

	})
	LoadAll(configuration.Text(jdata))
}

func TestWatcher(t *testing.T) {
	defer func() {
		CleanLoaders()
	}()
	CleanLoaders()
	var data string
	var data1 string
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
	err = ioutil.WriteFile(file.AbsolutePath(), []byte("test"), 0700)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second)

	go func() {
		RegisterLoaderAndWatch(file, func(file configuration.Configuration) {
			d, err := configuration.Read(file)
			if err != nil {
				t.Fatal(err)
			}
			data = string(d)
		})
		RegisterLoaderAndWatch(file, func(file configuration.Configuration) {
			d, err := configuration.Read(file)
			if err != nil {
				t.Fatal(err)
			}
			data1 = string(d)
		})
		RegisterLoaderAndWatch(configuration.Text("test"), func(file configuration.Configuration) {
		})
		LoadAll()
		err = ioutil.WriteFile(file.AbsolutePath(), []byte("test1"), 0700)
		if err != nil {
			t.Fatal(err)
		}
	}()
	time.Sleep(time.Second)
	if data != "test1" {
		t.Fatal(data)
	}
	if data1 != "test1" {
		t.Fatal(data)
	}
	util.Must(WatcherManager.Stop())

	defer func() {
		err := WatcherManager.Start()
		if err != nil {
			t.Fatal(err)
		}
	}()

	err = ioutil.WriteFile(file.AbsolutePath(), []byte("test2"), 0700)
	time.Sleep(time.Second)
	if data != "test1" {
		t.Fatal(data)
	}
	if data1 != "test1" {
		t.Fatal(data)
	}
}

// type testfile struct {
// 	configuration.File
// 	Data []byte
// 	E    chan int
// 	C    chan int
// }

// func (f *testfile) ReadRaw() ([]byte, error) {
// 	return f.Data, nil
// }

// func (f *testfile) AbsolutePath() string {
// 	return ""
// }
// func (f *testfile) ID() string {
// 	u := url.URL{
// 		Scheme: "testfile",
// 		Host:   "local",
// 		Path:   html.EscapeString(string(f.File)),
// 	}
// 	return u.String()
// }

// func newTestFile(name string) *testfile {
// 	return &testfile{
// 		File: configuration.File(name),
// 		Data: []byte{},
// 		E:    make(chan int, 2),
// 		C:    make(chan int),
// 	}
// }
