package repository

import (
	"p5/internal/models"
	"p5/internal/repository/postgresql"
	"p5/internal/repository/postgresql/users"
)

type UserRepository interface {
	GetPaginatedUsers(filters *models.UserFilter, page int, pageSize int) (models.PaginatedResponse, error)
	GetCommonFriends(id1 int, id2 int) ([]models.User, error)
}

type Repositories struct {
	UserRepository
}

func NewRepositories(db *postgresql.Dialect) *Repositories {
	return &Repositories{
		UserRepository: users.NewUserRepository(db),
	}
}
