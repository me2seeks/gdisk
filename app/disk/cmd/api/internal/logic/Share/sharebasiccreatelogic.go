package Share

import (
	"cloud-disk/app/disk/model"
	"cloud-disk/common/uuid"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	resp = new(types.ShareBasicCreateReply)
	idna := uuid.UUID()
	usr := new(model.UserRepository)
	err = l.svcCtx.Engine.
		Where("identity = ?", req.UserRepositoryIdentity).
		First(usr).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db failed error: %v", err)
	}
	if usr.Id == 0 {

		return nil, errors.New("not found")
	}

	data := &model.ShareBasic{
		Identity:               idna,
		UserIdentity:           userIdentity,
		UserRepositoryIdentity: req.UserRepositoryIdentity,
		RepositoryIdentity:     usr.RepositoryId,
		ExpiredTime:            req.ExpiredTime,
		Desc:                   req.Desc,
	}
	err = l.svcCtx.Engine.
		Select("identity", "user_identity", "repository_identity", "user_repository_identity", "expired_time", "created_at", "updated_at").
		Create(data).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "db failed error: %v", err)
	}

	resp.Identity = idna
	return resp, nil
}
