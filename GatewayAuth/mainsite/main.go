package main

import (
	"mainsite/config"
	"mainsite/controller"
	"mainsite/session"
	"mainsite/template"

	"github.com/TechMaster/logger"
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
	app.Use(session.Sess.Handler())

	public := app.Party("/public/")
	{
		public.Get("/", controller.ShowHomePage)
	}

	app.Get("/upload", controller.ShowUploadForm)
	app.Post("/upload", controller.UploadPhoto)

	template.InitViewEngine(app)
	_ = app.Listen(config.Config.Port)
}
