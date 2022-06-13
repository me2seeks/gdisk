package user

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"cloud-disk/common/result"

	"cloud-disk/app/usercenter/cmd/api/internal/logic/user"
	"cloud-disk/app/usercenter/cmd/api/internal/svc"
	"cloud-disk/app/usercenter/cmd/api/internal/types"
)

func WxAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WXMiniAuthReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewWxAuthLogic(r.Context(), svcCtx)
		resp, err := l.WxAuth(&req)
		result.HttpResult(r, w, resp, err)
	}
}
