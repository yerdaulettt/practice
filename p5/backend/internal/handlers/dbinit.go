package handlers

import (
	"os"

	"p5/internal/models"
	"p5/internal/repository"
	"p5/internal/repository/postgresql"
)

func initDB() *repository.Repositories {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSL")

	dbconfig := &models.PostgresConfiguration{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   dbname,
		SSLMode:  ssl,
	}

	db := postgresql.NewDialect(dbconfig)
	repos := repository.NewRepositories(db)

	return repos
}
