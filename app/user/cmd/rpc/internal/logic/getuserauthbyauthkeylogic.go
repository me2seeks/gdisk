package logic

import (
	"cloud-disk/app/user/cmd/rpc/user"
	"cloud-disk/app/user/model"
	"cloud-disk/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"cloud-disk/app/user/cmd/rpc/internal/svc"
	"cloud-disk/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAuthByAuthKeyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByAuthKeyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByAuthKeyLogic {
	return &GetUserAuthByAuthKeyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByAuthKeyLogic) GetUserAuthByAuthKey(in *pb.GetUserAuthByAuthKeyReq) (*pb.GetUserAuthByAuthKeyResp, error) {

	userAuth, err := l.svcCtx.UserAuthModel.FindOneByAuthTypeAuthKey(l.ctx, in.AuthType, in.AuthKey)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user auth  failed"), "err : %v , in : %+v", err, in)
	}

	var respUserAuth user.UserAuth
	_ = copier.Copy(&respUserAuth, userAuth)

	return &pb.GetUserAuthByAuthKeyResp{
		UserAuth: &respUserAuth,
	}, nil
}
