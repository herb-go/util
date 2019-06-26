package tomlconfig

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/herb-go/util"
	"github.com/herb-go/util/config"
)

const UnmarshalerNameTOML = "toml"

//Load load toml file and unmarshaler to interface.
//Return any error if rasied
func Load(file util.FileObject, v interface{}) error {
	return config.Load(UnmarshalerNameTOML, file, v)
}

//MustLoad load toml file and unmarshaler to interface.
//Panic if  any error rasied
func MustLoad(file util.FileObject, v interface{}) {
	err := Load(file, v)
	if err != nil {
		panic(err)
	}
}

//Save save interface to toml file
//Return any error if rasied
func Save(file util.FileObject, v interface{}) error {
	buffer := bytes.NewBuffer([]byte{})
	err := toml.NewEncoder(buffer).Encode(v)
	if err != nil {
		return err
	}
	return util.WriteFile(file, buffer.Bytes(), util.DefaultFileMode)
}

//MustSave save interface to toml file
//Panic if  any error rasied
func MustSave(file util.FileObject, v interface{}) {
	err := Save(file, v)
	if err != nil {
		panic(err)
	}
}

var tomlUnmarshal = func(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}

func init() {
	config.RegisterUnmarshaler(UnmarshalerNameTOML, tomlUnmarshal)
}
