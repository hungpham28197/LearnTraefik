package controller

import (
	"mainsite/rbac"
	"mainsite/session"

	"github.com/kataras/iris/v12"
)

func ShowHomePage(ctx iris.Context) {
	if authinfo := session.GetAuthInfoViewData(ctx); authinfo != nil {
		ctx.ViewData("roles", rbac.RolesNames(authinfo.Roles))
	}

	_ = ctx.View("index")
}
