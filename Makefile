build:
	GOOS=js GOARCH=wasm go build -o output/main.wasm

.PHONY : serve
serve:
	servedir --dir ./output