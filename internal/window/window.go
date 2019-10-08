package window

import "syscall/js"

type window struct {
	js.Value
}

func Window() *window {
	return &window{js.Global()}
}

// Alert creates a popup alert in the browser with the requested method
func (w *window) Alert(message string) {
	w.Call("alert", message)
}

// Document returns the document object on the current window
func (w *window) Document() js.Value {
	return w.Get("document")
}
