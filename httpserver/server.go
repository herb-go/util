package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/herb-go/herb/service/httpservice"
	"github.com/herb-go/util"
)

//MustListenAndServeHTTP listen and serve http server with given server,config and handler.
//Panic if any error raised.
func MustListenAndServeHTTP(s *http.Server, config *httpservice.Config, app http.Handler) {
	go func() {
		var err error
		l, err := config.Listen()
		if err != nil {
			panic(err)
		}
		defer l.Close()
		s.Handler = app
		if config.TLS {
			util.Println("Listening https " + l.Addr().String())
			err = s.ServeTLS(l, config.TLSCertPath, config.TLSKeyPath)
		} else {
			util.Println("Listening " + l.Addr().String())
			err = s.Serve(l)
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

// ShutdownHTTP  shutdown  http server.
func ShutdownHTTP(Server *http.Server) {
	WithContextShutdown(context.Background(), Server)
}

//ShutdownHTTPWithTimeout shutdown  http server ith given timeout.
func ShutdownHTTPWithTimeout(Server *http.Server, Timeout time.Duration) {
	ctx, fn := context.WithTimeout(context.Background(), Timeout)
	fn()
	WithContextShutdown(ctx, Server)

}

//WithContextShutdown shutdown  http server ith given context.
func WithContextShutdown(ctx context.Context, Server *http.Server) {
	util.Println("Http server shuting down...")
	Server.Shutdown(ctx)
	util.Println("Http server Stoped.")
}
