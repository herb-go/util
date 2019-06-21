package util

import (
	"fmt"
	"testing"
)

func TestUtil(t *testing.T) {
	func() {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal(r)
			}
			err := r.(error)
			if err == nil {
				t.Fatal(err)
			}
		}()
		Must(fmt.Errorf("%s", "testerror"))
	}()
}
