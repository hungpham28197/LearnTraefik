package session

import (
	"auth/config"
	"time"

	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

const (
	SESSION_COOKIE = "sessid"
	SESS_AUTH      = "authenticated"
	SESS_USER      = "user"
	AUTHINFO       = "authinfo"
)

/*
Cấu hình Session Manager
*/
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
