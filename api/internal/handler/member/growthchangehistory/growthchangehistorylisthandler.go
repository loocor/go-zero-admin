package handler

import (
	"net/http"

	"zdmin/api/internal/logic/member/growthchangehistory"
	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GrowthChangeHistoryListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GrowthChangeHistoryListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGrowthChangeHistoryListLogic(r.Context(), ctx)
		resp, err := l.GrowthChangeHistoryList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
