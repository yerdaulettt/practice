package usecase

import (
	"p5/internal/models"
	"p5/internal/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r *repository.Repositories) repository.UserRepository {
	return &UserUsecase{repo: r}
}

func (u *UserUsecase) GetPaginatedUsers(filters *models.UserFilter, page int, pageSize int) (models.PaginatedResponse, error) {
	response, err := u.repo.GetPaginatedUsers(filters, page, pageSize)
	return response, err
}

func (u *UserUsecase) GetCommonFriends(id1 int, id2 int) ([]models.User, error) {
	response, err := u.repo.GetCommonFriends(id1, id2)
	return response, err
}
