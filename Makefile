build:
	env GOOS=js GOARCH=wasm go build -o flappy.wasm github.com/diki-haryadi/flappy-ebiten