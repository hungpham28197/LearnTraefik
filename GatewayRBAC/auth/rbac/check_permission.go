package rbac

import (
	"auth/session"

	"github.com/TechMaster/eris"
	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
)

func CheckRoutePermission(ctx iris.Context) {
	route := ctx.GetCurrentRoute().String()
	authinfo, _ := session.GetAuthInfo(ctx)

	//Nếu route không thuộc nhóm public routes cần kiểm tra phân quyền
	if authinfo != nil {
		//Gán authinfo để cho handler tiếp theo sử dụng
		ctx.ViewData(session.AUTHINFO, authinfo)
	}

	//Nếu route nằm trong public routes thì cho qua luôn
	if publicRoutes[route] {
		ctx.Next()
		return
	}

	//Chưa đăng nhập mà đòi vào protected route, vậy phải kick ra
	if authinfo == nil {
		logger.Log(ctx, eris.Warning("Bạn chưa đăng nhập").UnAuthorized())
		return
	}

	//Nếu route không thuộc nhóm public routes cần kiểm tra phân quyền
	rolesInRoute := routesRoles[route]
	for role := range authinfo.Roles {
		if rolesInRoute[role] {
			//Khi route.roles và authinfo.Roles giao nhau có nghĩa là người dùng có quyền đi tiếp
			ctx.Next()
			return
		}
	}

	logger.Log(ctx, eris.Warning("Bạn không quyền thực hiện tác vụ này").UnAuthorized())
}
