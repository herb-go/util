package commonaction

import "net/http"

var successmsg = []byte("\"success\"")

//SuccessAction action which response a "success" msg.
var SuccessAction = func(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(successmsg)
	if err != nil {
		panic(err)
	}
}
