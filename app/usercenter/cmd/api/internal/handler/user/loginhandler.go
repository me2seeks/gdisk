package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"trytry/common/result"

	"trytry/app/usercenter/cmd/api/internal/logic/user"
	"trytry/app/usercenter/cmd/api/internal/svc"
	"trytry/app/usercenter/cmd/api/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result.HttpResult(r, w, resp, err)
	}
}
