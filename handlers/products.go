package handlers

import (
	"github.com/ak-yudha/lili_test/database"
	"github.com/ak-yudha/lili_test/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Div Rhino Trivia App!")
}

func CreateProduct(c *fiber.Ctx) error {
	fact := new(models.Product)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}

func ListProduct(c *fiber.Ctx) error {
	products := []models.Product{}
	name := c.FormValue("name")
	delivery_condition := c.FormValue("delivery_condition")

	if name == "" && delivery_condition == "" {
		database.DB.Db.Order("created_at desc").Find(&products)
	} else {
		database.DB.Db.Where("name LIKE ? ", "%"+name+"%").Or("delivery_condition LIKE ? ", "%"+delivery_condition+"%").Find(&products)
	}

	return c.Status(200).JSON(fiber.Map{"Data": products})
}
