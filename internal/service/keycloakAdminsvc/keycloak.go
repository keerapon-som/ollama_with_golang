package keycloakadminsvc

import "ollamawithgo/internal/keycloak"

type GetAccessTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type GetAccessTokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	Scope            string `json:"scope"`
}

type GetAllUsersResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Enabled  bool   `json:"enabled"`
	Email    string `json:"email"`
}

type KeycloakAdminSVC interface {
	GetAccessToken() (GetAccessTokenResponse, error)
	GetAllUsers() ([]GetAllUsersResponse, error)
	CreateANewUser(payload keycloak.PayloadCreateNewUser) (GetAllUsersResponse, error)
}
