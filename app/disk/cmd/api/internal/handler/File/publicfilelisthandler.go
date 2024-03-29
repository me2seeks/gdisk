package File

import (
	"cloud-disk/common/result"
	"net/http"

	"cloud-disk/app/disk/cmd/api/internal/logic/File"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublicFileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublicFileListRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := File.NewPublicFileListLogic(r.Context(), svcCtx)
		resp, err := l.PublicFileList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
