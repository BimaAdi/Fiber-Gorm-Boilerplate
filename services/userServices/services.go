package userservices

import (
	"fmt"
	"reflect"

	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/models"
	userserializers "github.com/BimaAdi/fiberPostgresqlBoilerPlate/serializers/userSerializers"
)

func GetAllUserService(page int, size int) (*userserializers.UserPaginationResponseSerializer, error) {
	// default parameter
	if page == 0 {
		page = 1
	}

	if size == 0 {
		size = 5
	}

	// get all users from database
	limit := size
	offset := (page - 1) * size
	users := []models.User{}
	if err := models.DBConn.Order("id DESC").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}

	// count all users in database
	var count int64
	if err := models.DBConn.Model(&users).Count(&count).Error; err != nil {
		return nil, err
	}

	// model to json
	responses := userserializers.ManyUserModelToUserResponseSerializer(users, page, size, int(count))
	return &responses, nil
}

func GetDetailUserService(id int) (*userserializers.UserResponseSerializer, error) {
	user := models.User{}
	if err := models.DBConn.First(&user, id).Error; err != nil {
		return nil, err
	}

	userResponse := userserializers.UserModelToUserResponseSerializer(user)
	return &userResponse, nil
}

func CreateUserService(serializer userserializers.UserRequestSerializer) (*userserializers.UserResponseSerializer, error) {
	newUser := userserializers.UserRequestSerializerToUserModel(serializer)
	if err := models.DBConn.Create(&newUser).Error; err != nil {
		return nil, err
	}

	fmt.Println("createdDataType is: ", reflect.TypeOf(newUser))
	response := userserializers.UserModelToUserResponseSerializer(newUser)
	return &response, nil
}

func UpdateUserService(id int, serializer userserializers.UserRequestSerializer) (*userserializers.UserResponseSerializer, error) {
	// check is User exist
	updatedUser := models.User{}
	if err := models.DBConn.First(&updatedUser, id).Error; err != nil {
		return nil, err
	}

	// Update the data
	updatedUser.Username = serializer.Username
	updatedUser.IsAdmin = serializer.IsAdmin
	updatedUser.Password = serializer.Password
	if err := models.DBConn.Save(updatedUser).Error; err != nil {
		return nil, err
	}

	response := userserializers.UserModelToUserResponseSerializer(updatedUser)
	return &response, nil
}

func DeleteUserService(id int) error {
	user := models.User{}
	// check is user exist?
	if err := models.DBConn.First(&user, id).Error; err != nil {
		return err
	}

	// if exist delete the user
	if err := models.DBConn.Unscoped().Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}
