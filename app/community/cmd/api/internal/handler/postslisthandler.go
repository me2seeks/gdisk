package handler

import (
	"cloud-disk/common/result"
	"net/http"

	"cloud-disk/app/community/cmd/api/internal/logic"
	"cloud-disk/app/community/cmd/api/internal/svc"
	"cloud-disk/app/community/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PostsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostsListRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewPostsListLogic(r.Context(), svcCtx)
		resp, err := l.PostsList(&req)
		result.HttpResult(r, w, resp, err)

	}
}
