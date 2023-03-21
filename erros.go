package util

import (
	"errors"
	"fmt"
)

type ErrorType string

const ErrTypeFileObjectSchemeNotavaliable = ErrorType("schemenotavaliable")
const ErrTypeFileObjectNotWriteable = ErrorType("notwriteable")

type FileObjectError struct {
	Type   ErrorType
	Msg    string
	FileID string
}

func (f *FileObjectError) Error() string {
	return fmt.Sprintf(f.Msg, f.FileID)
}
func NewFileObjectSchemeError(id string) error {
	return &FileObjectError{
		Type:   ErrTypeFileObjectSchemeNotavaliable,
		Msg:    "file scheme of file object \"%s\" is not avaliable",
		FileID: id,
	}
}

func NewFileObjectNotWriteableError(id string) error {
	return &FileObjectError{
		Type:   ErrTypeFileObjectNotWriteable,
		Msg:    "file object \"%s\" is not writeable",
		FileID: id,
	}
}

func GetErrorType(err error) ErrorType {
	if err == nil {
		return ""
	}
	e, ok := err.(*FileObjectError)
	if ok == false {
		return ""
	}
	return e.Type
}

func Catch(f func()) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			s, ok := r.(string)
			if ok {
				err = errors.New(s)
				return
			}
			err = r.(error)
		}
	}()
	f()
	return nil
}
