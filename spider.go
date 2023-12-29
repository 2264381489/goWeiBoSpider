package main

import (
	"flag"
	"fmt"
	"goSpider/internal/config"
	"goSpider/internal/handler"
	"goSpider/internal/svc"
	"goSpider/swagger/spider/doc"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/spider-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/swagger/*any",
		Handler: doc.WrapHandler(),
	})
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
