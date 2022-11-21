package svc

import (
	"cloud-disk/app/disk/cmd/api/internal/config"
	"github.com/hibiken/asynq"
)

// create asynq client.
func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass})
}
