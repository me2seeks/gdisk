package logic

import (
	"cloud-disk/app/user/cmd/rpc/user"
	"cloud-disk/app/user/model"
	"cloud-disk/common/tool"
	"cloud-disk/common/uuid"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"cloud-disk/app/user/cmd/rpc/internal/svc"
	"cloud-disk/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: RPC[user] UserModel.FindOneByPhone verify:%s,err:%v", in.Email, err)
	}
	if err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.ErrUserAlreadyRegisterError), "ERROR: RPC[user]  用户已存在 :%s,err:%v", in.Email, err)
	}

	var identity string
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		u := new(model.User)
		var identity = uuid.UUID()
		u.Email = in.Email
		u.Name = in.Nickname
		u.Identity = identity

		if len(u.Name) == 0 {
			u.Name = tool.Krand(8, tool.KC_RAND_KIND_ALL)
		}
		if len(in.Password) > 0 {
			u.Password = tool.Md5ByString(in.Password)
		}

		_, err := l.svcCtx.UserModel.Insert(l.ctx, session, u)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: RPC[user] UserModel.Insert err::%v,u:%+v", err, u)
		}

		//lastId, err := insertResult.LastInsertId()
		//if err != nil {
		//	return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: RPC[user] insertResult.LastInsertId:%v,u:%+v", err, u)
		//}

		userAuth := new(model.UserAuth)
		userAuth.Identity = u.Identity
		userAuth.AuthKey = in.AuthKey
		userAuth.AuthType = in.AuthType
		if _, err := l.svcCtx.UserAuthModel.Insert(l.ctx, session, userAuth); err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: RPC[user] user_auth Insert err: %v", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		Identity: identity,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "ERROR: RPC[user] 获取token错误: %d", identity)
	}

	return &pb.RegisterResp{
		Token:        tokenResp.Token,
		RefreshToken: tokenResp.RefreshToken,
	}, nil
}
