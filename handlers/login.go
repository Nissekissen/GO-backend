package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Nissekissen/GO-testing/google"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

type User struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Picture        string `json:"picture"`
	Verified_email bool   `json:"verified_email"`
}

var (
	oauthStateString = "pseudo-random"
)

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func Login(c *fiber.Ctx) error {
	url := google.GetConfig().AuthCodeURL(oauthStateString)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

func Callback(c *fiber.Ctx) error {
	user, err := getUserInfo(c.Query("state"), c.Query("code"))
	if err != nil {
		fmt.Println(err.Error())
		return c.Redirect("/", fiber.StatusTemporaryRedirect)
	}

	fmt.Println(user)

	return c.JSON(user)

}

func getUserInfo(state string, code string) (*User, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := google.GetConfig().Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	user := User{}

	err = getJson("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken, &user)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	fmt.Println(user)

	return &user, nil
}
