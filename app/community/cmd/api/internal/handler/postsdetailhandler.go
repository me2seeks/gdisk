package handler

import (
	"cloud-disk/common/result"
	"net/http"

	"cloud-disk/app/community/cmd/api/internal/logic"
	"cloud-disk/app/community/cmd/api/internal/svc"
	"cloud-disk/app/community/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PostsDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostsDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPostsDetailLogic(r.Context(), svcCtx)
		resp, err := l.PostsDetail(&req)
		result.HttpResult(r, w, resp, err)

	}
}
