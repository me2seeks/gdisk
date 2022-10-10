package File

import (
	"cloud-disk/app/define"
	"cloud-disk/app/disk/model"
	"cloud-disk/common/uuid"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, UserIdentity string) (resp *types.UserRepositorySaveReply, err error) {
	// 用户新增文件
	usr := &model.UserRepository{
		Identity:           uuid.UUID(),
		Uid:                UserIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Name:               req.Name,
		Ext:                req.Ext,
	}

	resp = new(types.UserRepositorySaveReply)
	var Size struct {
		TotalSize int `json:"total_size"`
	}
	l.svcCtx.Engine.
		Table("user_repository").
		Select("sum(repository_pool.size) as total_size").
		Where("user_repository.uid = ? AND user_repository.deleted_at IS NULL", UserIdentity).
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Take(&Size)
	if Size.TotalSize >= define.UserRepositoryMaxSize {

		return nil, errors.New("容量不足")
	}

	var count int64
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("name = ? AND pid = ? AND uid = ? AND deleted_at IS NULL", req.Name, req.ParentId, UserIdentity).
		Count(&count).Error
	if count > 0 {
		return nil, errors.New("exist")
	}

	err = l.svcCtx.Engine.
		Select("identity", "pid", "uid", "repository_identity", "name", "ext", "created_at", "updated_at").
		Create(usr).Error
	if err != nil {

		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db failed error: %v", err)
	}

	return
}
