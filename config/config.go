package config

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/herb-go/util"
)

func LoadJSON(file util.FileObject, v interface{}) error {

	bs, err := util.ReadFile(file)
	if err != nil {
		return NewError(file.ID(), err)
	}
	r := bytes.NewBuffer(bs)
	var bytes = []byte{}
	err = nil
	var line string
	for err != io.EOF {
		line, err = r.ReadString(10)
		line = strings.TrimSpace(line)
		if len(line) > 2 && line[0:2] == "//" {
			continue
		}
		bytes = append(bytes, []byte(line)...)
	}
	err = json.Unmarshal(bytes, v)
	if err != nil {
		return NewError(file.ID(), err)
	}
	return nil
}
func MustLoadJSON(file util.FileObject, v interface{}) {
	err := LoadJSON(file, v)
	if err != nil {
		panic(err)
	}
}
