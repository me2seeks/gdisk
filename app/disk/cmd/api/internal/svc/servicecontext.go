package svc

import (
	"cloud-disk/app/disk/cmd/api/internal/config"
	"cloud-disk/app/disk/cmd/rpc/disk"
	"cloud-disk/app/disk/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	DiskRpc     disk.Disk
	RedisClient *redis.Redis
	Engine      *gorm.DB // gorm
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DiskRpc: disk.NewDisk(zrpc.MustNewClient(c.DiskRpcConf)),
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		Engine: model.Init(c.DB.DataSource),
	}
}
