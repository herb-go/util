package httpconfig

import (
	"github.com/herb-go/herb/middleware/forwarded"
	"github.com/herb-go/herb/middleware/misc"
	"github.com/herb-go/herb/service"
	"github.com/herb-go/herb/service/httpservice"
)

type Config struct {
	Forwarded forwarded.Middleware
	Config    httpservice.Config
	Headers   misc.Headers
	Hosts     service.Hosts
}

func New() *Config {
	return &Config{
		Forwarded: forwarded.Middleware{},
		Config:    *httpservice.NewConfig(),
		Headers:   misc.Headers{},
		Hosts:     service.Hosts{},
	}
}
