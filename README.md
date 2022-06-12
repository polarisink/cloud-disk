# CloudDisk

> 轻量级网盘,基于go-zero和xorm

## 一、使用的命令
```shell
# 创建API服务
goctl api new core
# 启动服务
go run core.go -f etc/core-api.yml
# 使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero
```
