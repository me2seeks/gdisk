package svc

import (
	"cloud-disk/app/disk/cmd/rpc/internal/config"
	"cloud-disk/app/disk/model"
	"cloud-disk/app/user/cmd/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	UserRpc user.Usercenter

	FileModel   model.FileModel
	FolderModel model.FileFolderModel
	StoreModel  model.FileStoreModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),

		FileModel:   model.NewFileModel(conn, c.Cache),
		FolderModel: model.NewFileFolderModel(conn, c.Cache),
		StoreModel:  model.NewFileStoreModel(conn, c.Cache),
	}
}
