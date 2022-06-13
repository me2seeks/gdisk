package fileFolder

import (
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FolderlistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFolderlistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FolderlistLogic {
	return &FolderlistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FolderlistLogic) Folderlist(req *types.ListReq) (resp *types.ListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
