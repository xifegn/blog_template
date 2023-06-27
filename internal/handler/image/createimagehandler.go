package handler

import (
	"blog_template/internal/logic/image"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateImageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := image.NewCreateImageLogic(r.Context(), svcCtx)
		resp, err := l.CreateImage(&req)
		response.Response(w, resp, err)
	}
}
