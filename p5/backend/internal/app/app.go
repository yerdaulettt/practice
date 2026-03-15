package app

import (
	"log"
	"net/http"

	"p5/internal/handlers"
	"p5/internal/middleware"
)

func Run() {
	r := http.NewServeMux()

	r.HandleFunc("GET /users", handlers.GetPaginatedUsers)
	r.HandleFunc("GET /common-friends", handlers.GetCommonFriends)

	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", middleware.Log(r)))
}
