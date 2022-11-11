package File

import (
	"cloud-disk/app/disk/model"
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/uuid"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFolderCreateLogic {
	return &PublicFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFolderCreateLogic) PublicFolderCreate(req *types.PublicFolderCreateRequest) (resp *types.PublicFolderCreateReply, err error) {
	u := ctxdata.GetUidFromCtx(l.ctx)
	// 判断当前文件名在该层级下是否已存在
	var cnt int64
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("name = ? AND parent_id = ? AND uid = ?", req.Name, req.ParentId, u).
		Count(&cnt).Error

	resp = new(types.PublicFolderCreateReply)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db failed error: %v", err)
	}
	if cnt > 0 {
		return nil, errors.New("文件已存在")
	}

	resp.Identity = uuid.UUID()
	// 创建文件夹
	data := &model.RepositoryPool{
		Identity: resp.Identity,
		Owner:    u,
		ParentId: req.ParentId,
		Name:     req.Name,
		IsPublic: 1,
	}
	err = l.svcCtx.Engine.
		Table("repository_pool").
		Select("identity", "owner", "name", "parent_id", "created_at", "updated_at", "is_public").
		Create(data).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db failed error: %v", err)
	}
	return resp, nil
}
