package usecase

import "p4/pkg/modules"

type UserUseCaseInterface interface {
	GetUsers() ([]modules.User, error)
	NewUser(newUser modules.User) (int, error)
	DeleteUser(id int) (*modules.User, error)
	GetUserByid(id int) (*modules.User, error)
	UpdateUser(id int, usetToUpdate modules.User) (*modules.User, error)
}
