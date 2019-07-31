package commonform

import (
	"github.com/herb-go/herb/translate"
)

//MsgRequired message that shows field is required
var MsgRequired = translate.NewMessage("herbgo.form", "%[1]s is required")

//MsgDuplicated messages that shows field is  duplicated.
var MsgDuplicated = translate.NewMessage("herbgo.form", " %[1]s is duplicated.")

//MsgInvalide messages that shows field is  invalid.
var MsgInvalide = translate.NewMessage("herbgo.form", " %[1]s is invalid")

//MsgNotMatch messages that shows field is  not match.
var MsgNotMatch = translate.NewMessage("herbgo.form", "%[1]s is not match")

//MsgIncorrect messages that shows field is  incorrect.
var MsgIncorrect = translate.NewMessage("herbgo.form", "%[1]s is incorrect")

//NewFormatWrongMsg create new message that shows field format wrong with given format.
func NewFormatWrongMsg(format string) *translate.TemplateMessage {
	return translate.NewTemplateMessage("herbgo.form", "%[1]s format wrong(correct format is '{{format}}'.", map[string]string{"format": format})
}

//NewTooSmallMsg create new message that shows field is too small with given min value.
func NewTooSmallMsg(min string) *translate.TemplateMessage {
	return translate.NewTemplateMessage("herbgo.form", "%[1]s is too small (minimum is {{min}}).", map[string]string{"min": min})
}

//NewTooBigMsg create new message that shows field is too big with given max value.
func NewTooBigMsg(max string) *translate.TemplateMessage {
	return translate.NewTemplateMessage("herbgo.form", "%[1]s is too big (maximum is {{max}}).", map[string]string{"max": max})
}

//NewTooShortMsg create new message that shows field is too show with given min length.
func NewTooShortMsg(min string) *translate.TemplateMessage {
	return translate.NewTemplateMessage("herbgo.form", "%[1]s is too short (minimum is {{min}} characters).", map[string]string{"min": min})
}

//NewTooLongMsg create new message that shows field is too long with given max length.
func NewTooLongMsg(max string) *translate.TemplateMessage {
	return translate.NewTemplateMessage("herbgo.form", "%[1]s is too long (maximum is {{max}} characters).", map[string]string{"max": max})
}

//NewWrongLengthMsg create new message that shows field length wrong with given  length.
func NewWrongLengthMsg(length string) *translate.TemplateMessage {
	return translate.NewTemplateMessage("herbgo.form", " %[1]s is of the wrong length (should be {{length}} characters).", map[string]string{"length": length})
}
