package logic

import (
	"cloud-disk/app/user/cmd/rpc/user"
	"cloud-disk/app/user/model"
	"context"
	"github.com/jinzhu/copier"

	"cloud-disk/app/user/cmd/api/internal/svc"
	"cloud-disk/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
		AuthType: model.UserAuthTypeSystem,
		AuthKey:  req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	var resp types.LoginResp
	_ = copier.Copy(&resp, loginResp)
	return &resp, nil
}
