package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	Redis struct {
		Addr string
		Password string
	}
	Oss struct{
		Endpoint string
		AccessKeyId string
		AccessKeySecret string
		Bucket string
	}
}
