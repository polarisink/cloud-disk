package main

import (
	"cloud-disk/core/internal/common/errorx"
	"cloud-disk/core/internal/config"
	"cloud-disk/core/internal/handler"
	"cloud-disk/core/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/zero-contrib/logx/zapx"
	"net/http"
)

var configFile = flag.String("f", "./core/etc/core-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	initZap()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

//  初始化zapx日志
func initZap() {
	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)
	logx.Infof("Init logx...")
}
