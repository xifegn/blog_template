package handler

import (
	"blog_template/internal/logic/like"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RemoveLikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveLikedReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := like.NewRemoveLikeLogic(r.Context(), svcCtx)
		resp, err := l.RemoveLike(&req)
		response.Response(w, resp, err)
	}
}
