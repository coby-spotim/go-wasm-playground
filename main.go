package main

import (
	"github.com/coby-spotim/wasm-playground/internal/dom"
	"github.com/coby-spotim/wasm-playground/internal/dom/eventListener"
)

func main() {
	c := make(chan struct{}, 1)
	window := dom.NewWindow()
	document := dom.NewDocument(window)
	window.Alert("Hello, World!")
	println(document.DocumentURI())
	document.AddEventListener("test", &eventListener.ConsoleLogger{}, eventListener.Options{}, false, false)
	<-c
}
