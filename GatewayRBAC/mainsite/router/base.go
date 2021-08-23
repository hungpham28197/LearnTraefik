package router

import (
	"mainsite/controller"
	"mainsite/rbac"

	"github.com/kataras/iris/v12"
)

func RegisterRoute(app *iris.Application) {
	app.Get("/", controller.ShowHomePage) //Không dùng rbac có nghĩa là public method

	blog := app.Party("/blog")
	{
		blog.Get("/", controller.GetAllPosts) //Không dùng rbac có nghĩa là public method
		rbac.Get(blog, "/all", rbac.AllowAll(), controller.GetAllPosts)
		rbac.Get(blog, "/create", rbac.Forbid(rbac.MAINTAINER, rbac.SYSOP), controller.GetAllPosts)
		rbac.Get(blog, "/{id:int}", rbac.Allow(rbac.AUTHOR, rbac.EDITOR), controller.GetPostByID)
		rbac.Get(blog, "/delete/{id:int}", rbac.Allow(rbac.ADMIN, rbac.AUTHOR, rbac.EDITOR), controller.DeletePostByID)
		rbac.Any(blog, "/any", rbac.Allow(rbac.SYSOP), controller.PostMiddleware)
	}

	student := app.Party("/student")
	{
		rbac.Get(student, "/submithomework", rbac.Allow(rbac.STUDENT), controller.SubmitHomework)
	}

	trainer := app.Party("/trainer")
	{
		rbac.Get(trainer, "/createlesson", rbac.Allow(rbac.TRAINER), controller.CreateLesson)
	}

	sysop := app.Party("/sysop")
	{
		rbac.Get(sysop, "/backupdb", rbac.Allow(rbac.SYSOP), controller.BackupDB)
		rbac.Get(sysop, "/upload", rbac.Allow(rbac.MAINTAINER, rbac.SYSOP), controller.ShowUploadForm)
		rbac.Post(sysop, "/upload", rbac.Allow(rbac.MAINTAINER, rbac.SYSOP, rbac.SALE), iris.LimitRequestBodySize(300000), controller.UploadPhoto)
	}

	sales := app.Party("/sale")
	{
		rbac.Get(sales, "/runads", rbac.Allow(rbac.SALE), controller.RunAdvertise)
	}

}
