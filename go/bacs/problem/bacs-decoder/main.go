package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var listenOn = flag.String("listen", ":80", "Listen on host:port")

var templates = template.Must(template.ParseFiles(
	"templates/broker_index.html",
	"templates/footer.html",
	"templates/header.html",
	"templates/index.html",
	"templates/raw_index.html",
))

func handleTemplate(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, name, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func main() {
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", handleTemplate("index.html"))
	r.HandleFunc("/raw", handleTemplate("raw_index.html"))
	r.HandleFunc("/raw/task", Render(RawTaskRenderer))
	r.HandleFunc("/raw/result", Render(RawResultRenderer))
	r.HandleFunc("/broker", handleTemplate("broker_index.html"))
	r.HandleFunc("/broker/task", Render(BrokerTaskRenderer))
	r.HandleFunc("/broker/result", Render(BrokerResultRenderer))
	r.HandleFunc("/broker/status", Render(BrokerStatusRenderer))

	log.Fatal(http.ListenAndServe(*listenOn, r))
}
