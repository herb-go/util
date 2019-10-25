package httpserver

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

//Config http server config.
type Config struct {
	//Net net interface,"tcp" for example.
	Net string
	//Addr network addr.
	Addr string
	//BaseURL http scheme and host."http://127.0.0.1:8000" for example.
	BaseURL string
	//ReadTimeoutInSecond http conn read time out.
	ReadTimeoutInSecond int64
	//ReadTimeoutInSecond http conn read Header time out.
	ReadHeaderTimeoutInSecond int64
	//WriteTimeoutInSecond http conn write time out.
	WriteTimeoutInSecond int64
	//IdleTimeoutInSecond conn idle time out.
	IdleTimeoutInSecond int64
	//MaxHeaderBytes max header length in bytes.
	MaxHeaderBytes int
	//TLS whether use tls
	TLS bool
	//TLSCertPath tls cert file path
	TLSCertPath string
	//TLSKeyPath tls key file path
	TLSKeyPath string
}

//IsEmpty check if config is empty
func (c *Config) IsEmpty() bool {
	if c == nil {
		return true
	}
	if c.Addr != "" {
		return false
	}
	if c.Net != "" {
		return false
	}
	if c.BaseURL != "" {
		return false
	}
	if c.ReadTimeoutInSecond != 0 {
		return false
	}
	if c.ReadHeaderTimeoutInSecond != 0 {
		return false
	}
	if c.WriteTimeoutInSecond != 0 {
		return false
	}
	if c.IdleTimeoutInSecond != 0 {
		return false
	}
	if c.MaxHeaderBytes != 0 {
		return false
	}
	if !c.TLS {
		return false
	}
	if c.TLSCertPath != "" {
		return false
	}
	if c.TLSKeyPath != "" {
		return false
	}
	return true
}

//Server create http server with config.
func (c *Config) Server() *http.Server {
	server := &http.Server{
		Addr:              c.Addr,
		ReadTimeout:       time.Duration(c.ReadTimeoutInSecond) * time.Second,
		ReadHeaderTimeout: time.Duration(c.ReadHeaderTimeoutInSecond) * time.Second,
		WriteTimeout:      time.Duration(c.WriteTimeoutInSecond) * time.Second,
		IdleTimeout:       time.Duration(c.IdleTimeoutInSecond) * time.Second,
		MaxHeaderBytes:    c.MaxHeaderBytes,
	}
	server.ErrorLog = log.New(ioutil.Discard, "", 0)
	return server
}

//Listen listen net and addr in config.
//Return net listener and any error if raised.
func (c *Config) Listen() (net.Listener, error) {
	return net.Listen(c.Net, c.Addr)
}

//MustListen listen net and addr in config.
//Return net listener.
//Panic if any error raised.
func (c *Config) MustListen() net.Listener {
	l, err := c.Listen()
	if err != nil {
		panic(err)
	}
	return l
}
