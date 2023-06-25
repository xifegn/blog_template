package handler

import (
	"blog_template/internal/logic/download"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := download.NewDownloadLogic(r.Context(), svcCtx)
		err := l.Download(&req)
		response.Response(w, nil, err)
	}
}
