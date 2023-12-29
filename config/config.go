package config

type Config struct {
	Mongo MongoConfig // mongo 相关配置
}

type MongoConfig struct {
	Url string
	DB  string `json:",default=jarvis"`
}
