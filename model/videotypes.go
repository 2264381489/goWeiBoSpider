package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Video struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	VideoUrl    string    `bson:"videoUrl" json:"videoUrl"`
	VideoHDUrl  string    `bson:"videoHDUrl,omitempty" json:"videoHDUrl,omitempty"`
	VideoTitle  string    `bson:"videoTitle" json:"videoTitle"`
	Video720URL string    `bson:"video720URL,omitempty" json:"video720URL,omitempty"`
	Duration    float64   `bson:"duration" json:"duration"`
	Title       string    `bson:"title" json:"title"`
	UpdateAt    time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
