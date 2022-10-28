package user

import (
	"cloud-disk/common/result"
	"net/http"

	"cloud-disk/app/user/cmd/api/internal/logic/user"
	"cloud-disk/app/user/cmd/api/internal/svc"
	"cloud-disk/app/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterCountReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCountLogic(r.Context(), svcCtx)
		resp, err := l.Count(&req)
		result.HttpResult(r, w, resp, err)
	}
}
