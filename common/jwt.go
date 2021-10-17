package common

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/models"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

func GenerateJWTToken(user models.User) ([]byte, error) {
	token := jwt.New()
	token.Set("id", user.ID)
	token.Set("username", user.Username)
	signedToken, err := jwt.Sign(token, jwa.HS512, []byte("qwerty"))
	if err != nil {
		return nil, err
	}

	return signedToken, nil
}

func ValidateJWTToken(authorizationToken string) (*models.User, error) {
	// authorizationToken : Bearer {jwt token}
	// Parsing header token
	words := strings.Fields(authorizationToken)
	if len(words) != 2 {
		return nil, errors.New("invalid token")
	}
	jwtToken := words[1]
	fmt.Println(jwtToken)

	// validate jwt token
	resultToken, err := jwt.Parse([]byte(jwtToken), jwt.WithVerify(jwa.HS512, []byte("qwerty")))
	if err != nil {
		return nil, errors.New("invalid token")
	}

	// Get Jwt Payload
	id, isExist := resultToken.Get("id")
	if !isExist {
		return nil, errors.New("invalid token")
	}
	username, isExist := resultToken.Get("username")
	if !isExist {
		return nil, errors.New("invalid token")
	}

	// Get User from payload and validate it
	user := models.User{}
	if err := models.DBConn.First(&user, id).Error; err != nil {
		return nil, err
	}
	if user.Username != username {
		return nil, errors.New("invalid token")
	}

	return &user, nil
}
