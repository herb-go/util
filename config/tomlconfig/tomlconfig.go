package tomlconfig

import (
	"github.com/herb-go/herbconfig/source"
	_ "github.com/herb-go/herbconfig/loader/drivers/tomlconfig" //tomlconfig
	"github.com/herb-go/util/config"
)

const UnmarshalerNameTOML = "toml"

//Load load toml file and unmarshaler to interface.
//Return any error if rasied
func Load(file source.Source, v interface{}) error {
	return config.Load(UnmarshalerNameTOML, file, v)
}

//MustLoad load toml file and unmarshaler to interface.
//Panic if  any error rasied
func MustLoad(file source.Source, v interface{}) {
	err := Load(file, v)
	if err != nil {
		panic(err)
	}
}
