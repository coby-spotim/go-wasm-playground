package eventListener

import "syscall/js"

type ConsoleLogger struct {
}

func (cl *ConsoleLogger) HandleEvent(event js.Value) {
	println("Handling Event")
}
