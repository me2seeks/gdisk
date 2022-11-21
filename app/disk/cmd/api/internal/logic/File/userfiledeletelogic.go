package File

import (
	"cloud-disk/app/mqueue/cmd/job/jobtype"
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"time"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const RemoveTimeMinutes = 30 //defer remove object time

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest) (resp *types.UserFileDeleteReply, err error) {
	u := ctxdata.GetUidFromCtx(l.ctx)

	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND uid = ?", req.Identity, u).
		Update("del_state", globalkey.DelStateYes).
		Update("deleted_at", time.Now()).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新user_repository失败")
	}

	//2、Delayed remove of disk tasks.
	payload, err := json.Marshal(jobtype.RemoveDeletedObjectPayload{Identity: req.Identity})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("create defer close order task json Marshal fail err :%+v , sn : %s", err, req.Identity)
	} else {
		_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferRemoveDeletedObject, payload), asynq.ProcessIn(RemoveTimeMinutes*time.Minute))
		if err != nil {
			logx.WithContext(l.ctx).Errorf("create defer remove object task insert queue fail err :%+v , sn : %s", err, req.Identity)
		}
	}
	return nil, nil

}
