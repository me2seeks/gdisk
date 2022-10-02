package Statistics

import (
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StatisticsShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStatisticsShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatisticsShareLogic {
	return &StatisticsShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StatisticsShareLogic) StatisticsShare(req *types.StatisticsShareRequest) (resp *types.StatisticsShareReply, err error) {
	// todo: add your logic here and delete this line

	return
}
