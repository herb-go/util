package tools

import (
	"fmt"
)

func ErrorIfStringFieldNotInList(filename string, value string, valuelist ...string) error {
	for _, v := range valuelist {
		if value == v {
			return nil
		}
	}
	return fmt.Errorf("arg %s value \"%s\" in not avaliable", filename, value)
}

func ErrorIfIntFieldNotInRange(filename string, value int, min int, max int) error {
	if value >= min && value <= max {
		return nil
	}
	return fmt.Errorf("arg %s value \"%d\" in not avaliable", filename, value)
}

func ErrorIfInt64FieldNotInRange(filename string, value int64, min int64, max int64) error {
	if value >= min && value <= max {
		return nil
	}
	return fmt.Errorf("arg %s value \"%d\" in not avaliable", filename, value)
}
