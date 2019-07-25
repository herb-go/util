package config

import (
	"errors"
	"fmt"
)

type ConfigError struct {
	ConfigPath string
	RawErrror  error
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("load config %s fail.%s", e.ConfigPath, e.RawErrror)
}
func NewError(configPath string, rawError error) *ConfigError {
	return &ConfigError{
		ConfigPath: configPath,
		RawErrror:  rawError,
	}
}

var ErrConfigPathNotInited = errors.New("config path not inited")
