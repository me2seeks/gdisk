package File

import (
	"cloud-disk/common/result"
	"net/http"

	"cloud-disk/app/disk/cmd/api/internal/logic/File"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublicFileSaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublicFileSaveRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := File.NewPublicFileSaveLogic(r.Context(), svcCtx)
		resp, err := l.PublicFileSave(&req)
		result.HttpResult(r, w, resp, err)
	}
}
