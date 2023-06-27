package handler

import (
	"blog_template/internal/logic/follow"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddFollowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddFollowReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := follow.NewAddFollowLogic(r.Context(), svcCtx)
		resp, err := l.AddFollow(&req)
		response.Response(w, resp, err)
	}
}
