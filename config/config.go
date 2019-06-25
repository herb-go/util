package config

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"

	"github.com/herb-go/util"
)

const UnmarshalerNameJSON = "json"

func Load(drivername string, file util.FileObject, v interface{}) error {
	bs, err := util.ReadFile(file)
	if err != nil {
		return NewError(file.ID(), err)
	}
	err = Unmarshal(drivername, bs, v)
	if err != nil {
		return NewError(file.ID(), err)
	}
	return nil

}
func LoadJSON(file util.FileObject, v interface{}) error {
	return Load(UnmarshalerNameJSON, file, v)
}
func MustLoadJSON(file util.FileObject, v interface{}) {
	err := LoadJSON(file, v)
	if err != nil {
		panic(err)
	}
}

var jsonUnmarshal = func(data []byte, v interface{}) error {
	var err error
	r := bytes.NewBuffer(data)
	var bytes = []byte{}
	var line string
	for err != io.EOF {
		line, err = r.ReadString(10)
		line = strings.TrimSpace(line)
		if len(line) > 2 && line[0:2] == "//" {
			continue
		}
		bytes = append(bytes, []byte(line)...)
	}
	return json.Unmarshal(bytes, v)
}

func init() {
	RegisterUnmarshaler(UnmarshalerNameJSON, jsonUnmarshal)
}
