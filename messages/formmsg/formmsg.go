package formmessages

import (
	"github.com/herb-go/herb/translate"
)

//MsgRequired message that shows field is required
var MsgRequired = translate.NewMessage("form", "%[1]s required")

var MsgNotUnique = translate.NewMessage("form", " %[1]s not unique")
