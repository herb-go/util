package httpserver

import (
	"net/http"
)

//HeaderPanicID header for panic id.
var HeaderPanicID = "panic-id"

//HeaderPrivateRef header for private ref.
var HeaderPrivateRef = "private-ref"

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
				w.Header().Set(HeaderPrivateRef, ref)
			}
		}
	}
	return true
}

//PrivateRenderers private renders
var PrivateRenderers = []ErrorRenderer{
	PrivateRefRenderer,
}
