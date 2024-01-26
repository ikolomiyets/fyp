package oauth2

import "time"

// Jwks is the structure storing JWKS keys returned by the provider
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// AccessToken response from the OAuth2 server
type AccessToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

// UserInfo response from the OAuth2 userinfo endpoint
type UserInfo struct {
	Subject         string    `json:"sub"`
	Nickname        string    `json:"nickname"`
	Name            string    `json:"name"`
	Picture         string    `json:"picture"`
	UpdatedAt       time.Time `json:"updated_at"`
	Email           string    `json:"email"`
	IsEmailVerified bool      `json:"email_verified"`
}

// JSONWebKeys JWK details
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}
