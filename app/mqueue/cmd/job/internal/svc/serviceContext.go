package svc

import (
	"cloud-disk/app/disk/cmd/rpc/disk"
	"cloud-disk/app/mqueue/cmd/job/internal/config"
	//"cloud-disk/app/user/cmd/rpc/user"
	"github.com/hibiken/asynq"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	MiniProgram *miniprogram.MiniProgram

	DiskRpc disk.Disk
	//UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AsynqServer: newAsynqServer(c),
		//MiniProgram: newMiniprogramClient(c),
		DiskRpc: disk.NewDisk(zrpc.MustNewClient(c.DiskRpcConf)),
	}
}
