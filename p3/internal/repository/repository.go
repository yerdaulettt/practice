package repository

import (
	"p3/internal/repository/_postgres"
	"p3/internal/repository/_postgres/users"
	"p3/pkg/modules"
)

type UserRepository interface {
	GetUsers() ([]modules.User, error)
	GetUserByid(id int) (modules.User, error)
}

type Repositories struct {
	UserRepository
}

func NewRepositories(db *_postgres.Dialect) *Repositories {
	return &Repositories{
		UserRepository: users.NewUserRepository(db),
	}
}
