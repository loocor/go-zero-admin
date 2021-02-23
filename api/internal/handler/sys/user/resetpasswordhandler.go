package handler

import (
	"net/http"

	"zdmin/api/internal/logic/sys/user"
	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ReSetPasswordHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReSetPasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewReSetPasswordLogic(r.Context(), ctx)
		resp, err := l.ReSetPassword(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
