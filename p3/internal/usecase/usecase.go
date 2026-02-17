package usecase

import (
	"p3/internal/repository"
	"p3/pkg/modules"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(r *repository.Repositories) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (u *UserUseCase) GetUsers() []modules.User {
	users, err := u.repo.GetUsers()
	if err != nil {
		return nil
	}

	return users
}

func (u *UserUseCase) GetUserbyid(id int) modules.User {
	var user modules.User

	user, err := u.repo.GetUserByid(id)
	if err != nil {
		return user
	}

	return user
}
