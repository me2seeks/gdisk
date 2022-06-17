package logic

import (
	"cloud-disk/common/globalkey"
	"cloud-disk/common/tool"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"time"

	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFileLogic {
	return &UpdateFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改file
func (l *UpdateFileLogic) UpdateFile(in *pb.UpdateFileReq) (*pb.UpdateFileResp, error) {
	oldFile, err := l.svcCtx.FileModel.FindOne(l.ctx, in.File.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: Failed to FileModel.FindOne err: %v", err)
	}

	if in.File.FileName != "" {
		oldFile.FileName = in.File.FileName
		oldFile.Postfix = tool.GetSuffix(in.File.FileName)
	}
	if in.File.ParentFolderId > 0 {
		oldFile.ParentFolderId = in.File.ParentFolderId
	}
	if in.File.Size != 0 {
		oldFile.Size = in.File.Size
	}
	if in.File.DeleteState != 0 {
		oldFile.DelState = globalkey.DelStateYes
		oldFile.DeleteTime = time.Now()
	}

	err = l.svcCtx.FileModel.UpdateWithVersion(l.ctx, nil, oldFile)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: Failed to FileModel.UpdateWithVersion err: %v", err)
	}

	return &pb.UpdateFileResp{}, nil
}
