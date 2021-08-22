package repo

import (
	"auth/pmodel"
	"auth/rbac"
	"errors"
)

var users = []pmodel.User{
	{
		User:  "admin",
		Pass:  "1",
		Email: "admin@gmail.com",
		Roles: pmodel.Roles{rbac.ADMIN: true},
	},
	{
		User:  "huy",
		Pass:  "1",
		Email: "huy@gmail.com",
		Roles: pmodel.Roles{rbac.TRAINER: true},
	},
	{
		User:  "hien",
		Pass:  "1",
		Email: "hien@gmail.com",
		Roles: pmodel.Roles{rbac.SYSOP: true, rbac.TRAINER: true},
	},
	{
		User:  "hung",
		Pass:  "1",
		Email: "hung@gmail.com",
		Roles: pmodel.Roles{rbac.STUDENT: true},
	},
	{
		User:  "man",
		Pass:  "1",
		Email: "man@gmail.com",
		Roles: pmodel.Roles{rbac.SALE: true, rbac.EDITOR: true},
	},

	{
		User:  "vuong",
		Pass:  "1",
		Email: "vuong@gmail.com",
		Roles: pmodel.Roles{rbac.EMPLOYER: true},
	},
}

func QueryByEmail(email string) (user *pmodel.User, err error) {
	for _, obj := range users {
		if obj.Email == email {
			user = new(pmodel.User)
			*user = obj
			return user, nil
		}
	}
	return nil, errors.New("User not found")
}
