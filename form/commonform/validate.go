package commonform

import (
	"reflect"
	"strconv"
	"time"

	"github.com/herb-go/herb/ui/validator"

	"github.com/herb-go/herb/ui"
)

//SetTranslatedFieldLabels set translated field labels to form.
func SetTranslatedFieldLabels(form validator.Fields, module string, labels map[string]string) {
	form.SetComponentLabels(ui.GetMessages(form.Lang(), "app").Collection((labels)))
}

//ValidateStringLength validate string length
func ValidateStringLength(form validator.Fields, value string, field string, min int, max int) {
	l := len(value)
	if min == max {
		if l != min {
			form.AddErrorf(field, NewWrongLengthMsg(strconv.Itoa(min)).Translate(form.Lang()))
		}
	} else {
		if l < min {
			form.AddErrorf(field, NewTooShortMsg(strconv.Itoa(min)).Translate(form.Lang()))
		} else if l > max {
			form.AddErrorf(field, NewTooLongMsg(strconv.Itoa(max)).Translate(form.Lang()))
		}
	}
}

//ValidateIntRange validate int value range
func ValidateIntRange(form validator.Fields, value int, field string, min int, max int) {
	if value < min {
		form.AddErrorf(field, NewTooSmallMsg(strconv.Itoa(min)).Translate(form.Lang()))
	} else if value > max {
		form.AddErrorf(field, NewTooBigMsg(strconv.Itoa(max)).Translate(form.Lang()))
	}
}

//ValidateInt64Range validate int value range
func ValidateInt64Range(form validator.Fields, value int64, field string, min int64, max int64) {
	if value < min {
		form.AddErrorf(field, NewTooSmallMsg(strconv.FormatInt(min, 10)).Translate(form.Lang()))

	} else if value > max {
		form.AddErrorf(field, NewTooBigMsg(strconv.FormatInt(max, 10)).Translate(form.Lang()))
	}
}

//ValidateRequiredPointer validate required pointer field
func ValidateRequiredPointer(form validator.Fields, value interface{}, field string) {
	if value == nil || reflect.ValueOf(value).IsNil() {
		form.AddErrorf(field, MsgRequired.Translate(form.Lang()))
	}
}
func getLength(value interface{}) int {
	if value != nil {
		return reflect.Indirect(reflect.ValueOf(value)).Len()
	}
	return 0
}

//ValidateRequiredList validate required list field
func ValidateRequiredList(form validator.Fields, value interface{}, field string) {
	if getLength(value) == 0 {
		form.AddErrorf(field, MsgRequired.Translate(form.Lang()))
	}
}

//ValidateListLength validate list length
func ValidateListLength(form validator.Fields, value interface{}, field string, min int, max int) {
	l := getLength(value)
	if min == max {
		if l != min {
			form.AddErrorf(field, NewWrongLengthMsg(strconv.Itoa(min)).Translate(form.Lang()))
		}
	} else {
		if l < min {
			form.AddErrorf(field, NewTooShortMsg(strconv.Itoa(min)).Translate(form.Lang()))
		} else if l > max {
			form.AddErrorf(field, NewTooLongMsg(strconv.Itoa(max)).Translate(form.Lang()))
		}
	}
}

//ValidateRequiredString validate required string field
func ValidateRequiredString(form validator.Fields, value string, field string) {
	if value == "" {
		form.AddErrorf(field, MsgRequired.Translate(form.Lang()))
	}
}

//ValidateRequiredInt validate required int field
func ValidateRequiredInt(form validator.Fields, value int, field string) {
	if value == 0 {
		form.AddErrorf(field, MsgRequired.Translate(form.Lang()))
	}
}

//ValidateRequiredInt64 validate required int64 field
func ValidateRequiredInt64(form validator.Fields, value int64, field string) {
	if value == 0 {
		form.AddErrorf(field, MsgRequired.Translate(form.Lang()))
	}
}

//ValidateRequiredFloat32 validate required int field
func ValidateRequiredFloat32(form validator.Fields, value float32, field string) {
	if value == 0 {
		form.AddErrorf(field, MsgRequired.Translate(form.Lang()))
	}
}

//ValidateRequiredFloat64 validate required int64 field
func ValidateRequiredFloat64(form validator.Fields, value float64, field string) {
	if value == 0 {
		form.AddErrorf(field, MsgRequired.Translate(form.Lang()))
	}
}

//ValidateRequiredTime validate required int64 field
func ValidateRequiredTime(form validator.Fields, value time.Time, field string) {
	if value.IsZero() {
		form.AddErrorf(field, MsgRequired.Translate(form.Lang()))
	}
}

//ValidateStringInList validate required string field
func ValidateStringInList(form validator.Fields, value string, field string, l *StringList) {
	if !l.Has(value) {
		form.AddErrorf(field, MsgInvalide.Translate(form.Lang()))
	}
}

//ValidateIntInList validate required string field
func ValidateIntInList(form validator.Fields, value int, field string, l *IntList) {
	if !l.Has(value) {
		form.AddErrorf(field, MsgInvalide.Translate(form.Lang()))
	}
}

//ValidateInt64InList validate required string field
func ValidateInt64InList(form validator.Fields, value int64, field string, l *Int64List) {
	if !l.Has(value) {
		form.AddErrorf(field, MsgInvalide.Translate(form.Lang()))
	}
}
