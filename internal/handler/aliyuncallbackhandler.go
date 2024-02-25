package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goSpider/internal/logic"
	"goSpider/internal/svc"
	"goSpider/internal/types"
)

func AliyunCallbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AliYunCallbackReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAliyunCallbackLogic(r.Context(), svcCtx)
		resp, err := l.AliyunCallback(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
