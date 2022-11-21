package logic

import (
	"cloud-disk/app/mqueue/cmd/job/internal/svc"
	"cloud-disk/app/mqueue/cmd/job/jobtype"
	"context"
	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// register job
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	//defer job
	mux.Handle(jobtype.DeferRemoveDeletedObject, NewRemoveDeletedObjectHandler(l.svcCtx))

	//queue job , asynq support queue job
	// wait you fill..

	//中间件
	//// some middlewares
	//mux.Use(func(next asynq.Handler) asynq.Handler {
	//	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
	//		// just record a log
	//		fmt.Println(fmt.Printf("[%s] log - %+v", time.Now().Format("2006-01-02 15:04:05"), t.Payload))
	//
	//		return next.ProcessTask(ctx, t)
	//	})
	//})

	return mux
}
