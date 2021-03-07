package dom

import (
	"syscall/js"

	"github.com/probably-not/wasm-playground/internal/dom/eventListener"
)

type EventTarget struct {
	value js.Value
}

func (e *EventTarget) AddEventListener(eventType string, listener eventListener.EventListener, options eventListener.Options, useCapture bool, wantsUntrusted bool) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		listener.HandleEvent(args[0])
		return nil
	})
	e.value.Call("addEventListener", eventType, cb, options.ToJSValue(), useCapture, wantsUntrusted)
}
