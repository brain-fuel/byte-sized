package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// Command line flags
	port := flag.Int("port", 6955, "port to serve on")
	directory := flag.String("dir", ".", "directory to serve")
	flag.Parse()

	// Get absolute path of the directory
	dir, err := filepath.Abs(*directory)
	if err != nil {
		log.Fatal(err)
	}

	// Create file server handler
	fileServer := http.FileServer(http.Dir(dir))

	// Create custom handler to add logging
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		fileServer.ServeHTTP(w, r)
	})

	// Register handler and start server
	http.Handle("/", handler)

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting file server on port %d serving directory: %s", *port, dir)
	log.Printf("Visit http://localhost%s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
