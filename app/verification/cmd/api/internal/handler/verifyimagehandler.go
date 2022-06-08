package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"trytry/common/result"

	"trytry/app/verification/cmd/api/internal/logic"
	"trytry/app/verification/cmd/api/internal/svc"
	"trytry/app/verification/cmd/api/internal/types"
)

func verifyimageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyImageReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewVerifyimageLogic(r.Context(), svcCtx)
		resp, err := l.Verifyimage()
		result.HttpResult(r, w, resp, err)
	}
}
