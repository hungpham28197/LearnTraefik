package controller

import (
	"auth/repo"
	"auth/session"
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
	authinfo, err := session.GetAuthInfo(ctx)

	if err != nil {
		logger.Log(ctx, err)
	}

	if authinfo != nil {
		ctx.ViewData("authinfo", authinfo)
	}
	_ = ctx.View("index")
}

func ShowSecret(ctx iris.Context) {
	// Check if user is authenticated
	if !session.IsLogin(ctx) {
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

	session.SetAuthenticated(ctx, session.AuthenInfo{
		User:  user.User,
		Email: user.Email,
		Roles: user.Roles,
	})

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

	session.SetAuthenticated(ctx, session.AuthenInfo{
		User:  user.User,
		Email: user.Email,
		Roles: user.Roles,
	})

	sess := sessions.Get(ctx)
	sess.Set(session.SESS_AUTH, true)
	sess.Set(session.SESS_USER, session.AuthenInfo{
		User:  user.User,
		Email: user.Email,
		Roles: user.Roles,
	})
	//Login thành công thì quay về trang chủ
	_, _ = ctx.JSON("Login successfully")
}

func Logout(ctx iris.Context) {
	sess := sessions.Get(ctx)
	//Revoke users authentication
	sess.Clear()
	ctx.RemoveCookie(session.SESSION_COOKIE)
	ctx.Redirect("/")
}

func LogoutJSON(ctx iris.Context) {
	if !session.IsLogin(ctx) {
		ctx.StatusCode(iris.StatusUnauthorized)
		_, _ = ctx.JSON("You are not login yet")
	}

	sess := sessions.Get(ctx)
	//Revoke users authentication
	sess.Clear()
	ctx.RemoveCookie(session.SESSION_COOKIE)
	_, _ = ctx.JSON("Logout success")
}
