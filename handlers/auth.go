package handlers

import (
	"context"
	"fmt"

	"github.com/Nissekissen/GO-testing/database"
	"github.com/Nissekissen/GO-testing/google"
	"github.com/Nissekissen/GO-testing/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

func Authenticate(c *fiber.Ctx) error {
	var token models.Token
	if err := database.DB.DB.Where("access_token = ?", c.Cookies("accesstoken")).First(&token).Error; err != nil {
		fmt.Println("Unauthorized from middleware")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return c.Next()
}

// Refresh the access token using the refresh token and google api
func Refresh(c *fiber.Ctx) error {

	redirectUrl := c.Query("redirectUrl")

	// Get token from db
	var token models.Token
	if err := database.DB.DB.Where("access_token = ?", c.Cookies("accesstoken")).First(&token).Error; err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Refresh token
	newToken, err := google.GoogleOauthConfig.TokenSource(context.Background(), &oauth2.Token{
		RefreshToken: token.RefreshToken,
	}).Token()
	if err != nil {
		return c.JSON(err)
	}

	// Update token in db
	database.DB.DB.Model(&token).Update("access_token", newToken.AccessToken)

	// Set new cookie
	c.Cookie(&fiber.Cookie{
		Name:     "accesstoken",
		Value:    newToken.AccessToken,
		HTTPOnly: true,
		Secure:   true,
	})

	if redirectUrl != "" {
		return c.Redirect(redirectUrl, fiber.StatusTemporaryRedirect)
	}

	return c.JSON(newToken)
}
