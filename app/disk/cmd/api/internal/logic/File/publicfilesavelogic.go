package File

import (
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
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
	_ = ctxdata.GetUidFromCtx(l.ctx)
	l.svcCtx.Engine.Table("repository_pool").
		Where("identity = ? ", req.RepositoryIdentity).
		Update("is_public", 1)

	return resp, nil
}
