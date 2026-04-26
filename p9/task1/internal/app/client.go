package app

import (
	"log"
	"net/http"

	"p9task1/internal/handlers"
)

func Run(testURL string) {
	r := http.NewServeMux()

	handlers.SetURL(testURL)

	r.HandleFunc("POST /pay", handlers.PaymentHandler)

	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
