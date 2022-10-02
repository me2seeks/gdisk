package Statistics

import (
	"net/http"

	"trytry/common/result"

	"cloud-disk/app/disk/cmd/api/internal/logic/Statistics"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
)

func StatisticsShareHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StatisticsShareRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := Statistics.NewStatisticsShareLogic(r.Context(), svcCtx)
		resp, err := l.StatisticsShare(&req)
		result.HttpResult(r, w, resp, err)
	}
}
