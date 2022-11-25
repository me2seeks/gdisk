package user

import (
	"cloud-disk/common/result"
	"net/http"

	"cloud-disk/app/user/cmd/api/internal/logic/user"
	"cloud-disk/app/user/cmd/api/internal/svc"
	"cloud-disk/app/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateAvatarReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateAvatarLogic(r.Context(), svcCtx)
		resp, err := l.UpdateAvatar(&req)
		result.HttpResult(r, w, resp, err)
	}
}
