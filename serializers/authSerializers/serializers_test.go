package authserializers_test

import (
	"encoding/json"
	"testing"

	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/models"
	authserializers "github.com/BimaAdi/fiberPostgresqlBoilerPlate/serializers/authSerializers"
	"github.com/stretchr/testify/assert"
)

func TestUserModelToLoginSuccessSerializer(t *testing.T) {
	input := models.User{
		ID:       1,
		Username: "test@local.com",
		Password: "password",
		IsAdmin:  true,
	}

	inputToken := []byte("qwertyuiop")

	output := authserializers.UserModelToLoginSuccessSerializer(input, inputToken)

	// Check value
	assert.Equal(t, int(input.ID), int(output.Id),
		"model.User ID value and LoginSuccesSerializer ID value should same")
	assert.Equal(t, input.Username, output.Username,
		"model.User Username value and LoginSuccesSerializer Username value should same")
	assert.Equal(t, input.IsAdmin, output.IsAdmin,
		"model.User IsAdmin value and LoginSuccesSerializer IsAdmin value should same")
	assert.Equal(t, string(inputToken), string(output.Token),
		"token value and LoginSuccesSerializer Token value should same")

	// Check JSON key value format
	marshaled, _ := json.Marshal(output)
	assert.JSONEq(t, `{"id":1,"username":"test@local.com","is_admin":true,"token":"qwertyuiop"}`, string(marshaled),
		`JSON format should:
		{
			"id":1,
			"username":"test@local.com",
			"is_admin":true,
			"token":"qwertyuiop"
		}`)
}

func TestUserModelToLogoutSuccessSerializer(t *testing.T) {
	input := models.User{
		ID:       1,
		Username: "test@local.com",
		Password: "password",
		IsAdmin:  true,
	}

	output := authserializers.UserModelToLogoutSuccessSerializer(input)

	// Check Value
	assert.Equal(t, int(input.ID), int(output.Id),
		"model.User ID value and LoginSuccesSerializer ID value should same")
	assert.Equal(t, input.Username, output.Username,
		"model.User Username value and LoginSuccesSerializer Username value should same")
	assert.Equal(t, input.IsAdmin, output.IsAdmin,
		"model.User IsAdmin value and LoginSuccesSerializer IsAdmin value should same")

	// Check JSON key value format
	marshaled, _ := json.Marshal(output)
	assert.JSONEq(t, `{"id":1,"username":"test@local.com","is_admin":true}`, string(marshaled),
		`JSON format should:
		{
			"id":1,
			"username":"test@local.com",
			"is_admin":true
		}`)
}
