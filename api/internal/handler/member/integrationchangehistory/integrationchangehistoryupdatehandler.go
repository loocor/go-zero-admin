package handler

import (
	"net/http"

	"zdmin/api/internal/logic/member/integrationchangehistory"
	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func IntegrationChangeHistoryUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IntegrationChangeHistoryUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewIntegrationChangeHistoryUpdateLogic(r.Context(), ctx)
		resp, err := l.IntegrationChangeHistoryUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
