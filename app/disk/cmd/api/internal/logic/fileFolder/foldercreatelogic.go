package fileFolder

import (
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FoldercreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFoldercreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FoldercreateLogic {
	return &FoldercreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FoldercreateLogic) Foldercreate(req *types.CreateFolderReq) (resp *types.CreateFolderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
