package svc

import (
	"cloud-disk/app/community/cmd/api/internal/config"
	"cloud-disk/app/community/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config

	RedisClient *redis.Redis
	Engine      *gorm.DB // gorm
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		Engine: model.Init(c.DB.DataSource),
	}
}
