package config

import (
	"github.com/herb-go/herbconfig/configloader"
	_ "github.com/herb-go/herbconfig/configloader/drivers/jsonconfig" //jsonconfig
	"github.com/herb-go/herbconfig/configuration"
)

const UnmarshalerNameJSON = "json"

func Load(drivername string, file configuration.Configuration, v interface{}) error {
	bs, err := configuration.Read(file)
	if err != nil {
		return NewError(file.ID(), err)
	}
	err = configloader.LoadConfig(drivername, bs, v)
	if err != nil {
		return NewError(file.ID(), err)
	}
	return nil

}
func LoadJSON(file configuration.Configuration, v interface{}) error {
	return Load(UnmarshalerNameJSON, file, v)
}
func MustLoadJSON(file configuration.Configuration, v interface{}) {
	err := LoadJSON(file, v)
	if err != nil {
		panic(err)
	}
}

func init() {

}
