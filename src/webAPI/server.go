package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
