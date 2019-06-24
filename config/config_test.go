package config

import "testing"
import "github.com/herb-go/util"

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
}
