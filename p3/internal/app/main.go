package app

import (
	"log"
	"net/http"

	"p3/internal/handlers"
	"p3/internal/middleware"
)

func Run() {
	r := http.NewServeMux()

	r.HandleFunc("GET /healthcheck", handlers.Healthcheck)

	r.HandleFunc("GET /users", handlers.GetAllUsers)
	r.HandleFunc("GET /users/{id}", handlers.GetUserByID)
	r.HandleFunc("POST /users", handlers.NewUser)
	r.HandleFunc("DELETE /users/{id}", handlers.DeleteUser)
	r.HandleFunc("PATCH /users/{id}", handlers.UpdateUser)

	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", middleware.LogMiddleware(r)))
}
