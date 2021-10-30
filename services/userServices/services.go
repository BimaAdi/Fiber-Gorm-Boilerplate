package userservices

import (
	"errors"
	"os"
	"strconv"

	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/models"
	userserializers "github.com/BimaAdi/fiberPostgresqlBoilerPlate/serializers/userSerializers"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUserService(ui models.UserInterface, page int, size int) (*userserializers.UserPaginationResponseSerializer, error) {
	// default parameter
	if page == 0 {
		page = 1
	}

	if size == 0 {
		size = 5
	}

	users, count, err := ui.GetAllUser(page, size)
	if err != nil {
		return nil, err
	}

	// model to json
	responses := userserializers.ManyUserModelToUserResponseSerializer(users, page, size, count)
	return &responses, nil
}

func GetDetailUserService(ui models.UserInterface, id int) (*userserializers.UserResponseSerializer, error) {
	user, err := ui.GetDetailUser(id)
	if err != nil {
		return nil, err
	}

	userResponse := userserializers.UserModelToUserResponseSerializer(user)
	return &userResponse, nil
}

func CreateUserService(ui models.UserInterface, serializer userserializers.UserRequestSerializer) (*userserializers.UserResponseSerializer, error) {
	// hash the password using bcrypt
	bcryptCost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		return nil, err
	}
	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(serializer.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	serializer.Password = string(bytesPassword)

	// save new user
	newUser := userserializers.UserRequestSerializerToUserModel(serializer)
	createdUser, err := ui.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	// fmt.Println("createdDataType is: ", reflect.TypeOf(newUser))
	response := userserializers.UserModelToUserResponseSerializer(createdUser)
	return &response, nil
}

func UpdateUserService(ui models.UserInterface, id int, serializer userserializers.UserUpdateRequestSerializer) (*userserializers.UserResponseSerializer, error) {
	// check is User exist
	updatedUser, err := ui.GetDetailUser(id)
	if err != nil {
		return nil, err
	}

	// Check is Old Password correct
	if bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte(serializer.OldPassword)) != nil {
		return nil, errors.New("wrong password")
	}

	// Hash the password using bcrypt
	bcryptCost, err := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if err != nil {
		return nil, err
	}
	bytesPassword, err := bcrypt.GenerateFromPassword([]byte(serializer.NewPassword), bcryptCost)
	if err != nil {
		return nil, err
	}

	// Update the data
	updatedUser.Username = serializer.Username
	updatedUser.IsAdmin = serializer.IsAdmin
	updatedUser.Password = string(bytesPassword)
	updatedUser, err = ui.UpdateUser(*updatedUser, id)
	if err != nil {
		return nil, err
	}

	response := userserializers.UserModelToUserResponseSerializer(updatedUser)
	return &response, nil
}

func DeleteUserService(ui models.UserInterface, id int) error {
	// check is user exist?
	_, err := ui.GetDetailUser(id)
	if err != nil {
		return err
	}

	// if exist delete the user
	err = ui.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
