package usecase

import "p3/pkg/modules"

type UserUseCaseInterface interface {
	GetUsers() []modules.User
	GetUserbyid(id int) modules.User
}
