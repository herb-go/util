package config

import (
	"fmt"
	"sync"

	"github.com/herb-go/herbconfig/source"

	"github.com/herb-go/util"
)

var Debug = false

type Loader struct {
	File     source.Source
	Loader   func(source.Source)
	Position string
	Preload  func()
}

func (l *Loader) Load() {
	util.DebugPrintf("Herb-go util debug: Load config \"%s\"(%s)", l.File.ID(), l.File.AbsolutePath())
	if l.Position != "" {
		util.DebugPrint(l.Position)
	}

	if l.Preload != nil {
		l.Preload()
	}
	l.Loader(l.File)
}

var registeredLoaders = []*Loader{}

var Lock sync.RWMutex

func CleanLoaders() {
	registeredLoaders = []*Loader{}
}

func RegisterLoader(file source.Source, loader func(file source.Source)) {
	var position string
	lines := util.GetStackLines(8, 9)
	if len(lines) == 1 {
		position = fmt.Sprintf("%s\r\n", lines[0])
	}
	l := Loader{File: file, Loader: loader, Position: position}
	registeredLoaders = append(registeredLoaders, &l)
}

func RegisterLoaderAndWatch(file source.Source, loader func(source.Source)) *Loader {
	var position string
	lines := util.GetStackLines(8, 9)
	if len(lines) == 1 {
		position = fmt.Sprintf("%s\r\n", lines[0])
	}
	l := Loader{File: file, Loader: loader, Position: position}
	l.Preload = func() {
		WatcherManager.Watch(file, func() {
			loader(file)
		})
	}
	registeredLoaders = append(registeredLoaders, &l)

	return &l
}
func LoadAll(files ...source.Source) {
	if util.ConfigPath == "" {
		panic(ErrConfigPathNotInited)
	}
	defer Lock.RUnlock()
	Lock.RLock()
	util.Must(WatcherManager.Start())
NextLoader:
	for _, v := range registeredLoaders {
		if len(files) != 0 {
			for _, configfile := range files {
				if source.IsSame(v.File, configfile) {
					v.Load()
					continue NextLoader
				}
			}
		} else {
			v.Load()
		}
	}
}
