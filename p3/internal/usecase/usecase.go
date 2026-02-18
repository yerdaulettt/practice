package usecase

import (
	"fmt"
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
		fmt.Println("ERROR in usecase.go", err)
		return nil
	}

	return users
}

func (u *UserUseCase) NewUser(newUser modules.User) int {
	id, err := u.repo.NewUser(newUser)
	if err != nil {
		return -1
	}

	return id
}

func (u *UserUseCase) DeleteUser(id int) *modules.User {
	deletedUser, err := u.repo.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return deletedUser
}

func (u *UserUseCase) GetUserByid(id int) *modules.User {
	userwithID, err := u.repo.GetUserByid(id)
	if err != nil {
		fmt.Println("ERROR", err)
		return nil
	}

	return userwithID
}

func (u *UserUseCase) UpdateUser(id int, userToUpdate modules.User) *modules.User {
	updatedUser, err := u.repo.UpdateUser(id, userToUpdate)
	if err != nil {
		fmt.Println("ERROR", err)
		return nil
	}

	return updatedUser
}
