package form

import (
	"github.com/herb-go/herb/ui/userinput"
	"github.com/herb-go/herb/ui"
)

//TranslateableForm translateable form interaface
type TranslateableForm interface {
	model.Validator
	translate.TranslationLanguage
}
