
# Wasm Awesome

WebAssembly (Wasm) is a powerful tool for running high-performance applications on the web. It allows you to run code written in go language.

## Screenshots

![Flappy Run](/screenshots/flappy.png)


## Run Locally

Clone the project

```bash
  git clone https://github.com/diki-haryadi/wasm
```

Go to the project directory

```bash
  cd wasm
```

Build Wasm file 
```bash
    make build
```

Install dependencies

```bash
  go mod tidy
```

Start the server

```bash
  go run main.go serve
```


## License

[MIT](https://choosealicense.com/licenses/mit/)

