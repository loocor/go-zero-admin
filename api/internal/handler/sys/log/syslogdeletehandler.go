package handler

import (
	"net/http"

	"zdmin/api/internal/logic/sys/log"
	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func SysLogDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SysLogDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSysLogDeleteLogic(r.Context(), ctx)
		resp, err := l.SysLogDelete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
