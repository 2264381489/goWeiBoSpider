package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mblog struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Content        string             `bson:"content,omitempty" json:"content,omitempty"`
	MblogId        string             `bson:"mblogId,omitempty" json:"mblogId,omitempty"`
	UserId         int64              `bson:"userId,omitempty" json:"userId,omitempty"`
	Url            string             `bson:"url,omitempty" json:"url,omitempty"`
	Text           string             `bson:"text" json:"text"`
	AttitudesCount int64              `bson:"attitudesCount" json:"attitudesCount"`
	CommentsCount  int64              `bson:"commentsCount" json:"commentsCount"`
	RepostsCount   int64              `bson:"repostsCount" json:"repostsCount"`
	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
