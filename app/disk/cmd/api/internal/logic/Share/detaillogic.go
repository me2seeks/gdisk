package Share

import (
	"cloud-disk/app/user/cmd/rpc/user"
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (*types.ShareBasicDetailReply, error) {
	// logic：其他用户获取分享文件详情
	resp := new(types.ShareBasicDetailReply)
	// 1 更新分享记录的点击次数
	err := l.svcCtx.Engine.
		Table("share_basic").
		Where("identity = ?", req.Identity).
		Exec("UPDATE share_basic SET click_num = click_num + 1 where identity = ?", req.Identity).Error
	if err != nil {
		return resp, err
	}

	// 2 获取资源详细信息
	err = l.svcCtx.Engine.
		Table("share_basic").
		Select("share_basic.identity, share_basic.repository_identity, user_repository.name, repository_pool.ext, "+
			"repository_pool.path, repository_pool.size, share_basic.click_num, share_basic.desc, "+
			"share_basic.user_identity as owner, share_basic.expired_time, share_basic.updated_at").
		Joins("LEFT JOIN repository_pool ON repository_pool.identity = share_basic.repository_identity").
		Joins("LEFT JOIN user_repository ON user_repository.identity = share_basic.user_repository_identity").
		Where("share_basic.identity = ?", req.Identity).
		First(resp).Error
	u, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Identity: resp.Owner,
	})
	if err != nil {
		return resp, err
	}
	resp.Owner = u.User.Name
	resp.Avatar = u.User.Avatar
	return resp, nil
}
