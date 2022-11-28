package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"cloud-disk/app/community/cmd/api/internal/logic"
	"cloud-disk/app/community/cmd/api/internal/svc"
	"cloud-disk/app/community/cmd/api/internal/types"
	"cloud-disk/common/result"
)

func PostsUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostsUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPostsUpdateLogic(r.Context(), svcCtx)
		resp, err := l.PostsUpdate(&req)
		result.HttpResult(r, w, resp, err)
	}
}
