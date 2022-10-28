package Share

import (
	"cloud-disk/common/globalkey"
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareStatisticsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareStatisticsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareStatisticsLogic {
	return &ShareStatisticsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareStatisticsLogic) ShareStatistics(req *types.ShareStatisticsRequest) (*types.ShareStatisticsReply, error) {
	var s int64
	var c []int64
	//var sum int64
	l.svcCtx.Engine.Table("share_basic").Where("share_basic.del_state", globalkey.DelStateNo).Count(&s)
	l.svcCtx.Engine.Table("share_basic").Pluck("click_num", &c)

	var sum int64 = 0

	for _, v := range c {
		sum = sum + v
	}

	return &types.ShareStatisticsReply{
		ShareCount: s,
		ClickNum:   sum,
	}, nil
}
