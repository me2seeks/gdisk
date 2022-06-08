package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"trytry/app/usercenter/cmd/rpc/usercenter"
	"trytry/app/usercenter/model"

	"trytry/app/usercenter/cmd/api/internal/svc"
	"trytry/app/usercenter/cmd/api/internal/types"

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
	loginResp, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &usercenter.LoginReq{
		AuthType: model.UserAuthTypeSystem,
		AuthKey:  req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	var resp types.LoginResp
	_ = copier.Copy(&resp, loginResp)
	return &resp, nil
}
