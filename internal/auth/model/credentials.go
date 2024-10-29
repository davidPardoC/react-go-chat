package model

type Credentials struct {
	AccesToken   string `json:"acces_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
