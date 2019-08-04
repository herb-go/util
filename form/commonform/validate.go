package commonform

import (
	"strconv"

	"github.com/herb-go/herb/ui"
	"github.com/herb-go/util/form"
)

//SetTranslatedFieldLabels set translated field labels to form.
func SetTranslatedFieldLabels(form form.TranslateableForm, module string, labels map[string]string) {
	form.SetFieldLabelsCollection(ui.GetMessages(form.Lang(), "app").Collection((labels)))
}

//ValidateStringLength validate string length
func ValidateStringLength(form form.TranslateableForm, value string, field string, min int, max int) {
	l := len(value)
	if min == max {
		if l != min {
			form.AddErrorf(field, NewWrongLengthMsg(strconv.Itoa(min)).Translate(form.Lang()))
		}
	} else {
		if l < min {
			form.AddErrorf(field, NewTooShortMsg(strconv.Itoa(min)).Translate(form.Lang()))
		} else if l > max {
			form.AddErrorf(field, NewTooLongMsg(strconv.Itoa(min)).Translate(form.Lang()))
		}
	}
}

//ValidateIntRange validate int value range
func ValidateIntRange(form form.TranslateableForm, value int, field string, min int, max int) {
	if value < min {
		form.AddErrorf(field, NewTooSmallMsg(strconv.Itoa(min)).Translate(form.Lang()))

	} else if value > max {
		form.AddErrorf(field, NewTooBigMsg(strconv.Itoa(min)).Translate(form.Lang()))
	}
}

//ValidateInt64Range validate int value range
func ValidateInt64Range(form form.TranslateableForm, value int64, field string, min int64, max int64) {
	if value < min {
		form.AddErrorf(field, NewTooSmallMsg(strconv.FormatInt(min, 10)).Translate(form.Lang()))

	} else if value > max {
		form.AddErrorf(field, NewTooBigMsg(strconv.FormatInt(max, 10)).Translate(form.Lang()))
	}
}

//ValidateRequiredPointer validate required pointer field
func ValidateRequiredPointer(form form.TranslateableForm, value interface{}, field string) {
	if value == nil {
		form.AddErrorf(field, MsgRequired.Translate(form.Lang()))
	}
}

//ValidateRequiredString validate required string field
func ValidateRequiredString(form form.TranslateableForm, value string, field string) {
	if value == "" {
		form.AddErrorf(field, MsgRequired.Translate(form.Lang()))
	}
}

//ValidateRequiredInt validate required int field
func ValidateRequiredInt(form form.TranslateableForm, value int, field string) {
	if value == 0 {
		form.AddErrorf(field, MsgRequired.Translate(form.Lang()))
	}
}

//ValidateRequiredInt64 validate required int64 field
func ValidateRequiredInt64(form form.TranslateableForm, value int64, field string) {
	if value == 0 {
		form.AddErrorf(field, MsgRequired.Translate(form.Lang()))
	}
}
