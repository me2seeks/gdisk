package File

import (
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"cloud-disk/app/disk/model"
	"cloud-disk/common/xerr"
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileListLogic {
	return &PublicFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileListLogic) PublicFileList(_ *types.PublicFileListRequest) (*types.PublicFileListReply, error) {
	resp := types.PublicFileListReply{}

	err := l.svcCtx.Engine.Table("repository_pool").
		Select("repository_pool.name,repository_pool.ext ,repository_pool.id,repository_pool.parent_id,"+
			"repository_pool.path,repository_pool.size,repository_pool.updated_at,"+
			"user.name as owner,repository_pool.identity as repository_identity").
		Where("is_public", 1).
		Where("repository_pool.deleted_at IS NULL").
		Joins("left join user on repository_pool.owner = user.identity").
		Find(&resp.List).Error

	l.svcCtx.Engine.Table("repository_pool").
		Where("is_public", 1).
		Count(&resp.Count)

	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR 查询repository_pool list失败err:%v", err)
	}
	return &resp, nil
}
