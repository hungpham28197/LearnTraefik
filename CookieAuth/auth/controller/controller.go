package controller

import (
	"auth/model"
	"auth/repo"
	"errors"
	"fmt"

	"github.com/TechMaster/eris"
	"github.com/TechMaster/logger"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

/*
Lưu thông tin đăng nhập từ client gửi lên
*/
type LoginRequest struct {
	Email string
	Pass  string
}

func ShowHomePage(ctx iris.Context) {
	if authinfo, err := GetAuthInfo(ctx); err == nil {
		ctx.ViewData("authinfo", authinfo)
	}
	_ = ctx.View("index")
}

func ShowSecret(ctx iris.Context) {
	// Check if user is authenticated
	if !IsLogin(ctx) {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	_, _ = ctx.WriteString("Secret Page")
}

/*
Login thông qua form. Dành cho ứng dụng web server side renderings
*/
func Login(ctx iris.Context) {
	var loginReq LoginRequest

	if err := ctx.ReadForm(&loginReq); err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := repo.QueryByEmail(loginReq.Email)
	if err != nil { //Không tìm thấy user
		_, _ = ctx.WriteString("Login Failed")
		return
	}

	if user.Pass != loginReq.Pass {
		_, _ = ctx.WriteString("Wrong password")
		return
	}

	session := sessions.Get(ctx)
	session.Set(SESS_AUTH, true)
	session.Set(SESS_USER, model.AuthenInfo{
		User:  user.User,
		Email: user.Email,
		Roles: user.Roles,
	})
	ctx.SetCookieKV("iris", "foo")
	//Login thành công thì quay về trang chủ
	ctx.Redirect("/")
}

/*
Login thông qua axios.post dành cho ứng dụng Vue
Request.ContentType = 'application/json'
*/
func LoginJSON(ctx iris.Context) {
	var loginReq LoginRequest

	if err := ctx.ReadJSON(&loginReq); err != nil {
		logger.Log(ctx, eris.NewFrom(err).BadRequest())
		return
	}

	user, err := repo.QueryByEmail(loginReq.Email)
	if err != nil { //Không tìm thấy user
		logger.Log(ctx, eris.Warning("User not found").UnAuthorized())
		return
	}

	if user.Pass != loginReq.Pass {
		logger.Log(ctx, eris.Warning("Wrong password").UnAuthorized())
		return
	}

	session := sessions.Get(ctx)
	session.Set(SESS_AUTH, true)
	session.Set(SESS_USER, model.AuthenInfo{
		User:  user.User,
		Email: user.Email,
		Roles: user.Roles,
	})
	//Login thành công thì quay về trang chủ
	_, _ = ctx.JSON("Login successfully")
}

func Logout(ctx iris.Context) {
	session := sessions.Get(ctx)
	//Revoke users authentication
	session.Clear()
	ctx.RemoveCookie(SESSION_COOKIE)
	ctx.Redirect("/")
}

func IsLogin(ctx iris.Context) bool {
	login, _ := sessions.Get(ctx).GetBoolean(SESS_AUTH)
	return login
}

func GetAuthInfo(ctx iris.Context) (*model.AuthenInfo, error) {
	data := sessions.Get(ctx).Get(SESS_USER)
	if authinfo, ok := data.(model.AuthenInfo); ok {
		return &authinfo, nil
	}
	return nil, errors.New("User not yet login")
}
