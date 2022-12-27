package File

import (
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"cloud-disk/app/disk/model"
	"cloud-disk/common/ctxdata"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicFileSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileSaveLogic {
	return &PublicFileSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileSaveLogic) PublicFileSave(req *types.PublicFileSaveRequest) (resp *types.PublicFileSaveReply, err error) {
	u := ctxdata.GetUidFromCtx(l.ctx)
	l.svcCtx.Engine.Table("repository_pool").
		Where("identity = ? ", req.RepositoryIdentity).
		Updates(model.RepositoryPool{IsPublic: 1, Owner: u, ParentId: req.ParentId})

	return resp, nil
}
