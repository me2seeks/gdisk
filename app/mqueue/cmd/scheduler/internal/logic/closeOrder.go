package logic

import (
	"cloud-disk/app/mqueue/cmd/job/jobtype"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

func (l *MqueueScheduler) closeOrder() {

	task := asynq.NewTask(jobtype.DeferRemoveDeletedObject, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("*/1 * * * *", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!M queueSchedulerErr!!! ====> 【removeObject】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【removeObject】 registered an  entry: %q \n", entryID)
}
