package config

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

var unmarshalers = map[string]func(data []byte, v interface{}) error{}

func RegisterUnmarshaler(name string, driver func(data []byte, v interface{}) error) {
	lock.Lock()
	defer lock.Unlock()
	unmarshalers[name] = driver
}

func Unmarshal(drivername string, data []byte, v interface{}) error {
	unmarshaler := unmarshalers[drivername]
	if unmarshaler == nil {
		return fmt.Errorf("config unmarshaler \"%s\" not found", drivername)
	}
	return unmarshaler(data, v)
}
