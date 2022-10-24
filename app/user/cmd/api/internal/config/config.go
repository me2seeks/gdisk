package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Pass string
		Host string
		Type string
	}
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf

	WxMiniConf WxMiniConf

	UserRpcConf zrpc.RpcClientConf
}
