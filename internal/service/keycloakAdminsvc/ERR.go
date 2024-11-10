package keycloakadminsvc

import "errors"

var (
	ErrGetAcessToken = errors.New("get access token failed")
	ErrCreateNewUser = errors.New("create new user failed")
	ErrGetAllUsers   = errors.New("do failed")
)
