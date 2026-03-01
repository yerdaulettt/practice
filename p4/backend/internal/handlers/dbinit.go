package handlers

import (
	"context"
	"log"
	"os"
	"time"

	"p4/internal/repository"
	"p4/internal/repository/_postgres"
	"p4/pkg/modules"

	"github.com/redis/go-redis/v9"
)

func initConfig() *modules.PostgresqlConfig {
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
	address := os.Getenv("CACHE_HOST") + ":" + os.Getenv("CACHE_PORT")
	redisCache := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	log.Println("redis opened...")
	return redisCache
}
