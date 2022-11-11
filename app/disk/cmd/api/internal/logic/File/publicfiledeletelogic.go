package File

import (
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"time"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublicFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublicDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublicFileDeleteLogic {
	return &PublicFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublicFileDeleteLogic) PublicDelete(req *types.PublicFileDeleteRequest) (*types.PublicFileDeleteReply, error) {
	u := ctxdata.GetUidFromCtx(l.ctx)

	err := l.svcCtx.Engine.
		Table("repository_pool").
		Where("identity = ? AND owner = ?", req.Identity, u).
		Update("del_state", globalkey.DelStateYes).
		Update("deleted_at", time.Now()).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新repository_pool失败")
	}

	if err != nil {
		return nil, err
	}
	return &types.PublicFileDeleteReply{}, nil
}
