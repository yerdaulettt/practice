package app

import (
	"log"
	"net/http"

	"p3/internal/handlers"
	"p3/internal/middleware"
)

func Run() {
	r := http.NewServeMux()

	r.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"working"}`))
	})

	r.HandleFunc("GET /users", handlers.GetAllUsers)
	r.HandleFunc("GET /users/{id}", handlers.GetUserByID)

	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", middleware.LogMiddleware(r)))
}
