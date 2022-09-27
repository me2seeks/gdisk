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
	var userId int64
	var err error
	switch in.AuthType {
	case model.UserAuthTypeSystem:
		userId, err = l.loginByemail(in.AuthKey, in.Password)
	default:
		return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	if err != nil {
		return nil, err
	}
	//2、生成token 不用rpc调用
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&user.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "ERROR GenerateToken userId : %d", userId)
	}

	return &user.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByemail(email, password string) (int64, error) {

	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR 根据手机号查询用户信息失败，email:%s,err:%v", email, err)
	}
	if user == nil {
		return 0, errors.Wrapf(ErrUserNoExistsError, "email:%s", email)
	}
	if tool.Md5ByString(password) == user.Password {
		return 0, errors.Wrap(ErrUsernamePwdError, "密码匹配错误")
	}
	return user.Id, nil
}

func (l *LoginLogic) loginByMiniWx() error {
	//待定
	return nil
}
