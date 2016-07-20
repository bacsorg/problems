package main

import (
	"fmt"

	"github.com/bacsorg/problems/go/bacs/problem/decoder"
)

type brokerTaskRenderer struct{}
type brokerResultRenderer struct{}
type brokerStatusRenderer struct{}

var BrokerTaskRenderer brokerTaskRenderer
var BrokerResultRenderer brokerResultRenderer
var BrokerStatusRenderer brokerStatusRenderer

func (_ brokerTaskRenderer) Page() RendererPage {
	return RendererPage{
		Title: "Task",
		Types: map[string]string{
			"single": "Single",
		},
	}
}

func (_ brokerResultRenderer) Page() RendererPage {
	return RendererPage{
		Title: "Result",
		Types: map[string]string{
			"single": "Single",
		},
	}
}

func (_ brokerStatusRenderer) Page() RendererPage {
	return RendererPage{
		Title: "Status",
		Types: map[string]string{
			"single": "Single",
		},
	}
}

func (rend brokerTaskRenderer) Render(form RendererPageForm) RendererPage {
	page := rend.Page()
	page.Form = form
	switch form.Type {
	case "single":
		result, err := decoder.NewBrokerTaskDecoder(decoder.SingleTaskDecoder).
			DecodeBase64ToText(form.Base64Data)
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

func (rend brokerResultRenderer) Render(form RendererPageForm) RendererPage {
	page := rend.Page()
	page.Form = form
	switch form.Type {
	case "single":
		result, err := decoder.NewBrokerResultDecoder(decoder.SingleResultDecoder).
			DecodeBase64ToText(form.Base64Data)
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

func (rend brokerStatusRenderer) Render(form RendererPageForm) RendererPage {
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
