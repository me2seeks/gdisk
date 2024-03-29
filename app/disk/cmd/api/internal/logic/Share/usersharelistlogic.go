package Share

import (
	"cloud-disk/common/ctxdata"
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserShareListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserShareListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserShareListLogic {
	return &UserShareListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserShareListLogic) UserShareList(_ *types.UserShareListRequest) (resp *types.UserShareListReply, err error) {
	u := ctxdata.GetUidFromCtx(l.ctx)
	shareFile := make([]*types.ShareBasicDetailReply, 0)
	resp = new(types.UserShareListReply)

	err = l.svcCtx.Engine.
		Table("share_basic").
		Select("share_basic.identity, share_basic.repository_identity, user_repository.name, repository_pool.ext, "+
			"repository_pool.path, repository_pool.size, share_basic.click_num, share_basic.desc, "+
			"user.name as owner, user.avatar, share_basic.expired_time, share_basic.updated_at").
		Joins("LEFT JOIN repository_pool ON repository_pool.identity = share_basic.repository_identity").
		Joins("LEFT JOIN user_repository ON user_repository.identity = share_basic.user_repository_identity").
		Joins("left join user on share_basic.user_identity = user.identity").
		Where("share_basic.user_identity = ?", u).
		Where("share_basic.deleted_at IS NULL").
		Order("share_basic.click_num desc").
		Find(&shareFile).Error

	if err != nil {
		return
	}

	resp.List = shareFile
	return

}
