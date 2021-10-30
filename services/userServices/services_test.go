package userservices_test

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/models"
	userserializers "github.com/BimaAdi/fiberPostgresqlBoilerPlate/serializers/userSerializers"
	userservices "github.com/BimaAdi/fiberPostgresqlBoilerPlate/services/userServices"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUserService(t *testing.T) {
	// test get all user
	testUsers := []models.User{
		{
			ID:       1,
			Username: "alpha@local.com",
			Password: "password",
			IsAdmin:  true,
		},
		{
			ID:       2,
			Username: "beta@local.com",
			Password: "password",
			IsAdmin:  false,
		},
		{
			ID:       3,
			Username: "charlie@local.com",
			Password: "password",
			IsAdmin:  false,
		},
	}
	inputPage := 1
	inputSize := 3
	mockFunction := new(models.MockedUser)
	mockFunction.On("GetAllUser", 1, 3).Return(&testUsers, 8, nil)
	output, err := userservices.GetAllUserService(mockFunction, inputPage, inputSize)

	assert.Nil(t, err,
		"Should have no error when sucess")
	assert.NotNilf(t, output,
		"Output should not nil when success")

	for i := 0; i < 3; i++ {
		assert.Equal(t, int(testUsers[i].ID), int(output.Data[i].Id))
		assert.Equal(t, testUsers[i].Username, output.Data[i].Username)
		assert.Equal(t, testUsers[i].IsAdmin, output.Data[i].IsAdmin)
	}

	// Check JSON key value format
	marshaled, _ := json.Marshal(output)
	assert.JSONEq(t, `{"page":1,"size":3,"num_page":3,"data":[{"id":1,"username":"alpha@local.com","is_admin":true},{"id":2,"username":"beta@local.com","is_admin":false},{"id":3,"username":"charlie@local.com","is_admin":false}]}`,
		string(marshaled),
		`JSON Format Should:\n`+`
		{
			"page":1,
			"size":3,
			"num_page":3,
			"data":[
				{
					"id":1,
					"username":"alpha@local.com",
					"is_admin":true
				},{
					"id":2,
					"username":"beta@local.com",
					"is_admin":false
				},{
					"id":3,
					"username":"charlie@local.com",
					"is_admin":false
				}
			]
		}`)
}

func TestGetDetailUserService(t *testing.T) {
	// test User found
	testuser := models.User{
		ID:       1,
		Username: "test@local.com",
		Password: "password",
		IsAdmin:  true,
	}
	mockFunction := new(models.MockedUser)
	mockFunction.On("GetDetailUser", 1).Return(&testuser, nil)
	output, err := userservices.GetDetailUserService(mockFunction, 1)
	assert.Nil(t, err,
		"Should have no error when user found")
	assert.Equal(t, int(testuser.ID), int(output.Id),
		"UserSerializer ID value and testuser ID value should same")
	assert.Equal(t, testuser.Username, output.Username,
		"UserSerializer Username value and testuser Username value should same")
	assert.Equal(t, testuser.IsAdmin, output.IsAdmin,
		"UserSerializer IsAdmin value and testuser IsAdmin value should same")

	// test User not found
	mockFunction.On("GetDetailUser", 999).Return(&models.User{}, errors.New("User not Found"))
	output, err = userservices.GetDetailUserService(mockFunction, 999)
	assert.NotNil(t, err,
		"Should return error when user not found")
	assert.Nil(t, output,
		"Should return user as nil when user not found")
}

