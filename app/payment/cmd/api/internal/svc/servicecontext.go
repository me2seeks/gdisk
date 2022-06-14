package svc

import (
	"cloud-disk/app/order/cmd/rpc/order"
	"cloud-disk/app/payment/cmd/api/internal/config"
	"cloud-disk/app/payment/cmd/rpc/payment"
	"cloud-disk/app/user/cmd/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PaymentRpc    payment.Payment
	OrderRpc      order.Order
	UsercenterRpc user.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		PaymentRpc:    payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		UsercenterRpc: user.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
