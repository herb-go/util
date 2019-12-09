package config

import (
	"github.com/herb-go/herbconfig/configuration"
	"github.com/herb-go/herbconfig/configuration/watchers/fsnotifywatcher"
	"github.com/herb-go/util"
)

var WatcherManager = configuration.NewWatchManager()

func InitWatcherManager() {
}

func init() {
	util.Must(WatcherManager.RegisterWatcher(configuration.NewSchemeWatcher("file", fsnotifywatcher.NewWatcher())))
	util.Must(WatcherManager.RegisterWatcher(configuration.NewSchemeWatcher("relative", fsnotifywatcher.NewWatcher())))
}
