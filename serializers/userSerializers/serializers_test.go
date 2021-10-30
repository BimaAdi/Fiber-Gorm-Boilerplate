package userserializers_test

import (
	"encoding/json"
	"testing"

	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/models"
	userserializers "github.com/BimaAdi/fiberPostgresqlBoilerPlate/serializers/userSerializers"
	"github.com/stretchr/testify/assert"
)

func TestUserRequestSerializerToUserModel(t *testing.T) {
	input := userserializers.UserRequestSerializer{
		Username:        "test@local.com",
		Password:        "password",
		ConfirmPassword: "password",
		IsAdmin:         false,
	}

	output := userserializers.UserRequestSerializerToUserModel(input)

	// Username, Password and IsAdmin should return same value
	assert.Equal(t, input.Username, output.Username,
		"UserRequest Username value and models.User Username value should same")
	assert.Equal(t, input.Password, output.Password,
		"UserRequest Password value and models.User Password value should same")
	assert.Equal(t, input.IsAdmin, output.IsAdmin,
		"UserRequest IsAdmin value and models.User IsAdmin value should same")
}

func TestUserModelToUserResponseSerializer(t *testing.T) {
	input := models.User{
		ID:       1,
		Username: "test@local.com",
		Password: "password",
		IsAdmin:  false,
	}

	output := userserializers.UserModelToUserResponseSerializer(&input)

	// Id, Username and IsAdmin should return same value
	assert.Equal(t, int(input.ID), int(output.Id),
		"UserSerializer ID value and User ID value should same")
	assert.Equal(t, input.Username, output.Username,
		"UserSerializer Username value and User Username value should same")
	assert.Equal(t, input.IsAdmin, output.IsAdmin,
		"UserSerializer IsAdmin value and User IsAdmin value should same")

	// Check JSON key value format
	marshaled, _ := json.Marshal(output)
	assert.JSONEq(t, `{"id":1,"username":"test@local.com","is_admin":false}`, string(marshaled),
		`json key and value format should:
		{
			"id":1,
			"username":"test@local.com",
			"is_admin":false
		}`)
}

func TestManyUserModelToUserResponseSerializer(t *testing.T) {
	input := []models.User{
		{
			ID:       1,
			Username: "test1@local.com",
			Password: "password",
			IsAdmin:  true,
		},
		{
			ID:       2,
			Username: "test2@local.com",
			Password: "password",
			IsAdmin:  false,
		},
		{
			ID:       3,
			Username: "test3@local.com",
			Password: "password",
			IsAdmin:  false,
		},
	}
	inputPage := 1
	inputSize := 3
	inputNumData := 10

	output := userserializers.ManyUserModelToUserResponseSerializer(&input, inputPage, inputSize, &inputNumData)

	// Check value
	assert.Equal(t, inputPage, output.Page)
	assert.Equal(t, inputSize, output.Size)
	assert.Equal(t, 4, output.NumPage) // should be number of page
	// Check value inside output.Data
	for i := 0; i < 3; i++ {
		assert.Equal(t, int(input[i].ID), int(output.Data[i].Id))
		assert.Equal(t, input[i].Username, output.Data[i].Username)
		assert.Equal(t, input[i].IsAdmin, output.Data[i].IsAdmin)
	}

	// Check JSON key value format
	marshaled, _ := json.Marshal(output)

	assert.JSONEq(t, `{"page":1,"size":3,"num_page":4,"data":[{"id":1,"username":"test1@local.com","is_admin":true},{"id":2,"username":"test2@local.com","is_admin":false},{"id":3,"username":"test3@local.com","is_admin":false}]}`,
		string(marshaled),
		`JSON Format should:
		{
			"page":1,
			"size":3,
			"num_page":4,
			"data":[
				{
					"id":1,
					"username":"test1@local.com",
					"is_admin":true
				},{
					"id":2,
					"username":"test2@local.com",
					"is_admin":false
				},{
					"id":3,
					"username":"test3@local.com",
					"is_admin":false
				}
			]
		}`)
}
