package util

import "testing"

func TestFileCreator(t *testing.T) {
	tp := getFileTypeFromID("http://www.example.com")
	if tp != "http" {
		t.Fatal(tp)
	}
	tp = getFileTypeFromID("www.example.com")
	if tp != "" {
		t.Fatal(tp)
	}
	_, err := NewFileObect("notexistscheme://123")
	if err == nil || GetErrorType(err) != ErrTypeFileObjectSchemeNotavaliable {
		t.Fatal(err)
	}
	_, err = NewFileObect("")
	if err == nil || GetErrorType(err) != ErrTypeFileObjectSchemeNotavaliable {
		t.Fatal(err)
	}
	localfile := File("/tmp/test.go")
	file, err := NewFileObect(localfile.ID())
	if !IsSameFile(localfile, file) {
		t.Fatal(file.ID())
	}
	relativefile := ConfigFile("test.toml")
	file, err = NewFileObect(relativefile.ID())
	if !IsSameFile(relativefile, file) {
		t.Fatal(file.ID())
	}
	textfile := FileObjectText("test")
	file, err = NewFileObect(textfile.ID())
	if !IsSameFile(textfile, file) {
		t.Fatal(file.ID())
	}
}
