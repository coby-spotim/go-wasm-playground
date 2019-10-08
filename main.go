package main

import (
	. "github.com/coby-spotim/wasm-playground/internal/dom"
)

func main() {
	window := Window()
	document := Document(window)
	window.Alert("Hello, World!")
	println(document.DocumentURI())
}
