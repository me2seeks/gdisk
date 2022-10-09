package model

import (
	"github.com/8treenet/gcache/option"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

func Init(dataSource string) *gorm.DB {
	opt := option.DefaultOption{}
	opt.Expires = 300              //缓存时间，默认60秒。范围 30-900
	opt.Level = option.LevelSearch //缓存级别，默认LevelSearch。LevelDisable:关闭缓存，LevelModel:模型缓存， LevelSearch:查询缓存
	opt.AsyncWrite = true          //异步缓存更新, 默认false。 insert update delete 成功后是否异步更新缓存
	opt.PenetrationSafe = true     //开启防穿透, 默认false。

	engine, err := gorm.Open(mysql.Open("root:"+"000000"+"@tcp(43.143.125.75:33060)/cloud_disk?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		logx.Errorf("gorm New Engine Error:%v", err)
		return nil
	}

	return engine
}

var ErrNotFound = gorm.ErrRecordNotFound
