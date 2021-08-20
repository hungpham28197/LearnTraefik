package session

import (
	"mainsite/config"
	"time"

	"github.com/TechMaster/eris"
	"github.com/mitchellh/mapstructure"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

const (
	SESSION_COOKIE = "sessid"
	SESS_AUTH      = "authenticated"
	SESS_USER      = "user"
)

/*
Lưu thông tin về người đăng nhập sau khi đăng nhập thành công.
Cấu trúc này sẽ lưu vào session
*/
type AuthenInfo struct {
	User  string
	Email string
	Roles []string
}

var Sess = sessions.New(sessions.Config{
	Cookie:       SESSION_COOKIE,
	AllowReclaim: true,
	Expires:      time.Hour * 48, /*Có giá trị trong 2 ngày*/
})

func InitSession() *redis.Database {
	redisConfig := config.Config.Redis

	redisDB := redis.New(redis.Config{
		Network:   redisConfig.Network,
		Addr:      redisConfig.Addr,
		Password:  redisConfig.Password,
		Database:  redisConfig.Database,
		MaxActive: redisConfig.MaxActive,
		Timeout:   time.Duration(redisConfig.IdleTimeout) * time.Minute,
		Prefix:    redisConfig.Prefix,
		Driver:    redis.GoRedis(),
	})

	Sess.UseDatabase(redisDB)
	return redisDB
}
func SetAuthenticated(ctx iris.Context, authenInfo AuthenInfo) {
	sess := sessions.Get(ctx)
	sess.Set(SESS_AUTH, true)
	sess.Set(SESS_USER, authenInfo)
}

func GetAuthInfo(ctx iris.Context) (*AuthenInfo, error) {
	data := sessions.Get(ctx).Get(SESS_USER)
	if data == nil {
		return nil, nil
	}

	authinfo := new(AuthenInfo)
	if err := mapstructure.Decode(data, authinfo); err != nil {
		return nil, eris.NewFrom(err)
	}
	return authinfo, nil
}

func IsLogin(ctx iris.Context) bool {
	login, _ := sessions.Get(ctx).GetBoolean(SESS_AUTH)
	return login
}
