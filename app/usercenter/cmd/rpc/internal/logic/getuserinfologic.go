package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"trytry/app/usercenter/cmd/rpc/usercenter"
	"trytry/app/usercenter/model"
	"trytry/common/xerr"

	"trytry/app/usercenter/cmd/rpc/internal/svc"
	"trytry/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Uid)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR GetUserInfo find user pb field, id:%d , err:%v", in.Uid, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "id: %d", in.Uid)
	}
	var respUser usercenter.User
	_ = copier.Copy(&respUser, user)
	return &pb.GetUserInfoResp{
		User: &respUser,
	}, nil
}
