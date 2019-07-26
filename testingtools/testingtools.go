package testingtools

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/herb-go/util"
)

//SetRootPathRelativeToModules set rootoath relative to modules folder.
func SetRootPathRelativeToModules(paths ...string) {
	var rootpath string
	_, file, _, _ := runtime.Caller(1)
	files := []string{file, "../"}
	files = append(files, paths...)
	modulespath := filepath.Join(files...)
	vendormodulepath := filepath.Join("vendor", "modules")
	if strings.HasSuffix(modulespath, vendormodulepath) {
		rootpath = filepath.Join(modulespath, "../", "../", "../")
	} else {
		rootpath = filepath.Join(modulespath, "../", "../")
	}
	util.RootPath = rootpath
}
