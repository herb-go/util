package httpconfig

import (
	"io/ioutil"
	"net/http"

	"github.com/herb-go/herb/middleware/errorpage"

	"github.com/herb-go/herb/middleware/forwarded"
	"github.com/herb-go/herb/middleware/misc"
	"github.com/herb-go/herb/service"
	"github.com/herb-go/herb/service/httpservice"
)

type Config struct {
	Forwarded  forwarded.Middleware
	Config     httpservice.Config
	Headers    misc.Headers
	Hosts      service.Hosts
	ErrorPages ErrorPagesConfig
}

func New() *Config {
	return &Config{
		Forwarded:  forwarded.Middleware{},
		Config:     *httpservice.NewConfig(),
		Headers:    misc.Headers{},
		Hosts:      service.Hosts{},
		ErrorPages: ErrorPagesConfig{},
	}
}

type StatusCodePage struct {
	StatusCode int
	Path       string
}
type ErrorPagesConfig struct {
	ErrorPage          string
	StatusCodePages    []*StatusCodePage
	IgnoredStatusCodes []int
}

func renderFile(p string) func(w http.ResponseWriter, r *http.Request, statuscode int) {
	return func(w http.ResponseWriter, r *http.Request, statuscode int) {
		w.WriteHeader(statuscode)
		bs, err := ioutil.ReadFile(p)
		if err != nil {
			panic(err)
		}
		_, err = w.Write(bs)
		if err != nil {
			panic(err)
		}

	}
}

func (c *ErrorPagesConfig) Middleware() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	e := errorpage.New()
	if c.ErrorPage != "" {
		e.OnError(renderFile(c.ErrorPage))
	}
	for _, v := range c.StatusCodePages {
		if v.Path != "" {
			e.OnStatus(v.StatusCode, renderFile(v.Path))
		}
	}
	for _, v := range c.IgnoredStatusCodes {
		e.IgnoreStatus(v)
	}
	return e.ServeMiddleware
}
