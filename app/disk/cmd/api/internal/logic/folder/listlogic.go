package folder

import (
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"cloud-disk/app/disk/model"
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/xerr"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
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
	var files []types.FolderFile

	err := mr.Finish(func() error {
		builder := l.svcCtx.FolderModel.RowBuilder().Where(squirrel.Eq{
			"user_id":     uId,
			"folder_path": req.Path,
		})
		pointerFolders, err := l.svcCtx.FolderModel.FindAll(l.ctx, builder, "desc")
		if err != nil || err != model.ErrNotFound {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR failed to FolderModel.FindAll uid: %d path: %s err: %v", uId, req.Path, err)
		}
		for i, folder := range pointerFolders {
			copier.Copy(&folders[i], folder)
			folders[i].UpdateTime = folder.UpdateTime.Unix()
		}
		return nil
	}, func() error {
		builder := l.svcCtx.FileModel.RowBuilder().Where(squirrel.Eq{
			"user_id":   uId,
			"file_path": req.Path,
		})
		pointerFiles, err := l.svcCtx.FileModel.FindAll(l.ctx, builder, "desc")
		if err != nil || err != model.ErrNotFound {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR failed to FileModel.FindAll uid: %d path: %s err: %v", uId, req.Path, err)
		}

		for i, file := range pointerFiles {
			copier.Copy(&files[i], file)
			files[i].CreateTime = file.CreateTime.Unix()
			files[i].UpdateTime = file.UpdateTime.Unix()
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &types.ListResp{
		FolderList: folders,
		FileList:   files,
	}, nil
}
