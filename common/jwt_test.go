package common_test

import (
	"errors"
	"os"
	"testing"

	"github.com/BimaAdi/fiberGormBoilerPlate/common"
	"github.com/BimaAdi/fiberGormBoilerPlate/models"
	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWTToken(t *testing.T) {
	if os.Getenv("IS_PROD") != "TRUE" {
		err := godotenv.Load("../.env")
		if err != nil {
			panic(err)
		}
	}

	// Test JWT Payload
	input := models.User{
		ID:       1,
		Username: "test@local.com",
		Password: "password",
		IsAdmin:  false,
	}

	output, err := common.GenerateJWTToken(input)
	assert.Nil(t, err,
		"err should be nil if success generate token")
	jwtSecret := os.Getenv("JWT_SECRET")
	resultToken, _ := jwt.Parse([]byte(output), jwt.WithVerify(jwa.HS512, []byte(jwtSecret)))
	tokenId, _ := resultToken.Get("id")
	tokenUsername, _ := resultToken.Get("username")
	assert.Equal(t, int(input.ID), int(tokenId.(float64)),
		"JWT token must have correct ID in Payload")
	assert.Equal(t, input.Username, tokenUsername,
		"JWT token must have correct username in Payload")
}

func TestValidateJWTToken(t *testing.T) {
	if os.Getenv("IS_PROD") != "TRUE" {
		err := godotenv.Load("../.env")
		if err != nil {
			panic(err)
		}
	}

	// 1. Test valid Token
	input := models.User{
		ID:       1,
		Username: "test@local.com",
		Password: "password",
		IsAdmin:  false,
	}

	token, _ := common.GenerateJWTToken(input)
	inputBearerToken := "Bearer " + string(token)
	mockFunction := new(models.MockedUser)
	mockFunction.On("GetDetailUser", 1).Return(
		&input,
		nil,
	)

	output, err := common.ValidateJWTToken(mockFunction, inputBearerToken)
	assert.Nil(t, err,
		"When Token is valid error should nil")
	assert.Equal(t, input.ID, output.ID,
		"User ID should same with user ID on database")
	assert.Equal(t, input.Username, output.Username,
		"User Username should same with user Username on database")
	assert.Equal(t, input.IsAdmin, output.IsAdmin,
		"User IsAdmin should same with user IsAdmin on database")

	// 2. Test Invalid Token no Bearer prefix
	input2 := models.User{
		ID:       1,
		Username: "test@local.com",
		Password: "password",
		IsAdmin:  false,
	}

	token2, _ := common.GenerateJWTToken(input2)
	inputBearerToken2 := string(token2) // no bearer prefix
	mockFunction2 := new(models.MockedUser)
	mockFunction2.On("GetDetailUser", 1).Return(
		&input2,
		nil,
	)

	output2, err := common.ValidateJWTToken(mockFunction2, inputBearerToken2)
	assert.NotNil(t, err,
		"When Token doesn't have prefix \"Bearer\" error should not nil")
	assert.Nil(t, output2,
		"When Token doesn't have prefix \"Bearer\" user should nil")

	// 3. Test Invalid Token not jwt token
	input3 := models.User{
		ID:       1,
		Username: "test@local.com",
		Password: "password",
		IsAdmin:  false,
	}

	inputBearerToken3 := "Bearer blablablabla"
	mockFunction3 := new(models.MockedUser)
	mockFunction3.On("GetDetailUser", 1).Return(
		&input3,
		nil,
	)

	output3, err := common.ValidateJWTToken(mockFunction3, inputBearerToken3)
	assert.NotNil(t, err,
		"When Token is not JWT Token error should not nil")
	assert.Nil(t, output3,
		"When Token is not JWT Token user should nil")

	// 4. User not found on database
	input4 := models.User{
		ID:       1,
		Username: "test@local.com",
		Password: "password",
		IsAdmin:  false,
	}
	token4, _ := common.GenerateJWTToken(input4)
	inputBearerToken4 := "Bearer " + string(token4)
	mockFunction4 := new(models.MockedUser)
	mockFunction4.On("GetDetailUser", 1).Return(
		&models.User{},
		errors.New("User not Found"),
	)

	output4, err := common.ValidateJWTToken(mockFunction4, inputBearerToken4)
	assert.NotNil(t, err,
		"When user not found in database error should not nil")
	assert.Nil(t, output4,
		"When user not found in database user should nil")
}
