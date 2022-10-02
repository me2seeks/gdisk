package Statistics

import (
	"cloud-disk/app/disk/model"
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatisticsFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatisticsFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticsFileLogic {
	return &StatisticsFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatisticsFileLogic) StatisticsFile(req *types.StatisticsFileRequest) (resp *types.StatisticsFileReply, err error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	//
	if uid != 0 {
		err = l.svcCtx.Engine.
			Table("user_repository").
			Where("uid = ? AND deleted_at IS NULL", uid).
			Count(&resp.Count).Error
		if err != nil && err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR 查询file数量失败，Uid: %s,err:%v", in.Uid, err)
		}
	}
	err = l.svcCtx.Engine.
		Table("repository_poll").
		Where("deleted_at IS NULL").
		Count(&resp.Count).Error
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR 查询file数量失败，err:%v", err)
	}

	return
}
