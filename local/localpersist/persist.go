package localpersist

import (
	"github.com/herb-go/herb/persist"
	"github.com/herb-go/util"
)

var Factory = func(loader func(v interface{}) error) (persist.Store, error) {
	return persist.FolderStore(util.AppData("persistdata")), nil
}
