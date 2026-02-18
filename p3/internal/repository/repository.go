package repository

import (
	"p3/internal/repository/_postgres"
	"p3/internal/repository/_postgres/users"
	"p3/pkg/modules"
)

type UserRepository interface {
	GetUsers() ([]modules.User, error)
	NewUser(newUser modules.User) (int, error)
	DeleteUser(id int) (*modules.User, error)
	GetUserByid(id int) (*modules.User, error)
	UpdateUser(id int, userToUpdate modules.User) (*modules.User, error)
}

type Repositories struct {
	UserRepository
}

func NewRepositories(db *_postgres.Dialect) *Repositories {
	return &Repositories{
		UserRepository: users.NewUserRepository(db),
	}
}
