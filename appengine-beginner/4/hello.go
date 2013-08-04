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

	t, err := template.ParseFiles("tmpl/root.tmpl")
	if err != nil {
		c.Errorf("Error while parsing templates: " + err.Error())
		return
	}

	err = t.ExecuteTemplate(w, "root", nil)
}
