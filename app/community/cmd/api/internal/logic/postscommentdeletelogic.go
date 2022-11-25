package logic

import (
	"cloud-disk/app/community/model"
	"cloud-disk/common/ctxdata"
	"context"

	"cloud-disk/app/community/cmd/api/internal/svc"
	"cloud-disk/app/community/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type PostsCommentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsCommentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsCommentDeleteLogic {
	return &PostsCommentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsCommentDeleteLogic) PostsCommentDelete(req *types.PostsCommentDeleteRequest) (resp *types.PostsCommentDeleteReply, err error) {
	userIdentity := ctxdata.GetUidFromCtx(l.ctx)

	err = l.svcCtx.Engine.
		Where("user_identity = ? AND identity = ?", userIdentity, req.Identity).
		Delete(new(model.PostsCommentBasic)).Error

	resp = new(types.PostsCommentDeleteReply)
	if err != nil {
		return
	}
	return resp, nil
}
