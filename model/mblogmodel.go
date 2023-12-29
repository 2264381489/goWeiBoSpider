package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ MblogModel = (*customMblogModel)(nil)

type (
	// MblogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMblogModel.
	MblogModel interface {
		mblogModel
	}

	customMblogModel struct {
		*defaultMblogModel
	}
)

// NewMblogModel returns a model for the mongo.
func NewMblogModel(url, db, collection string) MblogModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customMblogModel{
		defaultMblogModel: newDefaultMblogModel(conn),
	}
}
