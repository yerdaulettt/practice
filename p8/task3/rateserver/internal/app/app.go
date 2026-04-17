package app

import (
	"log"
	"net/http"
	"p8t3/internal/handlers"
)

func Run() {
	r := http.NewServeMux()

	r.HandleFunc("GET /health", handlers.Healthcheck)
	r.HandleFunc("GET /convert", handlers.RateConvert)

	log.Println("starting")
	log.Fatal(http.ListenAndServe(":8080", r))
}
