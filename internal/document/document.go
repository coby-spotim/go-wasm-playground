package document

import "syscall/js"

type document struct {
	js.Value
}

// Document returns a struct representing the Document interface in the Browser
func Document() *document {
	return &document{js.Global().Get("document")}
}
