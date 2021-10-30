package models

type UserInterface interface {
	GetAllUser(page int, size int) (*[]User, *int, error)
	GetDetailUser(id int) (*User, error)
	CreateUser(user User) (*User, error)
	UpdateUser(user User, id int) (*User, error)
	DeleteUser(id int) error
}
