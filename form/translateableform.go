package form

import (
	"github.com/herb-go/herb/model"
	"github.com/herb-go/herb/translate"
)

//TranslateableForm translateable form interaface
type TranslateableForm interface {
	model.Validator
	translate.TranslationLanguage
}
