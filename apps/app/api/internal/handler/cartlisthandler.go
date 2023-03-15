package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mylebron/apps/app/api/internal/logic"
	"mylebron/apps/app/api/internal/svc"
	"mylebron/apps/app/api/internal/types"
)

func CartListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCartListLogic(r.Context(), svcCtx)
		resp, err := l.CartList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
