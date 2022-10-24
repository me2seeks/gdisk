package logic

import (
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/xerr"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"

	"cloud-disk/app/user/cmd/rpc/internal/svc"
	"cloud-disk/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.Identity)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOKEN_GENERATE_ERROR), "getJwtToken err userId:%d , err:%v", in.Identity, err)
	}
	refreshToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire+31536000, in.Identity)

	return &pb.GenerateTokenResp{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secretKey string, iat, seconds int64, identity string) (string, error) {

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[ctxdata.CtxKeyJwtUserId] = identity
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
