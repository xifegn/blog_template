package handler

import (
	"blog_template/internal/logic/summary"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func CreateSummaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSummaryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := summary.NewCreateSummaryLogic(r.Context(), svcCtx)
		resp, err := l.CreateSummary(&req)
		response.Response(w, resp, err)
	}
}
