package svc

import (
	"cloud-disk/app/verification/cmd/api/internal/config"
	"cloud-disk/app/verification/cmd/rpc/verification"
	"github.com/zeromicro/go-zero/zrpc"
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
