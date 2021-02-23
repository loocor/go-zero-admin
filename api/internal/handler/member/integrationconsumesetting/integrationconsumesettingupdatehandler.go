package handler

import (
	"net/http"

	"zdmin/api/internal/logic/member/integrationconsumesetting"
	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func IntegrationConsumeSettingUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IntegrationConsumeSettingUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewIntegrationConsumeSettingUpdateLogic(r.Context(), ctx)
		resp, err := l.IntegrationConsumeSettingUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
