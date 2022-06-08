package logic

import (
	"context"
	"github.com/pkg/errors"
	"trytry/app/verification/cmd/api/internal/svc"
	"trytry/app/verification/cmd/api/internal/types"
	"trytry/app/verification/cmd/rpc/verification"

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
		Key:   req.Key,
		Value: req.Value,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	return nil, nil
}
