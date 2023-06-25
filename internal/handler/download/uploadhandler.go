package handler

import (
	"blog_template/internal/logic/download"
	"blog_template/internal/svc"
	"blog_template/response"
	"net/http"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := download.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload()
		response.Response(w, resp, err)
	}
}
