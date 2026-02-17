package users

import (
	"errors"
	"time"

	"p3/internal/repository/_postgres"
	"p3/pkg/modules"
)

type Repository struct {
	db               *_postgres.Dialect
	executionTimeout time.Duration
}

func NewUserRepository(db *_postgres.Dialect) *Repository {
	return &Repository{
		db:               db,
		executionTimeout: 5 * time.Second,
	}
}

func (r *Repository) GetUsers() ([]modules.User, error) {
	var users []modules.User
	err := r.db.DB.Select(&users, "select id, name from users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) GetUserByid(id int) (modules.User, error) {
	var user modules.User
	err := r.db.DB.QueryRow("select * from users where id = 1", &user.Id, &user.Name, &user.Age, &user.Hobby, &user.Profession)

	if err != nil {
		return user, errors.New("sql error")
	}

	return user, nil
}
