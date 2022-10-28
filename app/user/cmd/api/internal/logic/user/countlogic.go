package logic

import (
	"cloud-disk/app/user/cmd/api/internal/svc"
	"cloud-disk/app/user/cmd/api/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type CountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountLogic {
	return &CountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountLogic) Count(req *types.RegisterCountReq) (*types.RegisterCountResp, error) {

	b := l.svcCtx.UserModel.CountBuilder("id")
	c, nil := l.svcCtx.UserModel.FindCount(l.ctx, b)

	return &types.RegisterCountResp{Count: c}, nil
}
