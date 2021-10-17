package authroutes

import (
	"github.com/BimaAdi/fiberPostgresqlBoilerPlate/common"
	authserializers "github.com/BimaAdi/fiberPostgresqlBoilerPlate/serializers/authSerializers"
	authservices "github.com/BimaAdi/fiberPostgresqlBoilerPlate/services/authServices"
	"github.com/gofiber/fiber/v2"
)

func LoginRoute(c *fiber.Ctx) error {
	requestFormat := new(authserializers.LoginSerializer)

	if err := c.BodyParser(requestFormat); err != nil {
		return err
	}

	res, err := authservices.LoginService(requestFormat.Username, requestFormat.Password)
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
	requestUser, err := common.ValidateJWTToken(headerAuthorization)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	serializer := authserializers.UserModelToLogoutSuccessSerializer(*requestUser)

	return c.Status(200).JSON(serializer)
}
