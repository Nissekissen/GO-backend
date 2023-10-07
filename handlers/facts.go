package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Nissekissen/GO-testing/models"
	"github.com/Nissekissen/GO-testing/database"
)

func ListFacts(c *fiber.Ctx) error {
	var facts []models.Fact
	database.DB.DB.Find(&facts)
	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.DB.Create(&fact)

	return c.Status(fiber.StatusCreated).JSON(fact)
}
