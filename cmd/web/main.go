package main

import (
	"log"
	"net/http"

	"github.com/muhammetburakgolec/burakgolec.com-go/pkg/handlers"
)

const port = ":5001"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/faqs", handlers.Faqs)
	http.HandleFunc("/features", handlers.Features)

	log.Println("Starting Serve On : http://localhost" + port)
	_ = http.ListenAndServe(port, nil)
}
