package helloappengine

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", onRoot)
}

func onRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
