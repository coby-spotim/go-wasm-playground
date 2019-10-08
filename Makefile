normal-build:
	GOOS=js GOARCH=wasm go build -o output/main.wasm

.PHONY : normal-serve
normal-serve:
	servedir --dir ./output

.PHONY : normal
normal: normal-build normal-serve

tiny-build:
	tinygo build -o output-tiny/main.wasm -target wasm --no-debug

.PHONY : tiny-serve
tiny-serve:
	servedir --dir ./output-tiny

.PHONY : tiny
tiny: tiny-build tiny-serve