package logic

import (
	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/app/disk/model"
	"cloud-disk/app/user/cmd/rpc/user"
	"cloud-disk/common/xerr"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailStoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailStoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailStoreLogic {
	return &DetailStoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// store 详情
func (l *DetailStoreLogic) DetailStore(in *pb.StoreDetailReq) (*pb.StoreDetailResp, error) {
	_, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Uid: in.Uid,
	})
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	builder := l.svcCtx.StoreModel.RowBuilder().Where(squirrel.Eq{
		"user_id": in.Uid,
	})
	storeDetail, err := l.svcCtx.StoreModel.FindOneByQuery(l.ctx, builder)
	if err != nil {
		if err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: Failed to StoreModel.FindOneByQuery error: %v", err)
		}
		l.svcCtx.StoreModel.Insert(l.ctx, nil, &model.Store{
			UserId:  in.Uid,
			MaxSize: 10240,
		})
	}

	var pbStoreDetail pb.StoreDetailResp
	_ = copier.Copy(&pbStoreDetail.Store, storeDetail)
	pbStoreDetail.Store.DeleteTime = storeDetail.DeleteTime.Unix()

	return &pbStoreDetail, nil
}
