package svc

import (
	"cloud-disk/app/order/cmd/mq/internal/config"
	"cloud-disk/app/user/cmd/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc      order.Order
	UserCenterRpc user.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		UserCenterRpc: user.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
