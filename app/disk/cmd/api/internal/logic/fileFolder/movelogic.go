package fileFolder

import (
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveLogic {
	return &MoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveLogic) Move(req *types.MovedReq) (resp *types.MovedResp, err error) {
	// todo: add your logic here and delete this line

	return
}
