package Share

import (
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserShareListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserShareListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserShareListLogic {
	return &UserShareListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserShareListLogic) UserShareList(_ *types.UserShareListRequest) (resp *types.UserShareListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
