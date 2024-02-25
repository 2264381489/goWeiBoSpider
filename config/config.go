package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	Mongo     MongoConfig // mongo 相关配置
	RedisConf redis.RedisConf
	rest.RestConf
}

type (
	MongoConfig struct {
		Url string
		DB  string `json:",default=spider"`
	}

	//// A RedisConf is a redis config.
	//RedisConf struct {
	//	Host     string
	//	Type     string `json:",default=node,options=node|cluster"`
	//	Pass     string `json:",optional"`
	//	Tls      bool   `json:",optional"`
	//	NonBlock bool   `json:",default=true"`
	//	// PingTimeout is the timeout for ping redis.
	//	PingTimeout time.Duration `json:",default=1s"`
	//}
)
