package usecase

import (
	"p5/internal/models"
	"p5/internal/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r *repository.Repositories) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (u *UserUsecase) GetUserByID(id int) (*models.User, error) {
	user, err := u.repo.GetUserByID(id)
	return user, err
}

func (u *UserUsecase) GetPaginatedUsers(page int, pageSize int) (models.PaginatedResponse, error) {
	response, err := u.repo.GetPaginatedUsers(page, pageSize)
	return response, err
}
