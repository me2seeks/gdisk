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

	userInfo, err := l.svcCtx.UserModel.FindOneByIdentity(l.ctx, in.Identity)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR GetUserInfo find user pb field, id:%d , err:%v", in.Identity, err)
	}
	if userInfo == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "id: %d", in.Identity)
	}

	var respUser user.UserDetail
	_ = copier.Copy(&respUser, userInfo)
	respUser.CreatedAt = userInfo.CreateAt.String()
	return &pb.GetUserInfoResp{
		User: &respUser,
	}, nil
}
