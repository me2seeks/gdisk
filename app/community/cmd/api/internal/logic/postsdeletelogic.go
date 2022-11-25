package logic

import (
	"cloud-disk/common/ctxdata"
	"context"

	"cloud-disk/app/community/cmd/api/internal/svc"
	"cloud-disk/app/community/cmd/api/internal/types"
	"cloud-disk/app/community/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostsDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostsDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostsDeleteLogic {
	return &PostsDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostsDeleteLogic) PostsDelete(req *types.PostsDeleteRequest) (resp *types.PostsDeleteReply, err error) {
	userIdentity := ctxdata.GetUidFromCtx(l.ctx)

	err = l.svcCtx.Engine.
		Where("user_identity = ? AND identity = ?", userIdentity, req.Identity).
		Delete(new(model.PostsBasic)).Error

	resp = new(types.PostsDeleteReply)
	if err != nil {

		return
	}
	return
}
