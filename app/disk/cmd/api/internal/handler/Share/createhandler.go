package Share

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"cloud-disk/common/result"

	"cloud-disk/app/disk/cmd/api/internal/logic/Share"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
)

func CreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShareBasicCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := Share.NewShareBasicCreateLogic(r.Context(), svcCtx)
		resp, err := l.ShareBasicCreate(&req)
		result.HttpResult(r, w, resp, err)
	}
}
