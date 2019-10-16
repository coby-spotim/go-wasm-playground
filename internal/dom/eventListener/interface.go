package eventListener

import "syscall/js"

type EventListener interface {
	HandleEvent(js.Value)
}
