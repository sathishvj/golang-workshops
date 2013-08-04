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
	c := appengine.NewContext(r)

	t, _ := template.ParseFiles("tmpl/root.tmpl")

	infos := getInfos(c)
	t.ExecuteTemplate(w, "root", infos)
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

func getInfos(c appengine.Context) []Info {
	q := datastore.NewQuery("Info").Order("-Time")

	var infos []Info
	q.GetAll(c, &infos)

	return infos
}
