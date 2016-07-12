package main

import (
	"flag"
	"html/template"
	"net/http"

	"github.com/bacsorg/problems/go/bacs/problem/decoder"
	"github.com/urfave/negroni"
)

var listenOn = flag.String("listen", ":80", "Listen on host:port")

var templates = template.Must(template.ParseFiles(
	"templates/decode.html",
	"templates/footer.html",
	"templates/header.html",
	"templates/index.html",
	"templates/result.html",
	"templates/single_index.html",
	"templates/single_result.html",
	"templates/single_task.html",
))

type Result struct {
	Name     string
	Protobuf string
	Back     string
}

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/single", func(w http.ResponseWriter, req *http.Request) {
		err := templates.ExecuteTemplate(w, "single_index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/single/task",
		func(w http.ResponseWriter, req *http.Request) {
			var err error
			if req.Method == "POST" {
				var text string
				text, err = decoder.SingleTaskDecoder.DecodeBase64ToText(
					req.FormValue("base64data"))
				if err == nil {
					err = templates.ExecuteTemplate(w, "result.html", Result{
						Name:     "Task",
						Protobuf: text,
						Back:     "/single/task",
					})
				}
			} else {
				err = templates.ExecuteTemplate(w, "single_task.html", nil)
			}
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	mux.HandleFunc("/single/result",
		func(w http.ResponseWriter, req *http.Request) {
			var err error
			if req.Method == "POST" {
				var text string
				text, err = decoder.SingleResultDecoder.DecodeBase64ToText(
					req.FormValue("base64data"))
				if err == nil {
					err = templates.ExecuteTemplate(w, "result.html", Result{
						Name:     "Result",
						Protobuf: text,
						Back:     "/single/result",
					})
				}
			} else {
				err = templates.ExecuteTemplate(w, "single_result.html", nil)
			}
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(*listenOn, n)
}
