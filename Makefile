MAIN_PATH=./cmd

run:
	@go run cmd/main.go

server:
	@go run server/main.go

wasm:
	GOOS=js GOARCH=wasm go build -ldflags="-w -s" -o web/main.wasm $(MAIN_PATH)
	cp $$(go env GOROOT)/lib/wasm/wasm_exec.js web/
wasms: wasm
	go run -v github.com/hajimehoshi/wasmserve@latest -http=:8081 $(MAIN_PATH)
