package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

// NewVideoModel returns a model for the mongo.
func NewVideoModel(url, db, collection string) VideoModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customVideoModel{
		defaultVideoModel: newDefaultVideoModel(conn),
	}
}
