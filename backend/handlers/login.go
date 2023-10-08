package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"context"

	"github.com/Nissekissen/GO-testing/database"
	"github.com/Nissekissen/GO-testing/google"
	"github.com/Nissekissen/GO-testing/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

// Handler for /login
func Login(c *fiber.Ctx) error {

	// If user is already logged in, refresh the token and redirect to /
	if c.Cookies("accesstoken") != "" {
		// Refresh the access token
		var token models.Token
		if err := database.DB.DB.Where("access_token = ?", c.Cookies("accesstoken")).First(&token).Error; err != nil {
			// Cookie is invalid, delete cookie and redirect to /login
			c.Cookie(&fiber.Cookie{
				Name:     "accesstoken",
				Value:    "",
				HTTPOnly: true,
				Secure:   true,
			})
			return c.Redirect("/login", fiber.StatusTemporaryRedirect)
		}

		return c.Redirect("/refresh?redirectUrl=/", fiber.StatusTemporaryRedirect)
	}

	url := google.GoogleOauthConfig.AuthCodeURL(
		google.OauthStateString,
		oauth2.AccessTypeOffline,
	)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

// Handler for /auth/callback
func Callback(c *fiber.Ctx) error {

	user, token, err := getUserInfo(c.Query("state"), c.Query("code"))
	if err != nil {
		fmt.Println(err.Error())
		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	}

	// Check if user is already in database
	var dbUser models.User
	if err := database.DB.DB.Where("id = ?", user.ID).First(&dbUser).Error; err != nil {
		fmt.Println("User not found in database, creating new user...")
		dbUser = models.User{
			Name:    user.Name,
			Email:   user.Email,
			Picture: user.Picture,
			ID:      user.ID,
		}
		database.DB.DB.Create(&dbUser)
	} else {
		fmt.Println("User found in database, updating user...")
		database.DB.DB.Model(&dbUser).Updates(models.User{
			Name:    user.Name,
			Email:   user.Email,
			Picture: user.Picture,
		})
	}

	var dbToken models.Token
	if err := database.DB.DB.Where("user_id = ?", user.ID).First(&dbToken).Error; err != nil {
		fmt.Println("Token not found in database, creating new token...")
		// Store token
		dbToken = models.Token{
			UserID:       user.ID,
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		}
		database.DB.DB.Create(&dbToken)

	} else {
		fmt.Println("Token found in database, updating token...")
		database.DB.DB.Model(&dbToken).Updates(models.Token{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "accesstoken",
		Value:    token.AccessToken,
		HTTPOnly: true,
		Secure:   true,
	})

	return c.Redirect("/", fiber.StatusTemporaryRedirect)

}

// Returns a GoogleUser object from the state and code (provided from oauth2 callback)
func getUserInfo(state string, code string) (*google.GoogleUser, *oauth2.Token, error) {
	if state != google.OauthStateString {
		return nil, nil, fmt.Errorf("invalid oauth state")
	}

	token, err := google.GoogleOauthConfig.Exchange(context.TODO(), code)
	if err != nil {
		return nil, nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	user := google.GoogleUser{}

	err = getJson("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken, &user)
	if err != nil {
		return nil, nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	return &user, token, nil
}

func getJsonFromBody(body io.ReadCloser, target interface{}) error {
	return json.NewDecoder(body).Decode(target)
}

// make a GET request and return body as target object
func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer r.Body.Close()

	return getJsonFromBody(r.Body, target)
}
