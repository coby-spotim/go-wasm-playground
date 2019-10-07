build:
	GOOS=js GOARCH=wasm go build -o output/main.wasm

tiny:
	tinygo build -o output-tiny/main.wasm -target wasm

.PHONY : serve
serve:
	servedir --dir ./output

.PHONY : tinyserve
tinyserve:
	servedir --dir ./output-tiny