package File

import (
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	resp = new(types.UserFileNameUpdateReply)
	if req.Name == "" {
		return nil, errors.New("文件名为空")
	}

	// 判断当前文件名在该层级下是否已存在
	var cnt int64
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("name = ?", req.Name).
		Where("parent_id = (select parent_id from user_repository ur where ur.identity = ?)", req.Identity).
		Count(&cnt).Error

	if err != nil {

		return
	}
	if cnt > 0 {
		return nil, errors.New("文件名已存在")
	}

	// 更新文件名
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).
		Update("name", req.Name).Error

	if err != nil {

		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db failed error: %v", err)
	}

	return resp, nil
}
