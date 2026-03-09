package postgresql

import (
	"database/sql"
	"fmt"

	"p5/internal/models"

	_ "github.com/lib/pq"
)

type Dialect struct {
	DB *sql.DB
}

func NewDialect(cfg *models.PostgresConfiguration) *Dialect {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &Dialect{DB: db}
}
