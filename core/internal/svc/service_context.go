package svc

import (
	"cloud-disk/core/internal/config"
	"cloud-disk/core/models"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB *redis.Client
	OSS *oss.Bucket
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.Mysql.DataSource),
		RDB: models.InitRedis(c),
		//todo now i can't define it in ctx but in define.go ,like a static method in java
		//OSS: models.InitOssBucket(c),
	}
}
