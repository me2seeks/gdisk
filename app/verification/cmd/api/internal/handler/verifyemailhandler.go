package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"cloud-disk/common/result"

	"cloud-disk/app/verification/cmd/api/internal/logic"
	"cloud-disk/app/verification/cmd/api/internal/svc"
	"cloud-disk/app/verification/cmd/api/internal/types"
)

func verifyemailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyEmailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewVerifyemailLogic(r.Context(), svcCtx)
		resp, err := l.Verifyemail(&req)
		result.HttpResult(r, w, resp, err)
	}
}
