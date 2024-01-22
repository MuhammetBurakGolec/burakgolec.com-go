package main

import (
	"log"
	"net/http"
)

const port = ":5001"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/faqs", Faqs)
	http.HandleFunc("/features", Features)
	http.HandleFunc("/login", Login)

	log.Println("Starting Serve On : http://localhost" + port)
	_ = http.ListenAndServe(port, nil)
}
