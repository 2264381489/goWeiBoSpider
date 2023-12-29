package svc

import (
	"goSpider/config"
	"goSpider/model"
)

type ServiceContext struct {
	Config config.Config

	MblogModel model.MblogModel
	UserModel  model.UserModel
	PicModel   model.PicModel
	VideoModel model.VideoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		MblogModel: model.NewMblogModel(c.Mongo.Url, c.Mongo.DB, model.CollectionMblog),
		UserModel:  model.NewUserModel(c.Mongo.Url, c.Mongo.DB, model.CollectionUser),
		PicModel:   model.NewPicModel(c.Mongo.Url, c.Mongo.DB, model.CollectionPic),
		VideoModel: model.NewVideoModel(c.Mongo.Url, c.Mongo.DB, model.CollectionVideo),
	}
}
