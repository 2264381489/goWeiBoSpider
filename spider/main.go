package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"goSpider/config"
	"goSpider/svc"
)

var configFile = flag.String("f", "etc/spider.yaml", "the config file")
var uid = flag.Int("u", 0, "the userId ")

func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	serviceContext := svc.NewServiceContext(c)
	ctx := context.Background()
	newSpider := NewSpider(ctx, serviceContext)
	if *uid == 0 {
		panic("未传入uid")
	}
	newSpider.Run(*uid)
}
