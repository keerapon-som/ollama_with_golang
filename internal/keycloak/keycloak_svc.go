package keycloak

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type keycloak struct {
}

func NewKeycloak() Keycloak {
	return &keycloak{}
}

func (k *keycloak) GetAccessToken(req GetAccessTokenRequest) (response GetAccessTokenResponse, err error) {
	url := req.BASEURL + "realms/" + req.REALM + "/protocol/openid-connect/token"
	method := "POST"

	payload := strings.NewReader("client_id=" + req.ClientID + "&client_secret=" + req.ClientSecret + "&grant_type=client_credentials")

	client := &http.Client{}
	httpReq, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return GetAccessTokenResponse{}, ErrMarshal
	}
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		return GetAccessTokenResponse{}, ErrRequestFailed
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("Get Access Token failed with status code: ", res.StatusCode)
		return GetAccessTokenResponse{}, ErrStatuscode
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return GetAccessTokenResponse{}, ErrReadbody
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err)
		return GetAccessTokenResponse{}, ErrUnmarshal
	}

	return response, nil

}

func (k *keycloak) GetAllUsers(req GetAllUsersRequest) (response []GetAllUsersResponse, err error) {
	url := req.BASEURL + "admin/realms/" + req.REALM + "/users"
	method := "GET"
	client := &http.Client{}
	httpReq, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, ErrMarshal
	}

	httpReq.Header.Add("Authorization", "Bearer "+req.BearerToken)
	res, err := client.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		return nil, ErrRequestFailed
	}

	if res.StatusCode != http.StatusOK {
		return nil, ErrStatuscode
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, ErrReadbody
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return nil, ErrUnmarshal
	}

	return response, nil
}

func (k *keycloak) CreateANewUser(req CreateANewUserRequest, payload PayloadCreateNewUser) (GetAllUsersResponse, error) {
	url := req.BASEURL + "admin/realms/" + req.REALM + "/users"
	method := "POST"

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return GetAllUsersResponse{}, ErrMarshal
	}

	// bodyPayload := strings.NewReader(string(jsonData))

	client := &http.Client{}
	httpReq, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return GetAllUsersResponse{}, ErrRequestFailed
	}

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("Authorization", "Bearer "+req.BearerToken)

	res, err := client.Do(httpReq)
	if err != nil {
		fmt.Println(err)
		return GetAllUsersResponse{}, ErrDoFailed
	}

	if res.StatusCode != http.StatusCreated {
		fmt.Println("Create a new user failed with status code: ", res.StatusCode)
		return GetAllUsersResponse{}, ErrStatuscode
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return GetAllUsersResponse{}, ErrReadbody
	}

	var response GetAllUsersResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return GetAllUsersResponse{}, ErrUnmarshal
	}

	return response, nil
}
