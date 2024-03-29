// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"goSpider/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/spider/create",
				Handler: SpiderHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/driver/aliyuncallback",
				Handler: AliyunCallbackHandler(serverCtx),
			},
		},
	)
}
