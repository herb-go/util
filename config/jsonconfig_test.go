package config

import (
	"strings"
	"testing"

	"github.com/herb-go/herbconfig/configuration"
)

func TestJSON(t *testing.T) {
	type DataStruct struct {
		Data string
	}
	var jdata = `//comment
	{"Data":"12345"}`
	var data = &DataStruct{}
	MustLoadJSON(configuration.Text(jdata), data)
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
			if !strings.Contains(err.Error(), configuration.Text(wrongdata).ID()) {
				t.Fatal(err)
			}
		}()
		MustLoadJSON(configuration.Text(wrongdata), data)
	}()
}
