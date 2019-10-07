package window

import "syscall/js"

// Alert creates a popup alert in the browser with the requested method
func Alert(message string) {
	js.Global().Call("alert", message)
}
