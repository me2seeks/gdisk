package folder

import (
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"cloud-disk/app/disk/model"
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type MoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MoveLogic {
	return &MoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MoveLogic) Move(req *types.MovedReq) (resp *types.MovedResp, err error) {
	uId := ctxdata.GetUidFromCtx(l.ctx)
	switch req.DestPath {
	//重命名
	case req.SrcPath:
		err := l.renameFileOrFolder(req.Id, uId, req.IsFile, req.Name, req.SrcPath)
		if err != nil {
			return nil, err
		}
	//删除
	case "":
		err := l.deleteFileOrFolder(req.Id, req.IsFile, req.SrcPath)
		if err != nil {
			return nil, err
		}
	//移动
	default:
		err := l.moveFileOrFolder(req.Id, req.IsFile, req.SrcPath, req.DestPath)
		if err != nil {
			return nil, err
		}
	}
	return &types.MovedResp{}, nil
}

func (l *MoveLogic) moveFileOrFolder(id int64, isFile bool, srcPath string, destPath string) error {
	if isFile {
		err := l.svcCtx.FileModel.UpdateWithVersion(l.ctx, nil, &model.File{
			Id:       id,
			FilePath: destPath,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "ERROR: Failed to FileModel.UpdateWithVersion err: %v", err)
		}
	} else {

	}
	return nil
}

func (l *MoveLogic) renameFileOrFolder(id, uId int64, isFile bool, newName, srcPath string) error {
	if isFile {
		err := l.svcCtx.FileModel.UpdateWithVersion(l.ctx, nil, &model.File{
			Id:       id,
			FileName: newName,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "ERROR: Failed to FileModel.UpdateWithVersion err: %v", err)
		}
	} else {

	}

	return nil
}

func (l *MoveLogic) deleteFileOrFolder(id int64, isFile bool, srcPath string) error {
	if isFile {
		_, err := l.svcCtx.FileModel.Update(l.ctx, nil, &model.File{
			Id:         id,
			DelState:   globalkey.DelStateYes,
			DeleteTime: time.Now(),
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: Failed to FileModel.Update err: %v", err)
		}
		//	延迟删除对象存储上的文件

	} else {

	}
	return nil
}
