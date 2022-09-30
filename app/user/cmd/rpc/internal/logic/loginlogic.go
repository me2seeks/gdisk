package logic

import (
	"cloud-disk/app/user/cmd/rpc/user"
	"cloud-disk/app/user/model"
	"cloud-disk/common/tool"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"

	"cloud-disk/app/user/cmd/rpc/internal/svc"
	"cloud-disk/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*user.LoginResp, error) {
	var identity string
	var err error
	switch in.AuthType {
	case model.UserAuthTypeSystem:
		identity, err = l.loginByEmail(in.AuthKey, in.Password)
	default:
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	if err != nil {
		return nil, err
	}
	//2、生成token 不用rpc调用
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		Identity: identity,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "ERROR GenerateToken userId : %d", identity)
	}

	return &user.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByEmail(email, password string) (string, error) {

	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err != nil && err != model.ErrNotFound {
		return "", errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR 根据Email查询用户信息失败，email:%s,err:%v", email, err)
	}
	if user == nil {
		return "", errors.Wrapf(ErrUserNoExistsError, "email:%s", email)
	}
	if tool.Md5ByString(password) == user.Password {
		return "", errors.Wrap(ErrUsernamePwdError, "密码匹配错误")
	}
	return user.Identity, nil
}

func (l *LoginLogic) loginByMiniWx() error {
	//待定
	return nil
}
