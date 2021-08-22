package controller

import (
	"github.com/TechMaster/eris"
	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
)

func BackupDB(ctx iris.Context) {
	_, _ = ctx.WriteString(ctx.HandlerName())
}

//GET /upload
func ShowUploadForm(ctx iris.Context) {
	_ = ctx.View("upload")
}

/*
POST /upload
Viết hàm upload ảnh vào đây
*/
func UploadPhoto(ctx iris.Context) {
	uploadfiles, _, err := ctx.UploadFormFiles("./uploads")
	if err != nil {
		logger.Log(ctx, eris.NewFrom(err))
	}
	filenames := "Upload successful \n"
	for _, upload := range uploadfiles {
		filenames = filenames + upload.Filename + "/n"
	}
	_, _ = ctx.WriteString(filenames)
}
