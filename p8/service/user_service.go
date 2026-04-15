package service

import "p8/repository"

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) GetUserByID(id int) (*repository.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) CreateUser(user *repository.User) error {
	return s.repo.CreateUser(user)
}
