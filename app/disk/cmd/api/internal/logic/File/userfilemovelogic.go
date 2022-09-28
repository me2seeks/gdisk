package File

import (
	"cloud-disk/app/disk/model"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	resp = new(types.UserFileMoveReply)
	// parentId
	parentData := new(model.UserRepository)
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND user_identity = ?", req.ParentIdentity, userIdentity).
		First(parentData).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db failed error: %v", err)
	}
	if parentData.Id == 0 {
		return nil, errors.New("文件夹不存在")
	}

	// 更新记录的ParentId
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND deleted_at IS NULL", req.Identity).
		Update("parent_id", int64(parentData.Id)).Error
	if err != nil {

		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db failed error: %v", err)
	}
	return resp, nil
}
