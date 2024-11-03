build:
	env GOOS=js GOARCH=wasm go build -o ./public/flappy-ebiten.wasm github.com/diki-haryadi/wasm/flappy
	#cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./public
serve:
	go run github.com/hajimehoshi/wasmserve@latest .