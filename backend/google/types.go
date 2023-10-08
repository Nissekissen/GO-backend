package google

type GoogleUser struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Picture        string `json:"picture"`
	Verified_email bool   `json:"verified_email"`
	Name           string `json:"name"`
	FirstName      string `json:"given_name"`
	LastName       string `json:"family_name"`
	Age            int    `json:"birthday"`
}

type GoogleToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"` // seconds
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	IDToken      string `json:"id_token"`
}
