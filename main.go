package main

import (
	"github.com/coby-spotim/wasm-playground/internal/dom"
)

func main() {
	window := dom.NewWindow()
	document := dom.NewDocument(window)
	window.Alert("Hello, World!")
	println(document.DocumentURI())
}
