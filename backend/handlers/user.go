package handlers

import (
	"fmt"

	"github.com/Nissekissen/GO-testing/database"
	"github.com/Nissekissen/GO-testing/models"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {

	// Get access token from cookie and get corresponding user from db
	var token models.Token
	if err := database.DB.DB.Where("access_token = ?", c.Cookies("accesstoken")).First(&token).Error; err != nil {
		fmt.Println("Unauthorized from /user")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	var user models.User
	if err := database.DB.DB.Where("id = ?", token.UserID).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	return c.Status(200).JSON(user)
}
