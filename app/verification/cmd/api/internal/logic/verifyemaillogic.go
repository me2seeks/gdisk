package logic

import (
	"cloud-disk/app/verification/cmd/api/internal/svc"
	"cloud-disk/app/verification/cmd/api/internal/types"
	"cloud-disk/app/verification/cmd/rpc/verification"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyemailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyemailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyemailLogic {
	return &VerifyemailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyemailLogic) Verifyemail(req *types.VerifyEmailReq) (resp *types.VerifyEmailResp, err error) {

	_, err = l.svcCtx.VerificationRpc.VerifyEmail(l.ctx, &verification.VerifyEmailReq{
		Email: req.Email,
	})

	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return resp, nil
}
