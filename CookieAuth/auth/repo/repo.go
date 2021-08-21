package repo

import (
	"auth/model"
	"errors"
)

var users = []model.User{
	{
		User:  "admin",
		Pass:  "123",
		Email: "admin@gmail.com",
		Roles: []string{"admin", "trainer"},
	},
	{
		User:  "long",
		Pass:  "456",
		Email: "long@gmail.com",
		Roles: []string{"student"},
	},
	{
		User:  "linh",
		Pass:  "456",
		Email: "linh@gmail.com",
		Roles: []string{"student", "sale"},
	},
}

func QueryByEmail(email string) (user *model.User, err error) {
	for _, obj := range users {
		if obj.Email == email {
			user = new(model.User)
			*user = obj
			return user, nil
		}
	}
	return nil, errors.New("User not found")
}
