package repository

import "myauth/ent"

type AuthRepository struct {
	client *ent.Client
}

func (auth *AuthRepository) RegisterUser(user *ent.User) *ent.UserCreate {
	return auth.client.User.Create().SetID(user.ID).SetUsername(user.Username).SetPassword(user.Password).SetEmail(user.Email)
}
