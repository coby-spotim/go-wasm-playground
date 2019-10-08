package document

import (
	"syscall/js"

	. "github.com/coby-spotim/wasm-playground/internal/window"
)

type document struct {
	js.Value
}

func Document() *document {
	return &document{Window().Document()}
}
