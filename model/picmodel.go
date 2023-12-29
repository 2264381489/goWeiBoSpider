package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"time"
)

var _ PicModel = (*customPicModel)(nil)

type (
	// PicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPicModel.
	PicModel interface {
		picModel
		InsertMany(ctx context.Context, data []*Pic) error
	}

	customPicModel struct {
		*defaultPicModel
	}
)

// NewPicModel returns a model for the mongo.
func NewPicModel(url, db, collection string) PicModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customPicModel{
		defaultPicModel: newDefaultPicModel(conn),
	}
}

func (m *defaultPicModel) InsertMany(ctx context.Context, data []*Pic) error {
	docs := make([]interface{}, 0, len(data))
	for _, pic := range data {

		pic.CreateAt = time.Now()
		pic.UpdateAt = time.Now()
		docs = append(docs, pic)
	}
	if _, err := m.conn.InsertMany(ctx, docs); err != nil {
		logx.Errorf("pics insert Many err :  %v", err)
		return err
	}
	return nil
}
