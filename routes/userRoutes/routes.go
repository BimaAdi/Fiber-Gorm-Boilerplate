package userroutes

import (
	"strconv"

	userserializers "github.com/BimaAdi/fiberPostgresqlBoilerPlate/serializers/userSerializers"
	userservices "github.com/BimaAdi/fiberPostgresqlBoilerPlate/services/userServices"
	"github.com/gofiber/fiber/v2"
)

func GetAllUserRoute(c *fiber.Ctx) error {

	query := new(userserializers.UserQueryParams)

	if err := c.QueryParser(query); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response, err := userservices.GetAllUserService(query.Page, query.Size)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(response)
}

func GetDetailUserRoute(c *fiber.Ctx) error {
	idInt, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "parameter should be integer",
		})
	}

	response, err := userservices.GetDetailUserService(idInt)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(404).JSON(fiber.Map{
				"message": "Data not found",
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	return c.JSON(response)
}

func CreateUserRoute(c *fiber.Ctx) error {
	requestFormat := new(userserializers.UserRequestSerializer)

	if err := c.BodyParser(requestFormat); err != nil {
		return err
	}

	errors := userserializers.ValidateUser(*requestFormat)
	if errors != nil {
		return c.Status(400).JSON(errors)

	}

	new_user, err := userservices.CreateUserService(*requestFormat)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(201).JSON(new_user)
}

func UpdateUserRoute(c *fiber.Ctx) error {
	// get query parameter
	idInt, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "parameter should be integer",
		})
	}

	// get json and validate
	requestFormat := new(userserializers.UserRequestSerializer)
	if err := c.BodyParser(requestFormat); err != nil {
		return err
	}
	errors := userserializers.ValidateUser(*requestFormat)
	if errors != nil {
		return c.Status(400).JSON(errors)

	}

	// update user
	response, err := userservices.UpdateUserService(idInt, *requestFormat)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(404).JSON(fiber.Map{
				"message": "Data not found",
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	return c.Status(200).JSON(response)
}

func DeleteUserRoute(c *fiber.Ctx) error {
	idInt, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "parameter should be integer",
		})
	}

	if err := userservices.DeleteUserService(idInt); err != nil {
		if err.Error() == "record not found" {
			return c.Status(404).JSON(fiber.Map{
				"message": "Data not found",
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}
	return c.SendStatus(204)
}
