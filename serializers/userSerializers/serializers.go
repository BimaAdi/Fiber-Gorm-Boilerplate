package userserializers

import (
	"math"

	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/models"
	"github.com/go-playground/validator/v10"
)

type UserQueryParams struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

// combine fiber body parser (https://docs.gofiber.io/api/ctx#bodyparser)
// with validator (https://github.com/go-playground/validator)
// into one struct tags
type UserRequestSerializer struct {
	Username        string `json:"username" validate:"required,email,min=6,max=32"`
	Password        string `json:"password" validate:"required,min=6,max=32"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	IsAdmin         bool   `json:"is_admin" validate:"eq=True|eq=False"`
}

type UserUpdateRequestSerializer struct {
	Username        string `json:"username" validate:"required,email,min=6,max=32"`
	OldPassword     string `json:"old_password" validate:"required,min=6,max=32"`
	NewPassword     string `json:"new_password" validate:"required,min=6,max=32"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=NewPassword"`
	IsAdmin         bool   `json:"is_admin" validate:"eq=True|eq=False"`
}

type UserResponseSerializer struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

type UserPaginationResponseSerializer struct {
	Page    int                      `json:"page"`
	Size    int                      `json:"size"`
	NumPage int                      `json:"num_page"`
	Data    []UserResponseSerializer `json:"data"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateUser(user UserRequestSerializer) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateupdateUser(user UserUpdateRequestSerializer) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func UserRequestSerializerToUserModel(serializer UserRequestSerializer) models.User {
	return models.User{
		Username: serializer.Username,
		Password: serializer.Password,
		IsAdmin:  serializer.IsAdmin,
	}
}

func UserModelToUserResponseSerializer(model *models.User) UserResponseSerializer {
	return UserResponseSerializer{
		Id:       int(model.ID),
		Username: model.Username,
		IsAdmin:  model.IsAdmin,
	}
}

func ManyUserModelToUserResponseSerializer(manyModel *[]models.User, page int, size int, numData *int) UserPaginationResponseSerializer {
	userSerializers := []UserResponseSerializer{}
	for _, item := range *manyModel {
		userSerializers = append(userSerializers, UserResponseSerializer{
			Id:       int(item.ID),
			Username: item.Username,
			IsAdmin:  item.IsAdmin,
		})
	}

	response := UserPaginationResponseSerializer{
		Page:    page,
		Size:    size,
		NumPage: int(math.Ceil((float64(*numData) / float64(size)))),
		Data:    userSerializers,
	}

	return response
}
