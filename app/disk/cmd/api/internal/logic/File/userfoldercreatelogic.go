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

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest) (resp *types.UserFolderCreateReply, err error) {
	userIdentity := ctxdata.GetUidFromCtx(l.ctx)

	if req.Name == "" {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.REUQEST_PARAM_ERROR), "name is empty")
	}

	// 判断当前文件名在该层级下是否已存在
	var cnt int64
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("name = ? AND parent_id = ? AND uid = ?", req.Name, req.ParentId, userIdentity).
		Count(&cnt).Error

	resp = new(types.UserFolderCreateReply)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db failed error: %v", err)
	}
	if cnt > 0 {
		return nil, errors.New("文件已存在")
	}

	resp.Identity = uuid.UUID()
	// 创建文件夹
	data := &model.UserRepository{
		Identity: resp.Identity,
		Uid:      userIdentity,
		ParentId: req.ParentId,
		Name:     req.Name,
	}
	err = l.svcCtx.Engine.
		Table("user_repository").
		Select("identity", "name", "user_identity", "parent_id", "created_at", "updated_at").
		Create(data).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db failed error: %v", err)
	}
	return resp, nil
}
