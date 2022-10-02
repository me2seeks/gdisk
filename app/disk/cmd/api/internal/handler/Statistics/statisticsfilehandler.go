package Statistics

import (
	"net/http"

	"trytry/common/result"

	"cloud-disk/app/disk/cmd/api/internal/logic/Statistics"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
)

func StatisticsFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StatisticsFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := Statistics.NewStatisticsFileLogic(r.Context(), svcCtx)
		resp, err := l.StatisticsFile(&req)
		result.HttpResult(r, w, resp, err)
	}
}
