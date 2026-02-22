package handlers

import (
	"context"
	"log"
	"os"
	"time"

	"p3/internal/repository"
	"p3/internal/repository/_postgres"
	"p3/pkg/modules"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func initConfig() *modules.PostgresqlConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	return &modules.PostgresqlConfig{
		Host:        host,
		Port:        port,
		Username:    dbUser,
		Password:    password,
		DBName:      dbname,
		SSLMode:     sslmode,
		ExecTimeout: 5 * time.Second,
	}
}

func dbStart() *repository.Repositories {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbConfig := initConfig()

	db := _postgres.NewPGXDialect(ctx, dbConfig)

	return repository.NewRepositories(db)
}

func initRedis() *redis.Client {
	redisCache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	log.Println("redis opened...")
	return redisCache
}
