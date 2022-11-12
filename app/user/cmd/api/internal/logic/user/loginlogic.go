package logic

import (
	"cloud-disk/app/user/cmd/rpc/pb"
	"cloud-disk/common/tool"
	"cloud-disk/common/xerr"
	"context"

	"github.com/pkg/errors"

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
	//loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginReq{
	//	AuthType: model.UserAuthTypeSystem,
	//	AuthKey:  req.Email,
	//	Password: req.Password,
	//})
	var identity string
	var err error

	//switch req.AuthType {
	//case model.UserAuthTypeSystem:
	//	identity, err = l.loginByEmail(req.Email, req.Password)
	//case model.UserAuthTypeSmallWX:
	//	identity= l.loginByMiniWx()
	//default:
	//	return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	//}
	identity, err = l.loginByEmail(req.Email, req.Password)

	if err != nil {
		return nil, err
	}

	tokenResp, err := l.svcCtx.UserRpc.GenerateToken(l.ctx, &pb.GenerateTokenReq{
		Identity: identity,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "ERROR GenerateToken userId : %d", identity)
	}
	return &types.LoginResp{
		Token:        tokenResp.Token,
		RefreshToken: tokenResp.RefreshToken,
	}, nil

}

func (l *LoginLogic) loginByEmail(email, password string) (string, error) {

	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, email)
	if err != nil {
		return "", errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR 根据Email查询用户信息失败，email:%s,err:%v", email, err)
	}

	if user == nil || tool.Md5ByString(password) != user.Password {
		return "", errors.Wrap(xerr.NewErrCode(xerr.ErrUserPwdError), "用户名或密码错误")
	}
	return user.Identity, nil
}

func (*LoginLogic) loginByMiniWx() error {
	//待定
	return nil
}
