package doc

import (
	"github.com/CloverOS/goctl-swag/doc"
)

var Route = RouteImpl{}

type RouteImpl struct {
}

func (r RouteImpl) GetRouteInfos() []doc.RouteInfos {
	return []doc.RouteInfos{
		{
			BasePath:   "/",
			HandlerFun: "SpiderHandler",
			Method:     "post",
			Path:       "/spider/create",
			Public:     false,
			RouteGroup: doc.RouteGroup{GroupName: ""},
			Summary:    "",
		},
	}
}
