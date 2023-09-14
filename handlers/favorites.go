package handlers

import (
	"github.com/ak-yudha/lili_test/database"
	"github.com/ak-yudha/lili_test/models"
	"github.com/gofiber/fiber/v2"
)

func CreateFavorites(c *fiber.Ctx) error {
	fact := new(models.Favorites)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}
