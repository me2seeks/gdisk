package File

import (
	"net/http"

	"trytry/common/result"

	"cloud-disk/app/disk/cmd/api/internal/logic/File"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
)

func FileDownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileDownloadRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := File.NewFileDownloadLogic(r.Context(), svcCtx)
		resp, err := l.FileDownload(&req)
		result.HttpResult(r, w, resp, err)
	}
}
