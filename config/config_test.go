package config

import "testing"
import "github.com/herb-go/util"
import "strings"

func TestJSON(t *testing.T) {
	type DataStruct struct {
		Data string
	}
	var jdata = `//comment
	{"Data":"12345"}`
	var data = &DataStruct{}
	MustLoadJSON(util.FileObjectText(jdata), data)
	if data.Data != "12345" {
		t.Fatal(data)
	}
	var wrongdata = `//comment
	{"Data":"12345}`
	func() {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatal(r)
			}
			err := r.(error)
			if !strings.Contains(err.Error(), util.FileObjectText(wrongdata).ID()) {
				t.Fatal(err)
			}
		}()
		MustLoadJSON(util.FileObjectText(wrongdata), data)
	}()
}
