package repo

import (
	"fmt"
	"log"
	"p7/internal/entity"
	"p7/pkg/postgres"
)

type UserRepo struct {
	PG *postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{PG: pg}
}

func (u *UserRepo) LoginUser(user *entity.LoginUserDTO) (*entity.User, error) {
	var userFromDB entity.User
	if err := u.PG.Conn.Where("username = ?", user.Username).First(&userFromDB).Error; err != nil {
		return nil, fmt.Errorf("Username not found: %v", err)
	}
	return &userFromDB, nil
}

func (u *UserRepo) RegisterUser(user *entity.User) (*entity.User, error) {
	err := u.PG.Conn.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) GetMe(id any) (*entity.User, error) {
	var me entity.User
	if err := u.PG.Conn.Where("id = ?", id).First(&me).Error; err != nil {
		return nil, err
	}
	return &me, nil
}

func (u *UserRepo) Promote(username string) (bool, error) {
	var id int
	if err := u.PG.Conn.Raw("update users set role = 'admin' where username = ? returning id", username).Row().Scan(&id); err != nil {
		log.Println(err)
		return false, err
	}

	log.Println(id)
	return true, nil
}
