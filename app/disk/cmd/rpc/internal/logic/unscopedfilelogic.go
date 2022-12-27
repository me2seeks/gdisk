package logic

import (
	"cloud-disk/app/disk/model"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"

	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnscopedFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnscopedFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnscopedFileLogic {
	return &UnscopedFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnscopedFileLogic) UnscopedFile(in *pb.UnscopedFileReq) (*pb.UnscopedFileResp, error) {
	err := l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND uid = ? AND del_state = ?", in.Identity).
		Where("repository_pool.deleted_at IS NOT NULL").
		Unscoped().Delete(model.UserRepository{})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新user_repository失败")
	}

	return &pb.UnscopedFileResp{}, nil
}
