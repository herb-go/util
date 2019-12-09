package tomlconfig

import (
	_ "github.com/herb-go/herbconfig/configloader/drivers/tomlconfig" //tomlconfig
	"github.com/herb-go/herbconfig/configuration"
	"github.com/herb-go/util/config"
)

const UnmarshalerNameTOML = "toml"

//Load load toml file and unmarshaler to interface.
//Return any error if rasied
func Load(file configuration.Configuration, v interface{}) error {
	return config.Load(UnmarshalerNameTOML, file, v)
}

//MustLoad load toml file and unmarshaler to interface.
//Panic if  any error rasied
func MustLoad(file configuration.Configuration, v interface{}) {
	err := Load(file, v)
	if err != nil {
		panic(err)
	}
}
