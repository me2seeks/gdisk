package logic

import (
	"cloud-disk/app/verification/cmd/rpc/internal/svc"
	"cloud-disk/app/verification/cmd/rpc/pb"
	"cloud-disk/common/oss"
	"cloud-disk/common/verify"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyimageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyimageLogic {
	return &VerifyimageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyimageLogic) VerifyImage(_ *pb.VerifyImageReq) (*pb.VerifyImageResp, error) {

	captchaId := verify.Instance().CreateImage()
	data := verify.Instance().GetImageByte(captchaId)

	//将图片上传到七牛云并获取url
	url, err := oss.UploadToQiNiu(captchaId, data)
	if err != nil {
		return nil, err
	}

	return &pb.VerifyImageResp{
		ImageUrl: url,
	}, nil
}
