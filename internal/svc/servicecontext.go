package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"goSpider/config"
)

type ServiceContext struct {
	Config   config.Config
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		BizRedis: redis.New(c.RedisConf.Host),
	}
}
