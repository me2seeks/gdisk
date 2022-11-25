package logic

import (
	"context"

	"cloud-disk/app/community/cmd/api/internal/svc"
	"cloud-disk/app/community/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsDetailLogic {
	return &PostsDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsDetailLogic) PostsDetail(req *types.PostsDetailRequest) (resp *types.PostsDetailReply, err error) {
	resp = new(types.PostsDetailReply)

	err = l.svcCtx.Engine.
		Table("posts_basic").
		Where("identity = ?", req.Identity).
		Exec("UPDATE posts_basic SET click_num = click_num + 1 where identity = ?", req.Identity).Error
	if err != nil {
		return
	}

	err = l.svcCtx.Engine.
		Table("posts_basic").
		Select("posts_basic.identity, posts_basic.title, posts_basic.tags, user.name as owner, user.avatar, "+
			"posts_basic.content, posts_basic.click_num, posts_basic.mention, "+
			"posts_basic.cover, posts_basic.updated_at, "+
			"(SELECT count(posts_fb.identity) from posts_fb where posts_fb.type = 'collect' and posts_fb.count = 1 and posts_fb.posts_identity = posts_basic.identity and posts_fb.deleted_at IS NULL) as collection, "+
			"(SELECT count(posts_fb.identity) from posts_fb where posts_fb.type = 'dislike' and posts_fb.count = 1 and posts_fb.posts_identity = posts_basic.identity and posts_fb.deleted_at IS NULL) as dislike, "+
			"(SELECT count(posts_fb.identity) from posts_fb where posts_fb.type = 'ilike' and posts_fb.count = 1 and posts_fb.posts_identity = posts_basic.identity and posts_fb.deleted_at IS NULL) as ilike").
		Joins("left join user on posts_basic.user_identity = user.identity").
		Where("posts_basic.identity = ?", req.Identity).
		First(resp).Error
	if err != nil {

		return
	}

	return
}
