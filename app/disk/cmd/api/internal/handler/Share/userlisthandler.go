package Share

import (
	"cloud-disk/common/result"
	"net/http"

	"cloud-disk/app/disk/cmd/api/internal/logic/Share"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserShareListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserShareListRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := Share.NewUserShareListLogic(r.Context(), svcCtx)
		resp, err := l.UserShareList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
