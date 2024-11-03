package main

import (
	_ "embed"
	"fmt"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

	http.HandleFunc("/embed", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/embed.html")
	})

	// Serve all static files from the public directory directly
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public"))))

	// Start the server
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
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
