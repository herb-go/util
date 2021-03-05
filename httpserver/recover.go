package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/herb-go/util"
)

//HeaderPanicID header for panic id.
var HeaderPanicID = "panicid"

//CreateRecoverMiddleware create recover middleware by given logger and renders.
func CreateRecoverMiddleware(logger *log.Logger, renderers []ErrorRenderer) func(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	return func(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		defer func() {
			if r := recover(); r != nil {
				err := r.(error)
				var result string
				if util.IsErrorIgnored(err) == false {
					lines := strings.Split(string(debug.Stack()), "\n")
					length := len(lines)
					maxLength := util.LoggerMaxLength*2 + 7
					if length > maxLength {
						length = maxLength
					}
					var output = make([]string, length-6)
					panicid := util.IDGenerator()
					var panicidinfo string
					if panicid != "" {
						if HeaderPanicID != "" {
							w.Header().Set(HeaderPanicID, panicid)
						}
						panicidinfo = fmt.Sprintf("[PanicID:%s] - ", panicid)
					}
					output[0] = fmt.Sprintf("Panic: %s%s - http request %s \"%s\" ", panicidinfo, err.Error(), req.Method, req.URL.String())
					output[0] += "\n" + lines[0]
					copy(output[1:], lines[7:])
					result = strings.Join(output, "\n")
					if logger != nil {
						logger.Println(result)
					} else {
						util.ErrorLogger(result)
					}
				}
				if util.Debug {
					http.Error(w, result, http.StatusInternalServerError)
				} else {
					for _, v := range renderers {
						notfinished := !v(w, req, err)
						if !notfinished {
							return
						}
					}
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}
		}()
		next(w, req)
	}
}

//RecoverMiddleware create recover middleware with given logger.
func RecoverMiddleware(logger *log.Logger) func(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	return CreateRecoverMiddleware(logger, nil)
}

//PrivateRecoverMiddleware create private recover middleware with given logger.
func PrivateRecoverMiddleware(logger *log.Logger) func(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	return CreateRecoverMiddleware(logger, PrivateRenderers)
}
