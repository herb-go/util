package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/herb-go/util"
)

// Op describes a set of file operations.
type Op uint32

var (
	Create = Op(fsnotify.Create)
	Write  = Op(fsnotify.Write)
	Remove = Op(fsnotify.Remove)
	Rename = Op(fsnotify.Rename)
	Chmod  = Op(fsnotify.Chmod)
)

type Event struct {
	e *fsnotify.Event
}

func (e *Event) Path() string {
	return e.e.Name
}

func (e *Event) IsCreate() bool {
	return e.e.Op&fsnotify.Create == fsnotify.Create
}

func (e *Event) IsWrite() bool {
	return e.e.Op&fsnotify.Write == fsnotify.Write
}

func (e *Event) IsRemove() bool {
	return e.e.Op&fsnotify.Remove == fsnotify.Remove
}

func (e *Event) IsReName() bool {
	return e.e.Op&fsnotify.Rename == fsnotify.Rename
}

func (e *Event) IsChmod() bool {
	return e.e.Op&fsnotify.Chmod == fsnotify.Chmod
}

type WatcherManager struct {
	*fsnotify.Watcher
	C               chan int
	registeredFuncs map[string][]func(event Event)
}

func (w *WatcherManager) On(path string, callback func(event Event)) {
	if w.registeredFuncs[path] == nil {
		w.registeredFuncs[path] = []func(event Event){callback}
	} else {
		w.registeredFuncs[path] = append(w.registeredFuncs[path], callback)
	}
	w.Add(path)
}
func (w *WatcherManager) Close() {
	w.Watcher.Close()
	close(w.C)
}
func (w *WatcherManager) Start() error {
	watecher, err := fsnotify.NewWatcher()
	w.Watcher = watecher
	w.C = make(chan int)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case event := <-w.Watcher.Events:
				fns := w.registeredFuncs[event.Name]
				for _, k := range fns {
					defer util.Recover()
					k(Event{&event})
				}
			case err := <-w.Watcher.Errors:
				util.LogError(err)
			case <-w.C:
				return
			}
		}
	}()
	return nil
}

func NewWatcherManager() *WatcherManager {
	w := &WatcherManager{
		registeredFuncs: map[string][]func(event Event){},
	}
	return w
}

var Watcher *WatcherManager

func init() {
	var err error
	Watcher = NewWatcherManager()
	err = Watcher.Start()
	if err != nil {
		panic(err)
	}
}
