package download

import (
	"net/http"

	"blog_template/internal/logic/download"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := download.NewDownloadLogic(r.Context(), svcCtx, w)
		err := l.Download(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
