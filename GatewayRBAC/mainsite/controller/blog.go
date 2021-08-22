package controller

import "github.com/kataras/iris/v12"

func GetAllPosts(ctx iris.Context) {
	_, _ = ctx.WriteString(ctx.HandlerName())
}

func GetPostByID(ctx iris.Context) {
	_, _ = ctx.WriteString(ctx.HandlerName())
}

func DeletePostByID(ctx iris.Context) {
	_, _ = ctx.WriteString(ctx.HandlerName())
}

func CreatePost(ctx iris.Context) {
	_, _ = ctx.WriteString(ctx.HandlerName())
}

func PostMiddleware(ctx iris.Context) {
	_, _ = ctx.WriteString(ctx.HandlerName())
}
