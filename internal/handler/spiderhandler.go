package handler

import (
	"goSpider/internal/logic"
	"goSpider/internal/svc"
	"goSpider/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SpiderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSpiderLogic(r.Context(), svcCtx)
		resp, err := l.Spider(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
