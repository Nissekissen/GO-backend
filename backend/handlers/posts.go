package handlers

import (
	"strconv"

	"github.com/Nissekissen/GO-testing/database"
	"github.com/Nissekissen/GO-testing/models"
	"github.com/gofiber/fiber/v2"
)

func GetPost(c *fiber.Ctx) error {

	postId := c.AllParams()["id"]
	if postId == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}

	// TODO: Get post from database
	var post models.Post
	if err := database.DB.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(post)
}

func GetPosts(c *fiber.Ctx) error {

	amount, err := strconv.Atoi(c.Query("amount"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	posts := []models.Post{}

	query := database.DB.DB.Model(&models.Post{})
	if amount > 0 {
		query.Limit(amount)
	}
	query.Find(&posts)

	return c.JSON(posts)
}

func CreatePost(c *fiber.Ctx) error {
	// Validate input
	type CreatePostInput struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	var input CreatePostInput
	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Get token data
	var token models.Token
	if err := database.DB.DB.Where("access_token = ?", c.Cookies("accesstoken")).First(&token).Error; err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create post
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  token.UserID,
	}
	database.DB.DB.Create(&post)

	return c.Status(fiber.StatusCreated).JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	// Get token data
	var token models.Token
	if err := database.DB.DB.Where("access_token = ?", c.Cookies("accesstoken")).First(&token).Error; err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Get post from db
	postId := c.AllParams()["id"]
	if postId == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}

	var post models.Post
	if err := database.DB.DB.Where("id = ?", postId).First(&post).Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Check if user owns post
	if post.UserID != token.UserID {
		return c.SendStatus(fiber.StatusForbidden)
	}

	// Delete post
	database.DB.DB.Delete(&post)

	return c.SendStatus(fiber.StatusNoContent)
}
