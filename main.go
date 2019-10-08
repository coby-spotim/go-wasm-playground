package main

import (
	. "github.com/coby-spotim/wasm-playground/internal/document"
	. "github.com/coby-spotim/wasm-playground/internal/window"
)

func main() {
	window := Window()
	document := Document()
	window.Alert("Hello, World!")
	println(document.DocumentURI())
}
