package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/schema"
)

var rendererTemplates = template.Must(template.ParseFiles(
	"templates/footer.html",
	"templates/header.html",
	"templates/render.html",
))

type Renderer interface {
	Page() RendererPage
	Render(form RendererPageForm) RendererPage
}

type RendererPage struct {
	Title        string
	Types        map[string]string
	Error        string
	RenderedData string
	Form         RendererPageForm
}

type RendererPageForm struct {
	Type       string
	Base64Data string
}

var rendererDecoder = schema.NewDecoder()

func Render(rend Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		page := rend.Page()
		if r.Method == "POST" {
			err = r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = rendererDecoder.Decode(&page, r.PostForm)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			page = rend.Render(page.Form)
		}
		err = rendererTemplates.ExecuteTemplate(w, "render.html", page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
