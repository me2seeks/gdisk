package logic

import (
	"cloud-disk/app/user/cmd/rpc/user"
	"cloud-disk/common/ctxdata"
	"context"
	"github.com/jinzhu/copier"

	"cloud-disk/app/user/cmd/api/internal/svc"
	"cloud-disk/app/user/cmd/api/internal/types"

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

func (l *DetailLogic) Detail(req *types.UserInfoReq) (*types.UserInfoResp, error) {
	identity := ctxdata.GetUidFromCtx(l.ctx)

	userInfoResp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Identity: identity,
	})
	if err != nil {
		return nil, err
	}

	var userInfo types.User
	_ = copier.Copy(&userInfoResp, userInfo)
	return &types.UserInfoResp{
		UserInfo: userInfo,
	}, nil
}
