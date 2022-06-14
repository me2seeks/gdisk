package logic

import (
	"cloud-disk/app/disk/model"
	"cloud-disk/app/user/cmd/rpc/user"
	"cloud-disk/common/xerr"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type StoreDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStoreDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StoreDetailLogic {
	return &StoreDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 判断是否能存入
func (l *StoreDetailLogic) StoreDetail(in *pb.StoreDetailReq) (*pb.StoreDetailResp, error) {
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
		return nil, errors.Wrapf(xerr.NewErrMsg("请联系管理员"), "ERROR: Failed to 获取用户仓库信息 uid: %d", in.Uid)
	}
	var pbStoreDetail pb.StoreDetailResp
	_ = copier.Copy(&pbStoreDetail, storeDetail)

	return &pbStoreDetail, nil
}
