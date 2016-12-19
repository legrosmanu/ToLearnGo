package main

import (
	"log"
	"net/http"
)

// "go run" must be launched on the webServer.go directory. It's to find the "gui" directory.
func main() {
	// The static files are on the gui directory.
	fs := http.FileServer(http.Dir("gui"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
