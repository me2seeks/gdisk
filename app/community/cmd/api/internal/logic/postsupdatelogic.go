package logic

import (
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/globalkey"
	"context"

	"cloud-disk/app/community/cmd/api/internal/svc"
	"cloud-disk/app/community/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsUpdateLogic {
	return &PostsUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsUpdateLogic) PostsUpdate(req *types.PostsUpdateRequest) (resp *types.PostsUpdateReply, err error) {
	userIdentity := ctxdata.GetUidFromCtx(l.ctx)

	resp = new(types.PostsUpdateReply)
	if req.Title == "" {
		return
	}

	var cntTitle int64
	err = l.svcCtx.Engine.
		Table("posts_basic").
		Where("title = ? AND identity != ? AND user_identity = ? AND del_state = ? ", req.Title, req.Identity, userIdentity, globalkey.DelStateNo).
		Count(&cntTitle).Error
	if err != nil {

		return
	}
	if cntTitle > 0 {

		return
	}

	// 更新
	err = l.svcCtx.Engine.
		Table("posts_basic").
		Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).
		Updates(map[string]interface{}{"title": req.Title, "content": req.Content, "tags": req.Tags, "mention": req.Mention, "cover": req.Cover}).Error

	if err != nil {

		return
	}

	return resp, nil
}
