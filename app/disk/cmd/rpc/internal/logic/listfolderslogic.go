package logic

import (
	"cloud-disk/app/disk/model"
	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/mr"

	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFoldersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListFoldersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFoldersLogic {
	return &ListFoldersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取路径下的文件和文件夹
func (l *ListFoldersLogic) ListFolders(in *pb.ListFolderReq) (*pb.ListFolderResp, error) {

	var folders []*pb.FolderDetail
	var files []*pb.FileDetail

	err := mr.Finish(func() error {
		builder := l.svcCtx.FolderModel.RowBuilder().Where(squirrel.Eq{
			"user_id":          in.Uid,
			"parent_folder_id": in.Pid,
			"del_state":        globalkey.DelStateNo,
		})
		pointerFolders, err := l.svcCtx.FolderModel.FindAll(l.ctx, builder, "desc")
		if err != nil || err != model.ErrNotFound {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR failed to FolderModel.FindAll uid: %d pid: %d err: %v", in.Uid, in.Pid, err)
		}
		for i, folder := range pointerFolders {
			copier.Copy(&folders[i], folder)
			folders[i].UpdateTime = folder.UpdateTime.Unix()
		}
		return nil
	}, func() error {
		builder := l.svcCtx.FileModel.RowBuilder().Where(squirrel.Eq{
			"user_id":          in.Uid,
			"parent_folder_id": in.Pid,
			"del_state":        globalkey.DelStateNo,
		})
		pointerFiles, err := l.svcCtx.FileModel.FindAll(l.ctx, builder, "desc")
		if err != nil || err != model.ErrNotFound {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR failed to FileModel.FindAll uid: %d pid: %d err: %v", in.Uid, in.Pid, err)
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

	return &pb.ListFolderResp{
		Files:   files,
		Folders: folders,
	}, nil
}
