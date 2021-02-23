package handler

import (
	"net/http"

	"zdmin/api/internal/logic/member/level"
	"zdmin/api/internal/svc"
	"zdmin/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func MemberLevelDeleteHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MemberLevelDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMemberLevelDeleteLogic(r.Context(), ctx)
		resp, err := l.MemberLevelDelete(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
