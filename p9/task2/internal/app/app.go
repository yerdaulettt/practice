package app

import (
	"log"
	"net/http"
	"time"

	"p9task2/internal/middleware"
	"p9task2/internal/models"

	"github.com/redis/go-redis/v9"
)

func Run() {
	cache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	defer cache.Close()

	redisCache := models.NewRedisCache(cache, 24*time.Hour)

	r := http.NewServeMux()

	r.HandleFunc("POST /pay", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		time.Sleep(5 * time.Second)
		w.Write([]byte(`{"status": "paid", "amount": 1000, "transaction_id": "uuid-1231asd"}`))
	})

	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", middleware.Log(middleware.Idempotency(redisCache, r))))
}
