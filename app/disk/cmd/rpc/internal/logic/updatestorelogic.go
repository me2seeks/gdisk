package logic

import (
	"cloud-disk/app/disk/model"
	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"time"

	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStoreLogic {
	return &UpdateStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新store大小
func (l *UpdateStoreLogic) UpdateStore(in *pb.UpdateStoreReq) (*pb.UpdateStoreResp, error) {
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

	return &pb.UpdateStoreResp{}, nil
}
