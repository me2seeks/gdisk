package store

import (
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.StoreDetailReq) (*types.StoreDetailResp, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	storeDetail, err := l.svcCtx.DiskRpc.DetailStore(l.ctx, &pb.StoreDetailReq{
		Uid: uid,
	})
	if err != nil {
		return nil, err
	}
	if storeDetail.Store.DelState == globalkey.DelStateYes {
		return nil, errors.Wrapf(xerr.NewErrMsg("仓库已被删除 请联系管理员"), "ERROR: Failed to 获取用户仓库信息 uid: %d", uid)
	}

	return &types.StoreDetailResp{
		CurrentSize: storeDetail.Store.CurrentSize,
		MaxSize:     storeDetail.Store.MaxSize,
	}, nil

}
