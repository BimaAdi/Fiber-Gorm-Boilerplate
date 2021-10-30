package models

import "github.com/stretchr/testify/mock"

type MockedUser struct {
	mock.Mock
}

func (m MockedUser) GetAllUser(page int, size int) (*[]User, *int, error) {
	args := m.Called(page, size)
	count := args.Int(1)
	return args.Get(0).(*[]User), &count, args.Error(2)
}

func (m MockedUser) GetDetailUser(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func (m MockedUser) CreateUser(user User) (*User, error) {
	args := m.Called(user)
	data := args.Get(0).(func(User) *User)
	return data(user), args.Error(1)
}

func (m MockedUser) UpdateUser(user User, id int) (*User, error) {
	args := m.Called(user, id)
	data := args.Get(0).(func(User, int) *User)
	return data(user, id), args.Error(1)
}

func (m MockedUser) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
