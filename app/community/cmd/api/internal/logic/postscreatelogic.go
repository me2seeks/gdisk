package logic

import (
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/globalkey"
	"cloud-disk/common/uuid"
	"context"

	"cloud-disk/app/community/cmd/api/internal/svc"
	"cloud-disk/app/community/cmd/api/internal/types"
	"cloud-disk/app/community/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsCreateLogic {
	return &PostsCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsCreateLogic) PostsCreate(req *types.PostsCreateRequest) (resp *types.PostsCreateReply, err error) {
	userIdentity := ctxdata.GetUidFromCtx(l.ctx)

	newPosts := &model.PostsBasic{
		Identity:     uuid.UUID(),
		UserIdentity: userIdentity,
		Title:        req.Title,
		Tags:         req.Tags,
		Content:      req.Content,
		Mention:      req.Mention,
		Cover:        req.Cover,
		ClickNum:     0,
	}

	resp = new(types.PostsCreateReply)

	var count int64
	err = l.svcCtx.Engine.
		Table("posts_basic").
		Where("title = ? AND user_identity = ? AND del_state = ? ", req.Title, userIdentity, globalkey.DelStateNo).
		Count(&count).Error

	if count > 0 {
		return
	}

	err = l.svcCtx.Engine.
		Select("identity", "user_identity", "title", "tags", "content", "mention", "cover", "click_num", "created_at", "updated_at").
		Create(newPosts).Error

	return resp, nil
}
