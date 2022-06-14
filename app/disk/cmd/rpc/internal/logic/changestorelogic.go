package logic

import (
	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/app/disk/model"
	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
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
	store := new(model.Store)
	if in.CurrentSize > in.MaxSize {
		return nil, errors.Wrapf(xerr.NewErrMsg("当前容量不允许超过最大容量"), "ERROR 当前容量不允许超过最大容量 uid: %d", in.Uid)
	}
	if in.DelState == globalkey.DelStateYes {
		store.DelState = globalkey.DelStateYes
		store.DeleteTime = time.Now()
	}
	store.UserId = in.Uid
	store.CurrentSize = in.CurrentSize
	store.MaxSize = in.MaxSize

	err := l.svcCtx.StoreModel.UpdateWithVersion(l.ctx, nil, store)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: Failed to StoreModel.UpdateWithVersion err: %v", err)
	}

	return &pb.ChangeStoreResp{}, nil
}
