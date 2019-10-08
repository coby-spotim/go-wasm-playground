package main

import (
	. "github.com/coby-spotim/wasm-playground/internal/document"
	. "github.com/coby-spotim/wasm-playground/internal/window"
)

func main() {
	Window().Alert("Hello, World!")
	println(Document().Get("location").Get("href").String())
}
