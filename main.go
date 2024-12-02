package main

import (
	_ "embed"
	"fmt"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func serve() {
	// Define the public directory path
	publicDir := "public"
	htmlPath := filepath.Join(publicDir, "index.html")

	// Log the paths for debugging
	log.Printf("Serving index.html from %s\n", htmlPath)
	log.Printf("Serving static files from %s\n", publicDir)

	// Serve index.html for the root route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	})

	http.HandleFunc("/tabler", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/tabler.html")
	})

	http.HandleFunc("/embed", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/embed.html")
	})

	// Serve all static files from the public directory directly
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))
	// Serve the chunked wasm file at /wasm/chunk/<index>
	http.HandleFunc("/wasm/chunk/", serveChunkedWasm)

	// Start the server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

const chunkSize = 64 * 1024 // 64 KB per chunk

// Serve the chunked wasm file
func serveChunkedWasm(w http.ResponseWriter, r *http.Request) {
	// Extract the chunk index from the URL
	// Example URL: /wasm/chunk/0, /wasm/chunk/1, etc.
	chunkIndex := r.URL.Path[len("/wasm/chunk/"):]
	index, err := strconv.Atoi(chunkIndex)
	if err != nil {
		http.Error(w, "Invalid chunk index", http.StatusBadRequest)
		return
	}

	// Read the wasm file
	wasmFilePath := "./public/flappy-ebiten.wasm" // Update this with your actual .wasm file path
	file, err := os.Open(wasmFilePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Get the total file size
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Could not get file info", http.StatusInternalServerError)
		return
	}

	fileSize := fileInfo.Size()

	// Calculate the starting and ending positions for this chunk
	start := int64(index) * chunkSize
	end := start + chunkSize
	if end > fileSize {
		end = fileSize
	}

	// Read the chunk
	chunk := make([]byte, end-start)
	_, err = file.ReadAt(chunk, start)
	if err != nil {
		http.Error(w, "Error reading file chunk", http.StatusInternalServerError)
		return
	}

	// Set the content type and serve the chunk
	w.Header().Set("Content-Type", "application/wasm")
	w.WriteHeader(http.StatusOK)
	w.Write(chunk)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [gui|serve]\n", os.Args[0])
		os.Exit(1)
	}

	if len(os.Args) > 1 && os.Args[1] == "serve" {
		serve()
	}

}
