package folder

import (
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/common/ctxdata"
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListReq) (*types.ListResp, error) {
	uId := ctxdata.GetUidFromCtx(l.ctx)

	var folders []types.Folder
	var files []types.File

	result, err := l.svcCtx.DiskRpc.ListFolders(l.ctx, &pb.ListFolderReq{
		Uid:  uId,
		Path: req.Path,
	})
	if err != nil {
		return nil, err
	}

	if result.Files != nil {
		copier.Copy(&files, result.Files)
	}
	if result.Folders != nil {
		copier.Copy(&folders, result.Folders)
	}

	return &types.ListResp{
		FolderList: folders,
		FileList:   files,
	}, nil
}
