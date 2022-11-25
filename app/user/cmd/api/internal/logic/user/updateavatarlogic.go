package logic

import (
	"cloud-disk/common/ctxdata"
	"context"

	"cloud-disk/app/user/cmd/api/internal/svc"
	"cloud-disk/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAvatarLogic {
	return &UpdateAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAvatarLogic) UpdateAvatar(req *types.UpdateAvatarReq) (resp *types.UpdateAvatarResp, err error) {
	identity := ctxdata.GetUidFromCtx(l.ctx)

	u, err := l.svcCtx.UserModel.FindOneByIdentity(l.ctx, identity)
	u.Avatar = req.Avatar
	_, err = l.svcCtx.UserModel.Update(l.ctx, nil, u)
	return
}
