package authservices_test

import (
	"errors"
	"os"
	"testing"

	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/models"
	authserializers "github.com/BimaAdi/fiberPostgresqlBoilerPlate/serializers/authSerializers"
	authservices "github.com/BimaAdi/fiberPostgresqlBoilerPlate/services/authServices"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestLoginService(t *testing.T) {
	// load environtement variable
	if os.Getenv("IS_PROD") != "TRUE" {
		err := godotenv.Load("../../.env")
		if err != nil {
			panic(err)
		}
	}

	// 1. Login with correct credential
	mockFunction := new(models.MockedUser)
	mockFunction.On("GetDetailUserByUsername", "test@local.com").Return(
		&models.User{
			ID:       1,
			Username: "test@local.com",
			Password: "$2a$10$3YZK9zY.jG8rFU5LnnlGme8AwbyVjYMwDKwS6NLxTB6lt8xUW2hPW", // from bcrypt hash for 12qwaszx
			IsAdmin:  false,
		},
		nil,
	)
	output, err := authservices.LoginService(mockFunction, "test@local.com", "12qwaszx")
	expectedOutput := authserializers.LoginSuccessSerializer{
		Id:       1,
		Username: "test@local.com",
		IsAdmin:  false,
		Token:    "",
	}

	assert.Nil(t, err)
	assert.Equal(t, expectedOutput.Id, output.Id)
	assert.Equal(t, expectedOutput.Username, output.Username)
	assert.Equal(t, expectedOutput.IsAdmin, output.IsAdmin)
	assert.NotNil(t, output.Token)

	// 2. Login with wrong credential
	mockFunction2 := new(models.MockedUser)
	mockFunction2.On("GetDetailUserByUsername", "test@local.com").Return(
		&models.User{
			ID:       1,
			Username: "test@local.com",
			Password: "$2a$10$3YZK9zY.jG8rFU5LnnlGme8AwbyVjYMwDKwS6NLxTB6lt8xUW2hPW", // from bcrypt hash for 12qwaszx
			IsAdmin:  false,
		},
		nil,
	)
	output2, err := authservices.LoginService(mockFunction2, "test@local.com", "wrong password")
	assert.Nil(t, output2)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "wrong password", err.Error())
	}

	// 3. Login with unknown user
	mockFunction3 := new(models.MockedUser)
	mockFunction3.On("GetDetailUserByUsername", "notindatabase@local.com").Return(
		&models.User{},
		errors.New("User not Found"),
	)
	output3, err := authservices.LoginService(mockFunction3, "notindatabase@local.com", "wrong password")
	assert.Nil(t, output3)
	assert.NotNil(t, err)
}
