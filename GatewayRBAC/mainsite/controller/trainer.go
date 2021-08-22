package controller

import "github.com/kataras/iris/v12"

func CreateLesson(ctx iris.Context) {
	_, _ = ctx.WriteString(ctx.HandlerName())
}
