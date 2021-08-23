package main

import (
	"auth/config"
	"auth/controller"
	"auth/rbac"
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

	rbacConfig := rbac.NewConfig()
	rbacConfig.RootAllow = true
	rbac.Init(rbacConfig) //Khởi động với cấu hình mặc định
	
	//đặt hàm này trên các hàm đăng ký route - controller
	app.Use(rbac.CheckRoutePermission)

	app.Get("/", controller.ShowHomePage)

	rbac.Get(app, "/secret", rbac.AllowAll(), controller.ShowSecret)

	app.Post("/login", controller.Login)
	app.Post("/loginjson", controller.LoginJSON)

	rbac.Get(app, "/logout", rbac.AllowAll(), controller.LogoutFromWeb)
	rbac.Get(app, "/logoutjson", rbac.AllowAll(), controller.LogoutFromREST)

	template.InitViewEngine(app)

	//Luôn để hàm này sau tất cả lệnh cấu hình đường dẫn với RBAC
	rbac.BuildPublicRoute(app)
	_ = app.Listen(config.Config.Port)
}
