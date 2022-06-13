package logic

import (
	"cloud-disk/app/mqueue/cmd/job/jobtype"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *MqueueScheduler) closeOrder() {

	task := asynq.NewTask(jobtype.DeferCloseHomestayOrder, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("*/1 * * * *", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!M queueSchedulerErr!!! ====> 【closeOrder】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【closeOrder】 registered an  entry: %q \n", entryID)
}
