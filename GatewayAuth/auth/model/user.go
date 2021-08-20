package model

/*
Thông tin tài khoản
*/

type User struct {
	User  string
	Pass  string
	Email string
	Roles []string
}
