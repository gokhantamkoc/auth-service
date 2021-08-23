package util

import (
	"context"
	"github.com/Nerzal/gocloak/v8"
	"github.com/gokhantamkoc/auth-service/api"
)

type KeycloakClientWrapper struct {
	client 	gocloak.GoCloak
	ctx		context.Context
	token	*gocloak.JWT
}

func (k *KeycloakClientWrapper) NewKeycloakClient(address string) {
	k.ctx = context.Background()
	k.client = gocloak.NewClient(address)
}

func (k *KeycloakClientWrapper) LoginAdmin(username string, password string, realmName string) error {
	var err error
	k.token, err = k.client.LoginAdmin(k.ctx, username, password, realmName)
	if err != nil {
		return err
	}
	return nil
}
func (k *KeycloakClientWrapper) Logout(realmName string) error {
	err := k.client.LogoutUserSession(k.ctx, k.token.AccessToken, realmName, k.token.SessionState)
	if err != nil {
		return err
	}
	return nil
}

func (k KeycloakClientWrapper) GetUsers(realmName string, first int, max int) ([]api.User, error) {
	params := gocloak.GetUsersParams{
		First:               gocloak.IntP(first),
		Max:                 gocloak.IntP(max),
	}

	keycloakUsers, err := k.client.GetUsers(k.ctx, k.token.AccessToken, realmName, params)
	if err != nil {
		return nil, err
	}

	var users []api.User
	for _, keycloakUser := range keycloakUsers {
		user := api.User{
			ID: *keycloakUser.ID,
			FirstName: *keycloakUser.FirstName,
			LastName: *keycloakUser.LastName,
			Email: *keycloakUser.Email,
			Username: *keycloakUser.Username,
			Attributes: keycloakUser.Attributes,
		}
		users = append(users, user)
	}
	return users, err
}

func (k KeycloakClientWrapper) GetUserByID(ID string, realmName string) (api.User, error) {
	user, err := k.client.GetUserByID(k.ctx, k.token.AccessToken, realmName, ID)
	foundUser := api.User{
		ID: *user.ID,
		FirstName: *user.FirstName,
		LastName: *user.LastName,
		Email: *user.Email,
		Username: *user.Username,
	}
	return foundUser, err
}

func (k KeycloakClientWrapper) CreateUser(userToAdd api.User, realmName string) (string, error) {
	user := gocloak.User{
		FirstName:                  gocloak.StringP(userToAdd.FirstName),
		LastName:                   gocloak.StringP(userToAdd.LastName),
		Email:                      gocloak.StringP(userToAdd.Email),
		Enabled:                    gocloak.BoolP(true),
		Username:                   gocloak.StringP(userToAdd.Username),
		Attributes:					userToAdd.Attributes,
	}
	id, err := k.client.CreateUser(k.ctx, k.token.AccessToken, realmName, user)
	return id, err
}

func (k KeycloakClientWrapper) UpdateUser(userToUpdate api.User, realmName string) error {
	user := gocloak.User{
		ID: 						gocloak.StringP(userToUpdate.ID),
		FirstName:                  gocloak.StringP(userToUpdate.FirstName),
		LastName:                   gocloak.StringP(userToUpdate.LastName),
		Email:                      gocloak.StringP(userToUpdate.Email),
		Enabled:                    gocloak.BoolP(true),
		Username:                   gocloak.StringP(userToUpdate.Username),
		Attributes:                 userToUpdate.Attributes,
	}
	return k.client.UpdateUser(k.ctx, k.token.AccessToken, realmName, user)
}