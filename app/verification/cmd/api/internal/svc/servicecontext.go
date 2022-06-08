package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"trytry/app/verification/cmd/api/internal/config"
	"trytry/app/verification/cmd/rpc/verification"
)

type ServiceContext struct {
	Config          config.Config
	VerificationRpc verification.Verification
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		VerificationRpc: verification.NewVerification(zrpc.MustNewClient(c.VerificationRpc)),
	}
}
