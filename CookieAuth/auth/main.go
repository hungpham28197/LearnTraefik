package main

import (
	"auth/controller"
	"auth/template"

	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func main() {
	app := iris.New()
	logFile := logger.Init() //Cần phải có 2 file error.html và info.html ở /views
	if logFile != nil {
		defer logFile.Close()
	}
	sess := sessions.New(sessions.Config{
		Cookie:       controller.SESSION_COOKIE,
		AllowReclaim: true,
	})

	app.Use(sess.Handler())

	app.Get("/", controller.ShowHomePage)
	app.Get("/secret", controller.ShowSecret)
	app.Post("/login", controller.Login)
	app.Post("/loginjson", controller.LoginJSON)
	app.Get("/logout", controller.Logout)
	app.Any("/auth", controller.Authenticate)

	template.InitViewEngine(app)
	_ = app.Listen(":3000")
}
