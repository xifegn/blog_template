package handler

import (
	"blog_template/internal/logic/follow"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveFollowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveFollowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := follow.NewRemoveFollowLogic(r.Context(), svcCtx)
		resp, err := l.RemoveFollow(&req)
		response.Response(w, resp, err)
	}
}
