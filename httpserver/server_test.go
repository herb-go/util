package httpserver

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/herb-go/herb/server"
	"github.com/herb-go/util"
)

func TestConfig(t *testing.T) {
	config := &server.HTTPConfig{}
	config.Net = "tcp"
	config.Addr = httpaddr
}
func TestHttp(t *testing.T) {
	output := bytes.NewBuffer([]byte{})
	util.Logger.SetOutput(output)
	defer func() {
		util.Logger.SetOutput(os.Stderr)
	}()
	config := &server.HTTPConfig{}
	config.Net = "tcp"
	config.Addr = httpaddr
	server := config.Server()
	if server.ReadTimeout != time.Duration(config.ReadTimeoutInSecond)*time.Second {
		t.Fatal(server.ReadTimeout)
	}
	if server.ReadHeaderTimeout != time.Duration(config.ReadHeaderTimeoutInSecond)*time.Second {
		t.Fatal(server.ReadHeaderTimeout)
	}
	if server.WriteTimeout != time.Duration(config.WriteTimeoutInSecond)*time.Second {
		t.Fatal(server.WriteTimeout)
	}
	if server.IdleTimeout != time.Duration(config.IdleTimeoutInSecond)*time.Second {
		t.Fatal(server.IdleTimeout)
	}
	if server.MaxHeaderBytes != config.MaxHeaderBytes {
		t.Fatal(server.MaxHeaderBytes)
	}
	app := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		RecoverMiddleware(nil)(w, r, func(w http.ResponseWriter, r *http.Request) {
			panic(errors.New("errorcontent"))
		})
	})
	MustListenAndServeHTTP(server, config, app)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	time.Sleep(time.Second)
	resp, err := client.Get("http://" + config.Addr + "/")
	if err != nil {
		t.Fatal(err)
	}
	bodycontent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(bodycontent), "errorcontent") {
		t.Fatal(string(bodycontent))
	}
	util.Debug = true
	defer func() {
		util.Debug = false
	}()
	resp, err = client.Get("http://" + config.Addr + "/")
	if err != nil {
		t.Fatal(err)
	}
	bodycontent, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(bodycontent), "errorcontent") {
		t.Fatal(string(bodycontent))
	}

	resp.Body.Close()
	ShutdownHTTP(server)
}

func TestHttps(t *testing.T) {
	config := &server.HTTPConfig{}
	config.Net = "tcp"
	config.Addr = httpaddr
	config.TLS = true
	config.TLSCertPath = "./testdata/server.crt"
	config.TLSKeyPath = "./testdata/server.key"
	server := config.Server()
	app := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })
	MustListenAndServeHTTP(server, config, app)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	time.Sleep(time.Second)
	resp, err := client.Get("https://" + config.Addr + "/")
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	ShutdownHTTPWithTimeout(server, time.Second)
}
