package users

import (
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

	err := r.db.DB.Select(&users, "select * from users")
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) NewUser(newUser modules.User) (int, error) {
	var id int

	err := r.db.DB.QueryRow(
		"insert into users (name, age, hobby, profession) values ($1, $2, $3, $4) returning id",
		newUser.Name, newUser.Age, newUser.Hobby, newUser.Profession).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r *Repository) DeleteUser(id int) (*modules.User, error) {
	var deletedUser modules.User

	err := r.db.DB.QueryRow("delete from users where id = $1 returning *", id).Scan(
		&deletedUser.Id, &deletedUser.Name, &deletedUser.Age, &deletedUser.Hobby, &deletedUser.Profession)

	if err != nil {
		return nil, err
	}

	return &deletedUser, nil
}

func (r *Repository) GetUserByid(id int) (*modules.User, error) {
	var userWithId modules.User

	err := r.db.DB.QueryRow("select * from users where id = $1", id).Scan(
		&userWithId.Id, &userWithId.Name, &userWithId.Age, &userWithId.Hobby, &userWithId.Profession)

	if err != nil {
		return nil, err
	}

	return &userWithId, nil
}

func (r *Repository) UpdateUser(id int, userToUpdate modules.User) (*modules.User, error) {
	var updatedUser modules.User

	err := r.db.DB.QueryRow(
		"update users set (name, age, hobby, profession) = ($1, $2, $3, $4) where id = $5 returning *",
		userToUpdate.Name, userToUpdate.Age, userToUpdate.Hobby, userToUpdate.Profession, id).Scan(
		&updatedUser.Id, &updatedUser.Name, &updatedUser.Age, &updatedUser.Hobby, &updatedUser.Profession)

	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}
