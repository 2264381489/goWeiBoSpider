package main

type ContainerType string

const (
	MainTheme ContainerType = "main_theme"
	WeiBo     ContainerType = "weibo"
)

type CardType int

const (
	Commend CardType = 156
	Mblog   CardType = 9
)

func (c CardType) ToInt() int {
	return int(c)
}

type MediaType string

const (
	VideoType   MediaType = "video"
	SearchTopic MediaType = "search_topic"
	Topic       MediaType = "topic"
	BigPic      MediaType = "bigPic"
)

func (m MediaType) String() string {
	return string(m)
}

const userWeiBoList = "https://m.weibo.cn/api/container/getIndex?type=uid&value=%d&containerid=%s" // param1 userId param2 containerId

const (
	MainThemeCId = 230283
	WeiBoCId     = 107603
)
