package config

import (
	"github.com/herb-go/herbconfig/loader"
	_ "github.com/herb-go/herbconfig/loader/drivers/jsonconfig" //jsonconfig
	"github.com/herb-go/herbconfig/source"
)

const UnmarshalerNameJSON = "json"

func Load(drivername string, file source.Source, v interface{}) error {
	bs, err := source.Read(file)
	if err != nil {
		return NewError(file.ID(), err)
	}
	err = loader.LoadConfig(drivername, bs, v)
	if err != nil {
		return NewError(file.ID(), err)
	}
	return nil

}
func LoadJSON(file source.Source, v interface{}) error {
	return Load(UnmarshalerNameJSON, file, v)
}
func MustLoadJSON(file source.Source, v interface{}) {
	err := LoadJSON(file, v)
	if err != nil {
		panic(err)
	}
}

func init() {

}