func TestCreateUserService(t *testing.T) {
	if os.Getenv("IS_PROD") != "TRUE" {
		err := godotenv.Load("../../.env")
		if err != nil {
			panic(err)
		}
	}

	// 1. Test success create user
	input := userserializers.UserRequestSerializer{
		Username:        "test@local.com",
		Password:        "password",
		ConfirmPassword: "password",
		IsAdmin:         false,
	}
	mockFunction := new(models.MockedUser)
	mockFunction.On("CreateUser", mock.AnythingOfType("models.User")).Return(
		func(user models.User) *models.User {
			return &models.User{
				ID:       1,
				Username: user.Username,
				Password: user.Password,
				IsAdmin:  user.IsAdmin,
			}
		},
		nil)

	output, err := userservices.CreateUserService(mockFunction, input)
	assert.Nil(t, err,
		"Should have no error when user successfully created")
	assert.Equal(t, input.Username, output.Username,
		"UserRequestSerializer(input) Username value and UserResponseSerializer(output) Username value should same")
	assert.Equal(t, input.IsAdmin, output.IsAdmin,
		"UserRequestSerializer(input) IsAdmin value and UserResponseSerializer(output) IsAdmin value should same")

	// 2. Test error when create user
	input2 := userserializers.UserRequestSerializer{
		Username:        "test@local.com",
		Password:        "password",
		ConfirmPassword: "password",
		IsAdmin:         false,
	}
	mockFunction2 := new(models.MockedUser)
	mockFunction2.On("CreateUser", mock.AnythingOfType("models.User")).Return(
		func(user models.User) *models.User {
			return nil
		},
		errors.New("There is a Problem connect to database"))

	output2, err := userservices.CreateUserService(mockFunction2, input2)
	assert.NotNil(t, err,
		"Should return error when error created user")
	assert.Nil(t, output2,
		"Should return UserResponseSerializer as nil when error occurred")
}

func TestUpdateUserService(t *testing.T) {
	if os.Getenv("IS_PROD") != "TRUE" {
		err := godotenv.Load("../../.env")
		if err != nil {
			panic(err)
		}
	}
	// 1. Test success update user
	input := userserializers.UserUpdateRequestSerializer{
		Username:        "test@local.com",
		OldPassword:     "12qwaszx",
		NewPassword:     "12qwaszx",
		ConfirmPassword: "12qwaszx",
		IsAdmin:         false,
	}
	mockFunction := new(models.MockedUser)
	mockFunction.On("GetDetailUser", 1).Return(
		&models.User{
			ID:       1,
			Username: "test@local.com",
			Password: "$2a$10$3YZK9zY.jG8rFU5LnnlGme8AwbyVjYMwDKwS6NLxTB6lt8xUW2hPW", // from bcrypt hash for 12qwaszx
			IsAdmin:  false,
		},
		nil,
	)
	mockFunction.On("UpdateUser", mock.AnythingOfType("models.User"), 1).Return(
		func(user models.User, id int) *models.User {
			return &models.User{
				ID:       uint(id),
				Username: user.Username,
				Password: user.Password,
				IsAdmin:  user.IsAdmin,
			}
		},
		nil,
	)
	output, err := userservices.UpdateUserService(mockFunction, 1, input)
	assert.Nil(t, err)
	assert.Equal(t, input.Username, output.Username)
	assert.Equal(t, input.IsAdmin, output.IsAdmin)

	// 2. Test old password not match
	input2 := userserializers.UserUpdateRequestSerializer{
		Username:        "test@local.com",
		OldPassword:     "notmatchwithpaswordondatabase",
		NewPassword:     "12qwaszx",
		ConfirmPassword: "12qwaszx",
		IsAdmin:         false,
	}
	mockFunction2 := new(models.MockedUser)
	mockFunction2.On("GetDetailUser", 1).Return(
		&models.User{
			ID:       1,
			Username: "test@local.com",
			Password: "$2a$10$3YZK9zY.jG8rFU5LnnlGme8AwbyVjYMwDKwS6NLxTB6lt8xUW2hPW", // from bcrypt hash for 12qwaszx
			IsAdmin:  false,
		},
		nil,
	)
	mockFunction2.On("UpdateUser", mock.AnythingOfType("models.User"), 1).Return(
		func(user models.User, id int) *models.User {
			return &models.User{
				ID:       uint(id),
				Username: user.Username,
				Password: user.Password,
				IsAdmin:  user.IsAdmin,
			}
		},
		nil,
	)
	output2, err := userservices.UpdateUserService(mockFunction2, 1, input2)
	assert.Nil(t, output2)
	assert.NotNil(t, err)

	// 3. Test User Not found
	input3 := userserializers.UserUpdateRequestSerializer{
		Username:        "test@local.com",
		OldPassword:     "12qwaszx",
		NewPassword:     "12qwaszx",
		ConfirmPassword: "12qwaszx",
		IsAdmin:         false,
	}
	mockFunction3 := new(models.MockedUser)
	mockFunction3.On("GetDetailUser", 1).Return(
		&models.User{},
		errors.New("User not Found"),
	)
	mockFunction3.On("UpdateUser", mock.AnythingOfType("models.User"), 1).Return(
		func(user models.User, id int) *models.User {
			return &models.User{
				ID:       uint(id),
				Username: user.Username,
				Password: user.Password,
				IsAdmin:  user.IsAdmin,
			}
		},
		nil,
	)
	output3, err := userservices.UpdateUserService(mockFunction3, 1, input3)
	assert.Nil(t, output3)
	assert.NotNil(t, err)

	// 4. Error when update user
	input4 := userserializers.UserUpdateRequestSerializer{
		Username:        "test@local.com",
		OldPassword:     "12qwaszx",
		NewPassword:     "12qwaszx",
		ConfirmPassword: "12qwaszx",
		IsAdmin:         false,
	}
	mockFunction4 := new(models.MockedUser)
	mockFunction4.On("GetDetailUser", 1).Return(
		&models.User{
			ID:       1,
			Username: "test@local.com",
			Password: "$2a$10$3YZK9zY.jG8rFU5LnnlGme8AwbyVjYMwDKwS6NLxTB6lt8xUW2hPW", // from bcrypt hash for 12qwaszx
			IsAdmin:  false,
		},
		nil,
	)
	mockFunction4.On("UpdateUser", mock.AnythingOfType("models.User"), 1).Return(
		func(user models.User, id int) *models.User {
			return &models.User{}
		},
		errors.New("Error when update user"),
	)
	output4, err := userservices.UpdateUserService(mockFunction4, 1, input4)
	assert.Nil(t, output4)
	assert.NotNil(t, err)
}

