package config

import (
	_ "github.com/herb-go/herbconfig/loader/types/timedurationtype" //time duration type
	"github.com/herb-go/herbconfig/source"
	"github.com/herb-go/herbconfig/source/watchers/fsnotifywatcher"
	"github.com/herb-go/util"
)

var WatcherManager = source.NewWatchManager()

func InitWatcherManager() {
}

func init() {
	util.Must(WatcherManager.RegisterWatcher(source.NewSchemeWatcher("file", fsnotifywatcher.NewWatcher())))
	util.Must(WatcherManager.RegisterWatcher(source.NewSchemeWatcher("relative", fsnotifywatcher.NewWatcher())))
}
