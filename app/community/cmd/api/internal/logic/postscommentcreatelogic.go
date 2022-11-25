package logic

import (
	"cloud-disk/app/community/model"
	"cloud-disk/common/ctxdata"
	"context"

	"cloud-disk/app/community/cmd/api/internal/svc"
	"cloud-disk/app/community/cmd/api/internal/types"
	"cloud-disk/common/uuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsCommentCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsCommentCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsCommentCreateLogic {
	return &PostsCommentCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsCommentCreateLogic) PostsCommentCreate(req *types.PostsCommentCreateRequest) (resp *types.PostsCommentCreateReply, err error) {
	userIdentity := ctxdata.GetUidFromCtx(l.ctx)

	newPostsComment := &model.PostsCommentBasic{
		Identity:      uuid.UUID(),
		UserIdentity:  userIdentity,
		PostsIdentity: req.PostsIdentity,
		ReplyIdentity: req.ReplyIdentity,
		ReplyName:     req.ReplyName,
		Content:       req.Content,
		Mention:       req.Mention,
		Like:          0,
		Dislike:       0,
		Read:          0,
	}

	resp = new(types.PostsCommentCreateReply)

	err = l.svcCtx.Engine.
		Select("identity", "user_identity", "posts_identity", "reply_identity", "reply_name", "content", "mention", "like", "dislike", "created_at", "updated_at").
		Create(newPostsComment).Error

	return
}
