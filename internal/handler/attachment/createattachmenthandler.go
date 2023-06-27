package handler

import (
	"blog_template/internal/logic/attachment"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateAttachmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateAttachmentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := attachment.NewCreateAttachmentLogic(r.Context(), svcCtx)
		resp, err := l.CreateAttachment(&req)
		response.Response(w, resp, err)
	}
}
