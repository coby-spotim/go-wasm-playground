package dom

import "syscall/js"

type window struct {
	value js.Value
}

func NewWindow() *window {
	return &window{
		value: js.Global(),
	}
}

// Alert creates a popup alert in the browser with the requested method
func (w *window) Alert(message string) {
	w.value.Call("alert", message)
}

// Document returns the document object on the current window
func (w *window) Document() js.Value {
	return w.value.Get("document")
}
