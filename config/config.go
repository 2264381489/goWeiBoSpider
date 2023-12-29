package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	Mongo MongoConfig // mongo 相关配置
	rest.RestConf
}

type MongoConfig struct {
	Url string
	DB  string `json:",default=jarvis"`
}
