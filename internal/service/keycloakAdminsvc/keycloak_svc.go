package keycloakadminsvc

import (
	"fmt"
	"ollamawithgo/config"
	"ollamawithgo/internal/keycloak"
)

type keycloakAdminSVC struct {
	keycloak keycloak.Keycloak
}

func NewKeycloakAdminSVC(keycloak keycloak.Keycloak) KeycloakAdminSVC {
	return &keycloakAdminSVC{
		keycloak: keycloak,
	}
}

var (
	cacheAdminToken string
)

func (s *keycloakAdminSVC) GetAccessToken() (resp GetAccessTokenResponse, err error) {
	k := config.GetConfig().Keycloak
	req := keycloak.GetAccessTokenRequest{
		ClientID:     k.AdminClientID,
		ClientSecret: k.AdminClientSecret,
		GrantType:    "client_credentials",
		BASEURL:      k.BASEURL,
		REALM:        "master",
	}

	AccessTokenData, err := s.keycloak.GetAccessToken(req)
	if err != nil {
		return GetAccessTokenResponse{}, ErrGetAcessToken
	}

	resp = GetAccessTokenResponse{
		AccessToken:      AccessTokenData.AccessToken,
		ExpiresIn:        AccessTokenData.ExpiresIn,
		RefreshExpiresIn: AccessTokenData.RefreshExpiresIn,
		TokenType:        AccessTokenData.TokenType,
		NotBeforePolicy:  AccessTokenData.NotBeforePolicy,
		Scope:            AccessTokenData.Scope,
	}

	cacheAdminToken = resp.AccessToken

	return resp, nil
}

func (s *keycloakAdminSVC) GetAllUsers() (resp []GetAllUsersResponse, err error) {
	k := config.GetConfig().Keycloak
	req := keycloak.GetAllUsersRequest{
		BASEURL:     k.BASEURL,
		REALM:       "master",
		BearerToken: cacheAdminToken,
	}

	users, err := s.keycloak.GetAllUsers(req)
	if err != nil {
		resp, err := s.GetAccessToken() // user New Token
		req.BearerToken = resp.AccessToken
		if err != nil {
			return nil, ErrGetAcessToken
		}

		users, err = s.keycloak.GetAllUsers(req)
		if err != nil {
			return nil, ErrGetAllUsers
		}

		fmt.Println("Use New Token")
	}

	for _, user := range users {
		resp = append(resp, GetAllUsersResponse{
			ID:       user.ID,
			Username: user.Username,
			Enabled:  user.Enabled,
			Email:    user.Email,
		})
	}

	return resp, nil
}

func (s *keycloakAdminSVC) CreateANewUser(payload keycloak.PayloadCreateNewUser) (resp GetAllUsersResponse, err error) {
	k := config.GetConfig().Keycloak
	req := keycloak.CreateANewUserRequest{
		BASEURL:     k.BASEURL,
		REALM:       "master",
		BearerToken: cacheAdminToken,
	}

	user, err := s.keycloak.CreateANewUser(req, payload)
	if err != nil {
		resp, err := s.GetAccessToken() // user New Token

		req.BearerToken = resp.AccessToken
		if err != nil {
			return GetAllUsersResponse{}, ErrGetAcessToken
		}

		user, err = s.keycloak.CreateANewUser(req, payload)
		if err != nil {
			fmt.Println("Error Create New User", err)
			return GetAllUsersResponse{}, ErrCreateNewUser
		}
	}

	resp = GetAllUsersResponse{
		ID:       user.ID,
		Username: user.Username,
		Enabled:  user.Enabled,
		Email:    user.Email,
	}

	return resp, nil
}
