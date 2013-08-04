package helloappengine

import (
	"appengine"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", onRoot)
}

func onRoot(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// TODO: call ParseFiles function of template package.  The file to be passed is "tmpl/root.tmpl".
	// It returns the parsed template (call it "t", and an error which you can call "err")

	if err != nil {
		c.Errorf("Error while parsing templates: " + err.Error())
		return
	}

	// TODO: call ExecuteTemplate on the parsed template file.  The template name is "root" and there is no data to pass to it.
	err = t.  (w, "root", nil)
}
