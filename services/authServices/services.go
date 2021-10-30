package authservices

import (
	"errors"

	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/common"
	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/models"
	authserializers "github.com/BimaAdi/fiberPostgresqlBoilerPlate/serializers/authSerializers"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(ui models.UserInterface, username string, password string) (*authserializers.LoginSuccessSerializer, error) {
	user, err := ui.GetDetailUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// compare the hash
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("wrong password")
	}

	signedToken, err := common.GenerateJWTToken(*user)
	if err != nil {
		return nil, err
	}

	// serialize user model
	serializer := authserializers.UserModelToLoginSuccessSerializer(*user, signedToken)

	return &serializer, nil
}
