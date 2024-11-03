build:
	env GOOS=js GOARCH=wasm go build -o flappy-ebiten.wasm github.com/diki-haryadi/flappy-ebiten