// Public models dùng chung giữa các dự án và module
package pmodel

//Danh sách các role
type Roles map[int]bool

//Thông tin tài khoản
type User struct {
	User  string
	Pass  string
	Email string
	Roles Roles
}

/*
Lưu thông tin về người đăng nhập sau khi đăng nhập thành công.
Cấu trúc này sẽ lưu vào session
*/
type AuthenInfo struct {
	User  string
	Email string
	Roles Roles //kiểu map[int]bool
}
