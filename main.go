package main

import (
	"github.com/coby-spotim/wasm-playground/internal/dom"
)

func main() {
	c := make(chan bool, 1)
	window := dom.NewWindow()
	document := dom.NewDocument(window)
	window.Alert("Hello, World!")
	println(document.DocumentURI())
	document.AddEventListener("test", &dom.ConsoleLogEventListener{}, dom.EventListenerOptions{}, false, false)
	<-c
}
