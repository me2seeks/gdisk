package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"trytry/app/verification/cmd/rpc/verification"

	"trytry/app/verification/cmd/api/internal/svc"
	"trytry/app/verification/cmd/api/internal/types"

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

	verifyimage, err := l.svcCtx.VerificationRpc.VerifyImage(l.ctx, &verification.VerifyImageReq{})
	var verifyImageResp types.VerifyImageResp
	_ = copier.Copy(&verifyImageResp, verifyimage)

	return &verifyImageResp, nil
}
