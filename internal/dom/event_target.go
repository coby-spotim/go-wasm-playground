package dom

import (
	"syscall/js"
)

type EventTarget struct {
	value js.Value
}

func (e *EventTarget) AddEventListener(eventType string, listener EventListener, options EventListenerOptions, useCapture bool, wantsUntrusted bool) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		listener.HandleEvent(args[0])
		return nil
	})
	e.value.Call("addEventListener", eventType, cb, options.ToJSValue(), useCapture, wantsUntrusted)
}
