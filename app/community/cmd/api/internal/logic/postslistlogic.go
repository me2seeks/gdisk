package logic

import (
	"context"

	"cloud-disk/app/community/cmd/api/internal/svc"
	"cloud-disk/app/community/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsListLogic {
	return &PostsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsListLogic) PostsList(req *types.PostsListRequest) (resp *types.PostsListReply, err error) {
	postsList := make([]*types.PostsListItem, 0)
	resp = new(types.PostsListReply)

	err = l.svcCtx.Engine.
		Table("posts_basic").
		Select("posts_basic.identity, posts_basic.title, posts_basic.tags, user.name as owner, user.avatar, " +
			"posts_basic.content, posts_basic.click_num, posts_basic.mention, " +
			"posts_basic.cover, posts_basic.updated_at, " +
			"(SELECT count(posts_comment_basic.identity) from posts_comment_basic where posts_comment_basic.posts_identity = posts_basic.identity and posts_comment_basic.deleted_at IS NULL) as reply_num").
		Joins("left join user on posts_basic.user_identity = user.identity").
		Where("posts_basic.deleted_at IS NULL").
		Order("posts_basic.updated_at desc").
		Find(&postsList).Error

	if err != nil {
		return
	}

	resp.List = postsList

	return resp, nil
}
