package keycloak

import "errors"

var (
	ErrRequestFailed = errors.New("request failed")
	ErrDoFailed      = errors.New("do failed")
	ErrReadbody      = errors.New("read body failed")
	ErrUnmarshal     = errors.New("unmarshal failed")
	ErrMarshal       = errors.New("marshal failed")
	ErrStatuscode    = errors.New("status code not 200")
)
