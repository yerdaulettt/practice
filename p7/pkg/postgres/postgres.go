package postgres

import (
	"log"
	"p7/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Conn *gorm.DB
}

func NewPostgresConn(url string) *Postgres {
	db, err := gorm.Open(postgres.Open(url))
	if err != nil {
		return nil
	}

	pg := &Postgres{Conn: db}

	err = pg.Conn.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println(err)
		return nil
	}

	return pg
}
