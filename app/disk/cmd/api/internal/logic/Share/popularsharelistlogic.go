package Share

import (
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PopularShareListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPopularShareListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PopularShareListLogic {
	return &PopularShareListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PopularShareListLogic) PopularShareList(_ *types.PopularShareListRequest) (resp *types.PopularShareListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
