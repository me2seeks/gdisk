package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"trytry/app/usercenter/cmd/rpc/internal/config"
	"trytry/app/usercenter/model"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),

		UserModel:     model.NewUserModel(conn, c.Cache),
		UserAuthModel: model.NewUserAuthModel(conn, c.Cache),
	}
}
