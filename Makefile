# export GOROOT=$(shell go env GOROOT)

build:
	tinygo build -o direct-response.wasm -scheduler=none -target=wasi ./main.go
