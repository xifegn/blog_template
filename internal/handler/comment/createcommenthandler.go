package handler

import (
	"blog_template/internal/logic/comment"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := comment.NewCreateCommentLogic(r.Context(), svcCtx)
		resp, err := l.CreateComment(&req)
		response.Response(w, resp, err)
	}
}
