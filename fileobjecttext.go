package util

import (
	"html"
	"os"
)

type FileObjectText string

func (f FileObjectText) ReadRaw() ([]byte, error) {
	return []byte(string(f)), nil
}

func (f FileObjectText) WriteRaw(data []byte, perm os.FileMode) error {
	return NewFileObjectNotWriteableError(f.ID())
}

func (f FileObjectText) AbsolutePath() string {
	return ""
}
func (f FileObjectText) ID() string {
	return "text://" + html.EscapeString(string(f))
}
func (f FileObjectText) Watcher() FileWatcher {
	return nil
}

func registerFileObejctTextCreator() {
	RegisterFileCreator("text", func(id string) (FileObject, error) {
		return FileObjectText(id[7:]), nil
	})
}

func init() {
	registerFileObejctTextCreator()
}
