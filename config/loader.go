package config

import (
	"fmt"
	"sync"

	"github.com/herb-go/util"
)

var Debug = false

type Loader struct {
	File     util.FileObject
	Loader   func(util.FileObject)
	Position string
	Preload  func()
}

func (l *Loader) Load() {
	if Debug || util.ForceDebug {
		fmt.Printf("Herb-go util debug: Load config \"%s\"", l.File.URI())
		if l.Position != "" {
			fmt.Print(l.Position)
		}

	}
	if l.Preload != nil {
		l.Preload()
	}
	l.Loader(l.File)

}

var registeredLoader = []Loader{}

var Lock sync.RWMutex

func RegisterLoader(file util.FileObject, loader func(file util.FileObject)) {
	var position string
	lines := util.GetStackLines(8, 9)
	if len(lines) == 1 {
		position = fmt.Sprintf("%s\r\n", lines[0])
	}
	l := Loader{File: file, Loader: loader, Position: position}
	registeredLoader = append(registeredLoader, l)
}

func RegisterLoaderAndWatch(file util.FileObject, loader func(util.FileObject)) *Loader {
	var position string
	lines := util.GetStackLines(8, 9)
	if len(lines) == 1 {
		position = fmt.Sprintf("%s\r\n", lines[0])
	}
	l := Loader{File: file, Loader: loader, Position: position}
	registeredLoader = append(registeredLoader, l)
	l.Preload = Watcher.Watch(file, func() {
		l.Load()
	})
	return &l
}
func LoadAll(files ...util.FileObject) {
	defer Lock.RUnlock()
	Lock.RLock()
	Watcher = NewWatcherManager()
	err := Watcher.Start()
	if err != nil {
		panic(err)
	}
NextLoader:
	for _, v := range registeredLoader {
		if len(files) != 0 {
			for _, configfile := range files {
				if util.IsSameFile(v.File, configfile) {
					v.Load()
					continue NextLoader
				}
			}
		} else {
			v.Load()
		}
	}
}
