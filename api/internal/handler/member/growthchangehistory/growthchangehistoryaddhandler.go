package handler

import (
	"net/http"

	"zdmin/api/internal/logic/member/growthchangehistory"
	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GrowthChangeHistoryAddHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GrowthChangeHistoryAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGrowthChangeHistoryAddLogic(r.Context(), ctx)
		resp, err := l.GrowthChangeHistoryAdd(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
