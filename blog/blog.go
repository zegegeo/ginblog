package main

import (
	"flag"
	"fmt"
	"net/http"
	"errors"

	"goblog/blog/internal/biz"
	"goblog/blog/internal/config"
	"goblog/blog/internal/handler"
	"goblog/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/blog-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(
		c.RestConf,
		rest.WithCors("http://localhost:5173"),
	)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(func(err error)(int, any){
		var e *biz.Error
		switch {
			case errors.As(err, &e):
				return http.StatusOK, biz.Fail(e)
			default:
				return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
