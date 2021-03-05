package httpserver

import (
	"fmt"
	"net/http"
)

//ErrorRenderer error http render.
//Return false if render finished.
type ErrorRenderer func(w http.ResponseWriter, r *http.Request, err error) bool

//PrivateRefError private ref error
type PrivateRefError interface {
	//ErrorPrivateRef error private ref
	ErrorPrivateRef() string
}

//PrivateRefRenderer error private ref render
func PrivateRefRenderer(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {
		if e, ok := err.(PrivateRefError); ok {
			ref := e.ErrorPrivateRef()
			if ref != "" {
				msg := fmt.Sprintf("%s\nRef:%s", http.StatusText(500), ref)
				http.Error(w, msg, 500)
				return false
			}
		}
	}
	return true
}

//PrivateRenderers private renders
var PrivateRenderers = []ErrorRenderer{
	PrivateRefRenderer,
}
