package handler

import (
	"net/http"

	"go-zero-admin/front/internal/logic/member/member"
	"go-zero-admin/front/internal/svc"
	"go-zero-admin/front/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func MemberLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MemberLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMemberLoginLogic(r.Context(), ctx)
		resp, err := l.MemberLogin(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
