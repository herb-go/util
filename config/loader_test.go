package config

import (
	"html"
	"io/ioutil"
	"net/url"
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
	err = util.WriteFile(file, []byte("test"), 0700)
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		RegisterLoaderAndWatch(file, func(file util.FileObject) {
			d, err := util.ReadFile(file)
			if err != nil {
				t.Fatal(err)
			}
			data = string(d)
		})
		RegisterLoaderAndWatch(file, func(file util.FileObject) {
			d, err := util.ReadFile(file)
			if err != nil {
				t.Fatal(err)
			}
			data1 = string(d)
		})
		RegisterLoaderAndWatch(util.FileObjectText("test"), func(file util.FileObject) {
		})
		LoadAll()
		err = util.WriteFile(file, []byte("test1"), 0700)
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
	Watcher.Close()

	defer func() {
		err := Watcher.Start()
		if err != nil {
			t.Fatal(err)
		}
	}()

	err = util.WriteFile(file, []byte("test2"), 0700)
	time.Sleep(time.Second)
	if data != "test1" {
		t.Fatal(data)
	}
	if data1 != "test1" {
		t.Fatal(data)
	}
}

type testfile struct {
	util.File
	Data []byte
	E    chan int
	C    chan int
}

func (f *testfile) ReadRaw() ([]byte, error) {
	return f.Data, nil
}

func (f *testfile) WriteRaw(data []byte, perm os.FileMode) error {
	f.Data = data
	return nil
}

func (f *testfile) AbsolutePath() string {
	return ""
}
func (f *testfile) ID() string {
	u := url.URL{
		Scheme: "testfile",
		Host:   "local",
		Path:   html.EscapeString(string(f.File)),
	}
	return u.String()
}
func (f *testfile) Watcher() util.FileWatcher {
	return func(callback func()) (unwatcher func()) {
		go func() {
			for {
				select {
				case <-f.E:
					go callback()
				case <-f.C:
					return
				}
			}
		}()
		return func() {
			close(f.C)
		}
	}
}
func newTestFile(name string) *testfile {
	return &testfile{
		File: util.File(name),
		Data: []byte{},
		E:    make(chan int, 2),
		C:    make(chan int),
	}
}
func TestUnwatch(t *testing.T) {
	var data string
	MustResetWatcher()
	defer func() {
		MustResetWatcher()
	}()
	file := newTestFile("test")
	err := util.WriteFile(file, []byte("test"), util.DefaultFileMode)
	if err != nil {
		t.Fatal(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)

	RegisterLoaderAndWatch(file, func(file util.FileObject) {
		d, err := util.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		data = string(d)
		wg.Done()
	})
	LoadAll()
	wg.Wait()
	if data != "test" {
		t.Fatal(data)
	}
	wg.Add(1)
	err = util.WriteFile(file, []byte("test1"), util.DefaultFileMode)
	if err != nil {
		t.Fatal(err)
	}
	file.E <- 0
	wg.Wait()
	if data != "test1" {
		t.Fatal(data)
	}
	MustResetWatcher()
	err = util.WriteFile(file, []byte("test2"), util.DefaultFileMode)
	if err != nil {
		t.Fatal(err)
	}
	wg.Add(1)

	file.E <- 0
	if data != "test1" {
		t.Fatal(data)
	}
}
