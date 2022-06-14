package svc

import (
	"cloud-disk/app/mqueue/cmd/job/internal/config"
	"cloud-disk/app/order/cmd/rpc/order"
	"cloud-disk/app/user/cmd/rpc/user"
	"github.com/hibiken/asynq"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	MiniProgram *miniprogram.MiniProgram

	OrderRpc      order.Order
	UsercenterRpc user.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AsynqServer:   newAsynqServer(c),
		MiniProgram:   newMiniprogramClient(c),
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		UsercenterRpc: user.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
