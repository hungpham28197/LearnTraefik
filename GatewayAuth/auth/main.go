package main

import (
	"auth/config"
	"auth/controller"
	"auth/session"
	"auth/template"

	"github.com/TechMaster/logger"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	config.ReadConfig()

	logFile := logger.Init() //Cần phải có 2 file error.html và info.html ở /views
	if logFile != nil {
		defer logFile.Close()
	}

	redisDb := session.InitSession()
	defer redisDb.Close()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.UseRouter(crs)

	app.Use(session.Sess.Handler())

	app.Get("/", controller.ShowHomePage)
	app.Get("/secret", controller.ShowSecret)
	app.Post("/login", controller.Login)
	app.Post("/loginjson", controller.LoginJSON)
	app.Get("/logout", controller.Logout)
	app.Get("/logoutjson", controller.LogoutJSON)
	app.Any("/auth", controller.Authenticate)

	template.InitViewEngine(app)
	_ = app.Listen(config.Config.Port)
}
