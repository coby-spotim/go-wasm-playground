package eventListener

import "syscall/js"

type Options struct {
	Capture        bool
	Once           bool
	Passive        bool
	MozSystemGroup bool
}

func (opts *Options) ToJSValue() js.Value {
	return js.ValueOf(map[string]interface{}{
		"capture":        opts.Capture,
		"once":           opts.Once,
		"passive":        opts.Passive,
		"mozSystemGroup": opts.MozSystemGroup,
	})
}
