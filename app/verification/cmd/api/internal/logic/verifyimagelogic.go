package logic

import (
	"cloud-disk/app/verification/cmd/rpc/verification"
	"context"
	"github.com/jinzhu/copier"

	"cloud-disk/app/verification/cmd/api/internal/svc"
	"cloud-disk/app/verification/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyimageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyimageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyimageLogic {
	return &VerifyimageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyimageLogic) Verifyimage() (resp *types.VerifyImageResp, err error) {

	verifyimage, _ := l.svcCtx.VerificationRpc.VerifyImage(l.ctx, &verification.VerifyImageReq{})
	var verifyImageResp types.VerifyImageResp
	_ = copier.Copy(&verifyImageResp, verifyimage)

	return &verifyImageResp, nil
}
