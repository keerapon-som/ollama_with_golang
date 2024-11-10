package keycloak

type GetAccessTokenRequest struct {
	BASEURL      string `json:"base_url"`
	REALM        string `json:"realm"`
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

type GetAllUsersRequest struct {
	BASEURL     string `json:"base_url"`
	REALM       string `json:"realm"`
	BearerToken string `json:"bearer_token"`
}

type GetAllUsersResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Enabled  bool   `json:"enabled"`
	Email    string `json:"email"`
}

type CreateANewUserRequest struct {
	BASEURL     string `json:"base_url"`
	REALM       string `json:"realm"`
	BearerToken string `json:"bearer_token"`
}

type PayloadCreateNewUser struct {
	Username  string `json:"username"`
	Enabled   bool   `json:"enabled"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Keycloak interface {
	GetAccessToken(req GetAccessTokenRequest) (GetAccessTokenResponse, error)
	GetAllUsers(req GetAllUsersRequest) ([]GetAllUsersResponse, error)
	CreateANewUser(req CreateANewUserRequest, payload PayloadCreateNewUser) (GetAllUsersResponse, error)
}