func TestDeleteUserService(t *testing.T) {
	// 1. Test Success Delete User
	mockFunction := new(models.MockedUser)
	mockFunction.On("GetDetailUser", 1).Return(
		&models.User{
			ID:       1,
			Username: "test@local.com",
			Password: "$2a$10$3YZK9zY.jG8rFU5LnnlGme8AwbyVjYMwDKwS6NLxTB6lt8xUW2hPW", // from bcrypt hash for 12qwaszx
			IsAdmin:  false,
		},
		nil,
	)
	mockFunction.On("DeleteUser", 1).Return(nil)

	err := userservices.DeleteUserService(mockFunction, 1)
	assert.Nil(t, err,
		"Should return error as nil when delete user success")

	// 2. Test User not found
	mockFunction2 := new(models.MockedUser)
	mockFunction2.On("GetDetailUser", 2).Return(
		&models.User{},
		errors.New("User not Found"),
	)
	mockFunction2.On("DeleteUser", 2).Return(nil)

	err2 := userservices.DeleteUserService(mockFunction2, 2)
	assert.NotNil(t, err2,
		"Should return error when user not found")

	// 3. Test Error when delete user
	mockFunction3 := new(models.MockedUser)
	mockFunction3.On("GetDetailUser", 3).Return(
		&models.User{
			ID:       3,
			Username: "test@local.com",
			Password: "$2a$10$3YZK9zY.jG8rFU5LnnlGme8AwbyVjYMwDKwS6NLxTB6lt8xUW2hPW", // from bcrypt hash for 12qwaszx
			IsAdmin:  false,
		},
		nil,
	)
	mockFunction3.On("DeleteUser", 3).Return(errors.New("Something wrong with database"))

	err3 := userservices.DeleteUserService(mockFunction3, 3)
	assert.NotNil(t, err3,
		"Should return error when there is problem with database")
}
