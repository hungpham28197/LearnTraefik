package model

/*
Thông tin tài khoản
*/

type User struct {
	User  string
	Pass  string
	Email string
	Roles  []string
}

/*
Lưu thông tin về người đăng nhập sau khi đăng nhập thành công.
Cấu trúc này sẽ lưu vào session
*/
type AuthenInfo struct {
	User  string
	Email string
	Roles []string
}
