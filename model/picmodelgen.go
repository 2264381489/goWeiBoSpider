package model

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type picModel interface {
	Insert(ctx context.Context, data *Pic) error
	FindOne(ctx context.Context, id string) (*Pic, error)
	Update(ctx context.Context, data *Pic) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultPicModel struct {
	conn *mon.Model
}

func newDefaultPicModel(conn *mon.Model) *defaultPicModel {
	return &defaultPicModel{conn: conn}
}

func (m *defaultPicModel) Insert(ctx context.Context, data *Pic) error {
	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m *defaultPicModel) FindOne(ctx context.Context, id string) (*Pic, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data Pic

	err = m.conn.FindOne(ctx, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPicModel) Update(ctx context.Context, data *Pic) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultPicModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": oid})
	return res, err
}
