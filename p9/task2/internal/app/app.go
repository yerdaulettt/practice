package app

import (
	"log"
	"net/http"
	"time"

	"p9task2/internal/middleware"
	"p9task2/internal/models"
)

func Run() {
	inMemoryStore := models.NewInMemory()

	r := http.NewServeMux()

	r.HandleFunc("POST /pay", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		time.Sleep(4 * time.Second)
		w.Write([]byte(`{"status": "paid", "amount": 1000, "transaction_id": "uuid-1231asd"}`))
	})

	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", middleware.Log(middleware.Idempotency(inMemoryStore, r))))
}
