package validator

import (
	"server/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var ValidatorUser = validator.New()

func ValidateUser(c *fiber.Ctx) error {
	var errors []*models.IError
	var body models.UserModel

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := ValidatorUser.Struct(body); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element models.IError
			element.Field = err.Field()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}
