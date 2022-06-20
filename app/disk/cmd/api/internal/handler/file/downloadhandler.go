package file

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"cloud-disk/common/result"

	"cloud-disk/app/disk/cmd/api/internal/logic/file"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadCertificateReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := file.NewDownloadLogic(r.Context(), svcCtx)
		resp, err := l.Download(&req)
		result.HttpResult(r, w, resp, err)
	}
}
