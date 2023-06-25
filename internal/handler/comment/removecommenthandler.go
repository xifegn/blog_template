package handler

import (
	"blog_template/internal/logic/comment"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func RemoveCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := comment.NewRemoveCommentLogic(r.Context(), svcCtx)
		resp, err := l.RemoveComment(&req)
		response.Response(w, resp, err)
	}
}
