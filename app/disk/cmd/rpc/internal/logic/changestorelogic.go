package logic

import (
	"cloud-disk/app/disk/model"
	"context"
	"time"

	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangeStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeStoreLogic {
	return &ChangeStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 改变store大小
func (l *ChangeStoreLogic) ChangeStore(in *pb.ChangeStoreReq) (*pb.ChangeStoreResp, error) {
	l.svcCtx.StoreModel.UpdateWithVersion(l.ctx, nil, model.FileStore{
		UserId:     in.Uid,
		MaxSize:    0,
		Version:    0,
		DeleteTime: time.Time{},
		DelState:   0,
	})
	return &pb.ChangeStoreResp{}, nil
}
