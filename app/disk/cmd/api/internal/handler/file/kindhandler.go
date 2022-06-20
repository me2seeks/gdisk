package file

import (
	"net/http"

	"trytry/common/result"

	"cloud-disk/app/disk/cmd/api/internal/logic/file"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
)

func KindHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.KindReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := file.NewKindLogic(r.Context(), svcCtx)
		resp, err := l.Kind(&req)
		result.HttpResult(r, w, resp, err)
	}
}
