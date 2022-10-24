package svc

import (
	"cloud-disk/app/user/cmd/api/internal/config"
	"cloud-disk/app/user/cmd/rpc/user"
	"cloud-disk/app/user/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UserRpc       user.User
	RedisClient   *redis.Redis
	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel

	SetUidToCtxMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		UserModel:     model.NewUserModel(conn, c.Cache),
		UserAuthModel: model.NewUserAuthModel(conn, c.Cache),
	}
}
