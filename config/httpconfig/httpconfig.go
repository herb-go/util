package httpconfig

import (
	"github.com/herb-go/herb/middleware/forwarded"
	"github.com/herb-go/herb/middleware/misc"
	"github.com/herb-go/herb/service/httpservice"
)

type Config struct {
	Forwarded forwarded.Middleware
	Config    httpservice.Config
	Headers   misc.Headers
}

func New() *Config {
	return &Config{
		Forwarded: forwarded.Middleware{},
		Config:    *httpservice.NewConfig(),
		Headers:   misc.Headers{},
	}
}
