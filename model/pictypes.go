package model

import (
	"time"
)

type Pic struct {
	ID string `bson:"_id,omitempty" json:"id,omitempty"`
	// TODO: Fill your own fields
	PicId    string    `bson:"picId" json:"picId"`
	Url      string    `bson:"url" json:"url"`
	LargeUrl string    `bson:"largeUrl" json:"largeUrl"`
	MblogId  string    `bson:"mblogId" json:"mblogId"`
	Desc     string    `bson:"desc" json:"desc"`
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
