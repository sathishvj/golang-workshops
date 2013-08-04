package helloappengine

import (
	"appengine"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", onRoot)
}

func onRoot(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	c.Debugf("Within the onRoot function ... ")

	fmt.Fprint(w, "Hello World!")

	c.Infof("Data successfully returned.")
}
