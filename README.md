# CloudDisk

> 轻量级网盘,基于go-zero和xorm

## 一、使用的命令
```shell
# 创建API服务
goctl api new core
# 启动服务
go run core.go -f etc/core-api.yaml
# 使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero
```

## 文档
[阿里云OSS-Golang文档](https://github.com/aliyun/aliyun-oss-go-sdk/blob/master/README-CN.md)
