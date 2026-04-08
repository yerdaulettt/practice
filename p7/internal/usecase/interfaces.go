package usecase

import "p7/internal/entity"

type UserInterface interface {
	LoginUser(user *entity.LoginUserDTO) (string, error)
	RegisterUser(user *entity.User) (*entity.User, string, error)
	GetMe(id any) (*entity.User, error)
}
