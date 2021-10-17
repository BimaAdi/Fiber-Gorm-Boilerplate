package authserializers

import "github.com/BimaAdi/fiberPostgresqlBoilerPlate/models"

type LoginSerializer struct {
	Username string `json:"username" validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}

type LoginSuccessSerializer struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	Token    string `json:"token"`
}

type LogoutSuccessSerializer struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
}

func UserModelToLoginSuccessSerializer(model models.User, token []byte) LoginSuccessSerializer {
	return LoginSuccessSerializer{
		Id:       int(model.ID),
		Username: model.Username,
		IsAdmin:  model.IsAdmin,
		Token:    string(token),
	}
}

func UserModelToLogoutSuccessSerializer(model models.User) LogoutSuccessSerializer {
	return LogoutSuccessSerializer{
		Id:       int(model.ID),
		Username: model.Username,
		IsAdmin:  model.IsAdmin,
	}
}
