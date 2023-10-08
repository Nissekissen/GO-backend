package google

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	OauthStateString  = "pseudo-random" // TODO: change this to random to protect against CSRF
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:5000/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/user.birthday.read",
			"https://www.googleapis.com/auth/plus.login",
			"openid",
		},
		Endpoint: google.Endpoint,
	}
)
