package main

import (
	"github.com/matthewmueller/golly/dom/document"
	"github.com/matthewmueller/golly/jsx"
	"github.com/matthewmueller/golly/testdata/49-jsx/header"
	"github.com/matthewmueller/golly/testdata/49-jsx/preact"
)

func main() {
	jsx.Use("preact.h", "./preact/preact.js")

	hdr := jsx.H("h2", map[string]interface{}{"class": "hi"},
		&jsx.Text{Value: "yo!"},
		header.New("lol", &jsx.Text{Value: "hi!"}),
	)

	preact.Render(hdr, document.Body)
	println(document.Body.InnerHTML())
}
