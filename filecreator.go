package util

import (
	"strings"
	"sync"
)

var registeredFileCreators = map[string]func(id string) (FileObject, error){}

var fileCreatorRegisterLock = sync.Mutex{}

func getFileTypeFromID(id string) string {
	index := strings.Index(id, "://")
	if index < 0 {
		return ""
	}
	return id[0:index]
}
func RegisterFileCreator(name string, creator func(id string) (FileObject, error)) {
	fileCreatorRegisterLock.Lock()
	defer fileCreatorRegisterLock.Unlock()
	registeredFileCreators[name] = creator
}

func NewFileObect(id string) (FileObject, error) {
	tp := getFileTypeFromID(id)
	creator := registeredFileCreators[tp]
	if tp == "" || creator == nil {
		return nil, NewFileObjectSchemeError(id)
	}
	return creator(id)
}
