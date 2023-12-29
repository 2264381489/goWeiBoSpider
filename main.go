package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"goSpider/config"
	"goSpider/spider"
	"goSpider/svc"
)

var configFile = flag.String("f", "etc/spider.yaml", "the config file")

func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	serviceContext := svc.NewServiceContext(c)
	ctx := context.Background()
	newSpider := spider.NewSpider(ctx, serviceContext)
	newSpider.Run()
}
