package models

import (
	"cloud-disk/core/internal/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

func Init(dataSource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	return engine
}

func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: c.Redis.Password, // no password set
		DB:       0,                // use default DB
	})
}

// InitOssBucket 初始化OssClient
func InitOssBucket(o config.Config) *oss.Bucket {
	client, err := oss.New(o.Oss.Endpoint, o.Oss.AccessKeyId, o.Oss.AccessKeySecret)
	if err != nil {
		log.Printf("Oss Client Init Error: %v", err)
	}
	bucket, err := client.Bucket(o.Oss.Bucket)
	if err!=nil {
		log.Printf("Oss Bucket Init Error: %v", err)
	}
	return bucket
}
