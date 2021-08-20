package controller

import (
	"auth/session"
	"fmt"

	"github.com/kataras/iris/v12"
)

func Authenticate(ctx iris.Context) {
	fmt.Println("Referer: ", ctx.GetReferrer())
	fmt.Println("Path: ", ctx.Path())
	fmt.Println("RouteName: ", ctx.RouteName())

	if session.IsLogin(ctx) {
		ctx.Next() //Cho phép đi tiếp
	} else {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
}
