package handler

import (
	"net/http"

	"zdmin/api/internal/logic/member/integrationchangehistory"
	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func IntegrationChangeHistoryListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IntegrationChangeHistoryListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewIntegrationChangeHistoryListLogic(r.Context(), ctx)
		resp, err := l.IntegrationChangeHistoryList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
