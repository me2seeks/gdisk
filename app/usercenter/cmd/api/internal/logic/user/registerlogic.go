package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"trytry/app/usercenter/cmd/api/internal/svc"
	"trytry/app/usercenter/cmd/api/internal/types"
	"trytry/app/usercenter/cmd/rpc/usercenter"
	"trytry/app/usercenter/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	//验证
	if res, _ := l.svcCtx.RedisClient.Get(req.Phone); req.Captcha != res {
		return nil, errors.Wrapf(nil, "验证码错误")
	}

	registerResp, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &usercenter.RegisterReq{
		Phone:    req.Phone,
		Password: req.Password,
		AuthKey:  req.Phone,
		AuthType: model.UserAuthTypeSystem,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "UsercenterRpc.Register err req: %+v", req)
	}
	var resp types.RegisterResp
	_ = copier.Copy(&resp, registerResp)
	return &resp, nil
}
