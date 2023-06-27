package handler

import (
	"blog_template/internal/logic/user"
	"blog_template/internal/svc"
	"blog_template/internal/types"
	"blog_template/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ModifyAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyAvatarReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewModifyAvatarLogic(r.Context(), svcCtx)
		resp, err := l.ModifyAvatar(&req)
		response.Response(w, resp, err)
	}
}
