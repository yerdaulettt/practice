package handlers

import (
	"context"
	"log"
	"os"
	"p3/internal/repository"
	"p3/internal/repository/_postgres"
	"p3/pkg/modules"
	"time"

	"github.com/joho/godotenv"
)

func initConfig() *modules.PostgresqlConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")

	return &modules.PostgresqlConfig{
		Host:        "localhost",
		Port:        "5432",
		Username:    dbUser,
		Password:    password,
		DBName:      dbname,
		SSLMode:     "disable",
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
