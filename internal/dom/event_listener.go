package dom

import (
	"syscall/js"
)

type EventListener interface {
	HandleEvent(js.Value)
}

type EventListenerOptions struct {
	Capture        bool
	Once           bool
	Passive        bool
	MozSystemGroup bool
}

func (opts *EventListenerOptions) ToJSValue() js.Value {
	return js.ValueOf(map[string]interface{}{
		"capture":        opts.Capture,
		"once":           opts.Once,
		"passive":        opts.Passive,
		"mozSystemGroup": opts.MozSystemGroup,
	})
}

type ConsoleLogEventListener struct {
}

func (clel *ConsoleLogEventListener) HandleEvent(event js.Value) {
	println("Handling Event")
}
