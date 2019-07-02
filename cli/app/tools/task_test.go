package tools

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/herb-go/util/cli/app"
)

func TestTask(t *testing.T) {
	var jobsoutput = ""
	tmpdir, err := ioutil.TempDir("", "herb-go-test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.RemoveAll(tmpdir)
	}()
	renderdata := map[string]interface{}{
		"data": "data",
	}
	task := NewTask(path.Join("./", "testdata"), tmpdir)
	err = task.Copy("/demo.txt", "/demo.txt")
	if err != nil {
		t.Fatal(err)
	}
	err = task.Render("/demo1.tmpl", "/output/demo1.txt", renderdata)
	if err != nil {
		t.Fatal(err)
	}
	files := task.ListFiles()
	if len(files) != 2 || files[0] != "/demo.txt" || files[1] != "/output/demo1.txt" {
		t.Fatal(files)
	}
	task.AddJob(func() error {
		jobsoutput = jobsoutput + "job1"
		return nil
	})
	task.AddJob(func() error {
		jobsoutput = jobsoutput + "job2"
		return nil
	})
	err = task.Exec()
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := ioutil.ReadFile(path.Join(tmpdir, "/demo.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != "123" {
		t.Fatal(string(bytes))
	}
	bytes, err = ioutil.ReadFile(path.Join(tmpdir, "/output/demo1.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != "data" {
		t.Fatal(string(bytes))
	}
	if jobsoutput != "job1job2" {
		t.Fatal(jobsoutput)
	}
}

func TestTaskFiles(t *testing.T) {
	App := app.NewApplication(app.NewApplicationConfig())
	inputr, inputw, _ := os.Pipe()
	outputw := bytes.NewBuffer([]byte{})
	App.Stdin = inputr
	App.Stdout = outputw

	tmpdir, err := ioutil.TempDir("", "herb-go-test")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		os.RemoveAll(tmpdir)
	}()
	renderdata := map[string]interface{}{
		"data":  "data",
		"data2": "data2",
	}
	task := NewTask(path.Join("./", "testdata"), tmpdir)
	result, err := task.ConfirmIf(App, false)
	if err != nil {
		t.Fatal(err)
	}
	if result != true {
		t.Fatal(result)
	}

	err = task.CopyFiles(map[string]string{"/demo.txt": "/demo.txt", "/demo2.txt": "/demo2.txt"})
	if err != nil {
		t.Fatal(err)
	}
	err = task.RenderFiles(map[string]string{"/output/demo1.txt": "/demo1.tmpl", "/output/demo3.txt": "/demo3.tmpl"}, renderdata)
	if err != nil {
		t.Fatal(err)
	}
	files := task.ListFiles()
	if len(files) != 4 || files[0] != "/demo.txt" || files[1] != "/demo2.txt" || files[2] != "/output/demo1.txt" || files[3] != "/output/demo3.txt" {
		t.Fatal(files)
	}

	result, err = task.ConfirmIf(App, false)
	if err != nil {
		t.Fatal(err)
	}
	if result != true {
		t.Fatal(result)
	}
	go func() {
		_, err = inputw.Write([]byte("n\n"))
		if err != nil {
			t.Fatal(err)
		}
	}()
	result, err = task.ConfirmIf(App, true)
	if err != nil {
		t.Fatal(err)
	}
	if result != false {
		t.Fatal(result)
	}
	go func() {
		_, err = inputw.Write([]byte("y\n"))
		if err != nil {
			t.Fatal(err)
		}
	}()

	bs, err := ioutil.ReadAll(outputw)
	if err != nil {
		t.Fatal(err)
	}
	output := string(bs)
	if !strings.Contains(output, "/demo.txt") || !strings.Contains(output, "/demo2.txt") || !strings.Contains(output, "/output/demo1.txt") || !strings.Contains(output, "/output/demo3.txt") {
		t.Fatal(output)
	}
	result, err = task.ConfirmIf(App, true)
	if err != nil {
		t.Fatal(err)
	}
	if result != true {
		t.Fatal(result)
	}
	err = task.Exec()
	if err != nil {
		t.Fatal(err)
	}
	bytes, err := ioutil.ReadFile(path.Join(tmpdir, "/demo.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != "123" {
		t.Fatal(string(bytes))
	}
	bytes, err = ioutil.ReadFile(path.Join(tmpdir, "/demo2.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != "234" {
		t.Fatal(string(bytes))
	}

	bytes, err = ioutil.ReadFile(path.Join(tmpdir, "/output/demo1.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != "data" {
		t.Fatal(string(bytes))
	}
	bytes, err = ioutil.ReadFile(path.Join(tmpdir, "/output/demo3.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != "data2" {
		t.Fatal(string(bytes))
	}
}
