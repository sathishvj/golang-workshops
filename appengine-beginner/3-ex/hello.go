package helloappengine

import (
	// TODO: add the "appengine" package here

	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", onRoot)
}

func onRoot(w http.ResponseWriter, r *http.Request) {
	// TODO: create a new context c, by calling NewContext function of appengine package.  It takes a *http.Request as a parameter.

	// TODO: call the Debugf function on the context
	c. ("Within the onRoot function ... ")

	fmt.Fprint(w, "Hello World!")

	// TODO: call the Infof function on the context
	c. ("Data successfully returned.")
}
