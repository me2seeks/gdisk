package svc

import (
	"cloud-disk/app/mqueue/cmd/scheduler/internal/config"
	"fmt"
	"github.com/hibiken/asynq"
	"time"
)

// create scheduler
func newScheduler(c config.Config) *asynq.Scheduler {

	location, _ := time.LoadLocation("Asia/Shanghai")
	return asynq.NewScheduler(
		asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
		}, &asynq.SchedulerOpts{
			Location: location,
			EnqueueErrorHandler: func(task *asynq.Task, opts []asynq.Option, err error) {
				fmt.Printf("Scheduler EnqueueErrorHandler <<<<<<<===>>>>> err : %+v , task : %+v", err, task)
			},
		})
}
