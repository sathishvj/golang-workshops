package helloappengine

import (
	"appengine"
	"appengine/datastore"
	"html/template"
	"net/http"
	"time"
)

type Info struct {
	Data string
	Time time.Time
}

func init() {
	http.HandleFunc("/", onRoot)
	http.HandleFunc("/add", onAdd)
}

func onRoot(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl/root.tmpl")

	t.ExecuteTemplate(w, "root", nil)
}

func onAdd(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	s := r.FormValue("info")
	addInfo(c, s)

	onRoot(w, r)
}

func addInfo(c appengine.Context, s string) {
	key := datastore.NewIncompleteKey(c, "Info", nil)

	newInfo := Info{
		s,
		time.Now(),
	}
	datastore.Put(c, key, &newInfo)
}