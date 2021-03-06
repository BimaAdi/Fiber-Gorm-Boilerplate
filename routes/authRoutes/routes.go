package authroutes

import (
	"github.com/BimaAdi/fiberGormBoilerPlate/common"
	"github.com/BimaAdi/fiberGormBoilerPlate/models"
	authserializers "github.com/BimaAdi/fiberGormBoilerPlate/serializers/authSerializers"
	authservices "github.com/BimaAdi/fiberGormBoilerPlate/services/authServices"
	"github.com/gofiber/fiber/v2"
)

func LoginRoute(c *fiber.Ctx) error {
	requestFormat := new(authserializers.LoginSerializer)

	if err := c.BodyParser(requestFormat); err != nil {
		return err
	}

	userRepo := models.User{}
	res, err := authservices.LoginService(userRepo, requestFormat.Username, requestFormat.Password)
	if err != nil {
		if err.Error() == "record not found" || err.Error() == "wrong password" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Wrong Credential",
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	return c.Status(200).JSON(res)
}

func LogoutRoute(c *fiber.Ctx) error {
	headerAuthorization := c.Get("Authorization")
	userRepo := models.User{}
	requestUser, err := common.ValidateJWTToken(userRepo, headerAuthorization)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	serializer := authserializers.UserModelToLogoutSuccessSerializer(*requestUser)

	return c.Status(200).JSON(serializer)
}
