package formmessages

import (
	"github.com/herb-go/herb/translate"
)

//MsgRequired message that shows field is required
var MsgRequired = translate.NewMessage("herbgo.form", "%[1]s required")

var MsgExist = translate.NewMessage("herbgo.form", " %[1]s existed")

var MsgFormatError = translate.NewMessage("herbgo.form", " %[1]s format error")

var MsgTooSmall = translate.NewMessage("herbgo.form", " %[1]s too small")

var MsgTooLarge = translate.NewMessage("herbgo.form", " %[1]s too large")

var MsgTooLong = translate.NewMessage("herbgo.form", " %[1]s too long")

var MsgTooShort = translate.NewMessage("herbgo.form", " %[1]s too short")

var MsgNotAvaliable = translate.NewMessage("herbgo.form", " %[1]s is not avaliable")

var MsgNotMatch = translate.NewMessage("herbgo.form", "%[1]s is not match")
