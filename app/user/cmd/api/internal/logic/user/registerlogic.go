package logic

import (
	"cloud-disk/app/user/cmd/api/internal/svc"
	"cloud-disk/app/user/cmd/api/internal/types"
	"cloud-disk/app/user/cmd/rpc/user"
	"cloud-disk/app/user/model"
	"cloud-disk/common/tool"
	"cloud-disk/common/uuid"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

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
	if res, _ := l.svcCtx.RedisClient.Get(req.Email); req.Captcha != res || req.Captcha == "" {
		return nil, xerr.NewErrCode(xerr.VERIFY_ERROR)
	}

	_, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR:  UserModel.FindOneByPhone verify:%s,err:%v", req.Email, err)
	}
	//TODO mysql反应慢
	//if err != model.ErrNotFound {
	//	return nil, errors.Wrapf(xerr.NewErrCode(xerr.ErrUserAlreadyRegisterError), "ERROR:  用户已存在 :%s,err:%v", req.Email, err)
	//}

	var identity string
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		u := new(model.User)
		var identity = uuid.UUID()
		u.Email = req.Email
		u.Name = req.Name
		u.Identity = identity

		if len(u.Name) == 0 {
			u.Name = tool.Krand(8, tool.KC_RAND_KIND_ALL)
		}
		if len(req.Password) > 0 {
			u.Password = tool.Md5ByString(req.Password)
		}

		_, err := l.svcCtx.UserModel.Insert(l.ctx, session, u)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR:  UserModel.Insert err::%v,u:%+v", err, u)
		}

		//lastId, err := insertResult.LastInsertId()
		//if err != nil {
		//	return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR:  insertResult.LastInsertId:%v,u:%+v", err, u)
		//}

		userAuth := new(model.UserAuth)
		userAuth.Identity = u.Identity
		userAuth.AuthKey = req.Email
		//userAuth.AuthType = req.AuthType
		if _, err := l.svcCtx.UserAuthModel.Insert(l.ctx, session, userAuth); err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: User_auth Insert err: %v", err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	tokenResp, err := l.svcCtx.UserRpc.GenerateToken(l.ctx, &user.GenerateTokenReq{
		Identity: identity,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "ERROR:  获取token错误: %d", identity)
	}

	return &types.RegisterResp{
		Token:        tokenResp.Token,
		RefreshToken: tokenResp.RefreshToken,
	}, nil
}
