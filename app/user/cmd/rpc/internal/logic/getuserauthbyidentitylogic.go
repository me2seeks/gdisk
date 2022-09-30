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

type GetUserAuthByIdentityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAuthByIdentityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAuthByIdentityLogic {
	return &GetUserAuthByIdentityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAuthByIdentityLogic) GetUserAuthByIdentity(in *pb.GetUserAuthByIdentityReq) (*pb.GetUserAuthByIdentityResp, error) {
	userAuth, err := l.svcCtx.UserAuthModel.FindOneByIdentityAuthType(l.ctx, in.Identity, in.AuthType)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user auth  failed"), "err : %v , in : %+v", err, in)
	}

	var respUserAuth user.UserAuth
	_ = copier.Copy(&respUserAuth, userAuth)
	return &pb.GetUserAuthByIdentityResp{
		UserAuth: &respUserAuth,
	}, nil
}
