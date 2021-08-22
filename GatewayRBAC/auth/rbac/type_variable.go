package rbac

import "auth/pmodel"

/*
Nếu thêm sửa xoá role thì cập nhật danh sách const này
*/
const (
	ADMIN      = 1
	STUDENT    = 2
	TRAINER    = 3
	SALE       = 4
	EMPLOYER   = 5
	AUTHOR     = 6
	EDITOR     = 7
	MAINTAINER = 8
	SYSOP      = 9 //Role mới
)

//Mảng này phải tương ứng với danh sách const khai báo ở trên
var allRoles = []int{ADMIN, STUDENT, TRAINER, SALE, EMPLOYER, AUTHOR, EDITOR, MAINTAINER, SYSOP}

//Dùng để in role kiểu int ra string cho dễ hiếu
var roleName = map[int]string{
	ADMIN:      "admin",
	STUDENT:    "student",
	TRAINER:    "trainer",
	SALE:       "sale",
	EMPLOYER:   "employer",
	AUTHOR:     "author",
	EDITOR:     "editor",
	MAINTAINER: "maintainer",
	SYSOP:      "system operator",
}

/*
Biểu thức hàm sẽ trả về danh sách role kiểu map[int]bool
*/
type RoleExp func() pmodel.Roles

/*
Ứng với một route = HTTP Verb + Path chúng ta có một map các role
Dùng để kiểm tra phân quyền
*/
var routesRoles = make(map[string]pmodel.Roles)

/*
pathsRoles có key là Path (không kèm HTTP Verb)
Dùng để in ra báo cáo cho dễ nhìn, vì các route chung một path sẽ được gom lại
*/
var pathsRoles = make(map[string]HTTPVerbRoles)

/*
kiểu HTTPVerbRoles là map có key là 'GET', 'POST', 'PUT', 'DELETE'
Value là map các role
HTTPVerbRoles dùng để gom các roles gán cho từng HTTP Verb ứng với một path
*/
type HTTPVerbRoles map[string]pmodel.Roles

/*
Danh sách các public routes dùng trong hàm CheckPermission
*/
var publicRoutes = make(map[string]bool)
