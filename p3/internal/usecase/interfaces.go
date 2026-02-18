package usecase

import "p3/pkg/modules"

type UserUseCaseInterface interface {
	GetUsers() []modules.User
	NewUser(newUser modules.User) int
	DeleteUser(id int) *modules.User
	GetUserByid(id int) *modules.User
	UpdateUser(id int, usetToUpdate modules.User) *modules.User
}
