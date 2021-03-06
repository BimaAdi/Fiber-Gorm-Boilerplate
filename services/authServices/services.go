package authservices

import (
	"errors"

	"github.com/BimaAdi/fiberGormBoilerPlate/common"
	"github.com/BimaAdi/fiberGormBoilerPlate/models"
	authserializers "github.com/BimaAdi/fiberGormBoilerPlate/serializers/authSerializers"
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
