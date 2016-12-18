package main

import (
	"log"
	"net/http"
)

// "go run" must be launched on the webServer.go directory. It's to find the "gui" directory.
func main() {
	fs := http.FileServer(http.Dir("gui"))
	http.Handle("/", fs)

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
