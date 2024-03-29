package logic

import (
	"cloud-disk/app/verification/cmd/rpc/internal/svc"
	"cloud-disk/app/verification/cmd/rpc/pb"
	"cloud-disk/common/tool"
	"cloud-disk/common/verify"
	"cloud-disk/common/xerr"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyemailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyemailLogic {
	return &VerifyemailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyemailLogic) VerifyEmail(in *pb.VerifyEmailReq) (*pb.VerifyEmailResp, error) {

	//图片验证
	//if !verify.Instance().Verify(in.Key, in.Value) {
	//	return nil, errors.Wrapf(xerr.NewErrCode(xerr.VERIFY_ERROR), "verify:%s", in.Email)
	//}

	value := tool.Krand(6, tool.KC_RAND_KIND_ALL)

	//将验证码放入Redis 并设置过期时间10分钟
	err := l.svcCtx.RedisClient.Setex(in.Email, value, 600)
	if err != nil {
		return nil, xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, xerr.MapErrMsg(xerr.SERVER_COMMON_ERROR))
	}

	//发邮件  之后将邮箱和授权码放外面
	err = verify.SendMail("chinaskillproject@163.com", in.Email, "验证码", "验证码:"+value, "JUBMSUDZGUVSHKYF")

	if err != nil {
		return &pb.VerifyEmailResp{}, err
	}

	return &pb.VerifyEmailResp{}, nil
}
