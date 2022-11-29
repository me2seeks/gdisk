package logic

import (
	"cloud-disk/app/disk/cmd/rpc/disk"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/app/mqueue/cmd/job/internal/svc"
	"cloud-disk/app/mqueue/cmd/job/jobtype"
	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

var ErrRemoveFal = xerr.NewErrMsg("remove object fail")

// RemoveDeletedObjectHandler 去除已经删除的对象
type RemoveDeletedObjectHandler struct {
	svcCtx *svc.ServiceContext
}

func NewRemoveDeletedObjectHandler(svcCtx *svc.ServiceContext) *RemoveDeletedObjectHandler {
	return &RemoveDeletedObjectHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask   : if return err != nil , asynq will retry
func (l *RemoveDeletedObjectHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	var p jobtype.RemoveDeletedObjectPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(ErrRemoveFal, "ERROR RemoveDeletedObjectHandler payload err:%v, payLoad:%+v", err, t.Payload())
	}

	l.svcCtx.DiskRpc.UpdateFile(ctx, &disk.UpdateFileReq{
		FileDetail: &pb.FileDetail{
			Identity: p.Identity,
			DelState: globalkey.DelStateYes,
		},
	})
	return nil
}
