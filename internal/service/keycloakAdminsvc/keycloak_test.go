package keycloakadminsvc

import (
	"fmt"
	"ollamawithgo/internal/keycloak"
	"testing"
)

func TestGetAccessTokem(t *testing.T) {
	k := NewKeycloakAdminSVC(keycloak.NewKeycloak())
	resp, err := k.GetAccessToken()
	if err != nil {
		t.Error("Error while getting access token")
	}

	fmt.Println(resp)
}

func TestGetAllUsers(t *testing.T) {
	k := NewKeycloakAdminSVC(keycloak.NewKeycloak())
	resp, err := k.GetAllUsers()
	if err != nil {
		t.Error("Error while getting all users")
	}

	fmt.Println(resp)
}

func TestCreateANewUser(t *testing.T) {
	k := NewKeycloakAdminSVC(keycloak.NewKeycloak())
	payload := keycloak.PayloadCreateNewUser{
		Username:  "tabcdezxsx",
		Email:     "eidxasdc@gmail.com",
		Enabled:   true,
		FirstName: "tesxszxt",
		LastName:  "teswzzzt",
	}
	resp, err := k.CreateANewUser(payload)
	if err != nil {
		fmt.Println(err)
		t.Error("Error while creating a new user")
	}

	fmt.Println(resp)
}
