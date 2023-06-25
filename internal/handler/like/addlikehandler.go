package handler

import (
	"blog_template/internal/logic/like"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func AddLikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddLikeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := like.NewAddLikeLogic(r.Context(), svcCtx)
		resp, err := l.AddLike(&req)
		response.Response(w, resp, err)
	}
}
