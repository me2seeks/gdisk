package File

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"cloud-disk/common/result"

	"cloud-disk/app/disk/cmd/api/internal/logic/File"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
)

func UserFileListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileListRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := File.NewUserFileListLogic(r.Context(), svcCtx)
		resp, err := l.UserFileList(&req, r.Header.Get("UserIdentity"))
		result.HttpResult(r, w, resp, err)
	}
}