package logic

import (
	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"time"

	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFolderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFolderLogic {
	return &UpdateFolderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改folder
func (l *UpdateFolderLogic) UpdateFolder(in *pb.UpdateFolderReq) (*pb.UpdateFolderResp, error) {
	oldFolder, err := l.svcCtx.FolderModel.FindOne(l.ctx, in.Folder.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: Failed to FolderModel.FindOne err: %v", err)
	}

	if in.Folder.FolderName != "" {
		oldFolder.FolderName = in.Folder.FolderName
	}
	if in.Folder.ParentFolderId > 0 {
		oldFolder.ParentFolderId = in.Folder.ParentFolderId
	}
	if in.Folder.DeleteState != 0 {
		oldFolder.DelState = globalkey.DelStateYes
		oldFolder.DeleteTime = time.Now()
	}

	err = l.svcCtx.FolderModel.UpdateWithVersion(l.ctx, nil, oldFolder)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: Failed to FolderModel.UpdateWithVersion err: %v", err)
	}

	return &pb.UpdateFolderResp{}, nil
}
