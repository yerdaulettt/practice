package usecase

import (
	"fmt"

	"p7/internal/entity"
	"p7/internal/usecase/repo"
	"p7/utils"

	"github.com/google/uuid"
)

type UserUseCase struct {
	repo *repo.UserRepo
}

func NewUserUseCase(r *repo.UserRepo) *UserUseCase {
	return &UserUseCase{repo: r}
}

func (u *UserUseCase) LoginUser(user *entity.LoginUserDTO) (string, error) {
	userFromRepo, err := u.repo.LoginUser(user)
	if err != nil {
		return "", fmt.Errorf("User from repo: %w", err)
	}
	if !utils.CheckPassword(userFromRepo.Password, user.Password) {
		return "", fmt.Errorf("Check password: %w", err)
	}
	token, err := utils.GenerateJWT(userFromRepo.ID, userFromRepo.Role)
	if err != nil {
		return "", fmt.Errorf("Generate JWT: %w", err)
	}
	return token, nil
}

func (u *UserUseCase) RegisterUser(user *entity.User) (*entity.User, string, error) {
	user, err := u.repo.RegisterUser(user)
	if err != nil {
		return nil, "", fmt.Errorf("register user: %w", err)
	}

	sessionId := uuid.New().String()
	return user, sessionId, nil
}
