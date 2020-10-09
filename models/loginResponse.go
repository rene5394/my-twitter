package models

// LoginResponse has the token returns with the login
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
