package main

import (
	"fmt"
	"log"
	"net/http"

	"p2/internal/handlers"
	"p2/internal/middleware"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("GET /tasks", handlers.GetTasks)
	r.HandleFunc("POST /tasks", handlers.CreateTask)
	r.HandleFunc("PATCH /tasks", handlers.UpdateTask)
	r.HandleFunc("DELETE /tasks", handlers.DeleteTask)

	fmt.Println("Starting server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.AuthAndLogMiddleware(r)))
}
