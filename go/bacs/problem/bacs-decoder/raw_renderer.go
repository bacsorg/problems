package main

import (
	"fmt"

	"github.com/bacsorg/problems/go/bacs/problem/decoder"
)

type rawTaskRenderer struct{}
type rawResultRenderer struct{}

var RawTaskRenderer rawTaskRenderer
var RawResultRenderer rawResultRenderer

func (_ rawTaskRenderer) Page() RendererPage {
	return RendererPage{
		Title: "Task",
		Types: map[string]string{
			"single": "Single",
		},
	}
}

func (_ rawResultRenderer) Page() RendererPage {
	return RendererPage{
		Title: "Result",
		Types: map[string]string{
			"single": "Single",
		},
	}
}

func (rend rawTaskRenderer) Render(form RendererPageForm) RendererPage {
	page := rend.Page()
	page.Form = form
	switch form.Type {
	case "single":
		result, err := decoder.SingleTaskDecoder.DecodeBase64ToText(
			form.Base64Data)
		if err != nil {
			page.Error = err.Error()
		} else {
			page.RenderedData = result
		}
	default:
		page.Error = fmt.Sprintf("Unknown data type: %s", form.Type)
	}
	return page
}

func (rend rawResultRenderer) Render(form RendererPageForm) RendererPage {
	page := rend.Page()
	page.Form = form
	switch form.Type {
	case "single":
		result, err := decoder.SingleResultDecoder.DecodeBase64ToText(
			form.Base64Data)
		if err != nil {
			page.Error = err.Error()
		} else {
			page.RenderedData = result
		}
	default:
		page.Error = fmt.Sprintf("Unknown data type: %s", form.Type)
	}
	return page
}
