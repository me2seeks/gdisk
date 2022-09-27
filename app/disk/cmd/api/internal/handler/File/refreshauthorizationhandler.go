package File

import (
	"net/http"

	"trytry/common/result"

	"cloud-disk/app/disk/cmd/api/internal/logic/File"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
)

func RefreshAuthorizationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefreshAuthorizationRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := File.NewRefreshAuthorizationLogic(r.Context(), svcCtx)
		resp, err := l.RefreshAuthorization(&req)
		result.HttpResult(r, w, resp, err)
	}
}
