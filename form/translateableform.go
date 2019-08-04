package form

import (
	"github.com/herb-go/herb/ui/validator"
	"github.com/herb-go/herb/ui"
)

//TranslateableForm translateable form interaface
type TranslateableForm interface {
	model.Validator
	ui.TranslationLanguage
}
