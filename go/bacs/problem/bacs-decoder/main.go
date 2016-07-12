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

func handleTemplate(w http.ResponseWriter, req *http.Request, name string) {
	err := templates.ExecuteTemplate(w, name, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func decodeHandler(w http.ResponseWriter, req *http.Request, name string,
	decode func(base64data string) (string, error)) {
	var err error
	if req.Method == "POST" {
		var text string
		text, err = decode(req.FormValue("base64data"))
		if err == nil {
			err = templates.ExecuteTemplate(w, "result.html", Result{
				Name:     name,
				Protobuf: text,
				Back:     req.URL.Path,
			})
		}
	} else {
		err = templates.ExecuteTemplate(w, "single_result.html", nil)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		handleTemplate(w, req, "index.html")
	})
	mux.HandleFunc("/single", func(w http.ResponseWriter, req *http.Request) {
		handleTemplate(w, req, "single_index.html")
	})
	mux.HandleFunc("/single/task",
		func(w http.ResponseWriter, req *http.Request) {
			decodeHandler(w, req, "Task",
				decoder.SingleTaskDecoder.DecodeBase64ToText)
		})
	mux.HandleFunc("/single/result",
		func(w http.ResponseWriter, req *http.Request) {
			decodeHandler(w, req, "Result",
				decoder.SingleResultDecoder.DecodeBase64ToText)
		})

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(*listenOn, n)
}
