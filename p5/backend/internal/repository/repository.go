package repository

import (
	"p5/internal/models"
	"p5/internal/repository/postgresql"
	"p5/internal/repository/postgresql/users"
)

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	GetPaginatedUsers(page int, pageSize int) (models.PaginatedResponse, error)
}

type Repositories struct {
	UserRepository
}

func NewRepositories(db *postgresql.Dialect) *Repositories {
	return &Repositories{
		UserRepository: users.NewUserRepository(db),
	}
}
