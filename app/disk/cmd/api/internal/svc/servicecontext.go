package svc

import (
	"cloud-disk/app/disk/cmd/api/internal/config"
	"cloud-disk/app/disk/cmd/rpc/disk"
	"cloud-disk/app/disk/model"
	"cloud-disk/app/user/cmd/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DiskRpc disk.Disk
	UserRpc user.User

	RedisClient *redis.Redis
	Engine      *gorm.DB // gorm
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DiskRpc: disk.NewDisk(zrpc.MustNewClient(c.DiskRpcConf)),
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		Engine: model.Init(c.DB.DataSource),
	}
}
