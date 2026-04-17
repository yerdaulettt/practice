package repository

type User struct {
	ID   int
	Name string
}

type UserRepository interface {
	GetUserByID(id int) (*User, error)
	CreateUser(user *User) error
	GetByEmai(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int) error
}
