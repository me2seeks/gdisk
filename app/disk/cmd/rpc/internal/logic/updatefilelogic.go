package logic

import (
	"cloud-disk/app/disk/model"
	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"time"

	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrFileNoExistsError = xerr.NewErrMsg("文件夹不存在")

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

func (l *UpdateFileLogic) UpdateFile(in *pb.UpdateFileReq) (*pb.UpdateFileResp, error) {
	fileDetail := new(model.UserRepository)
	//	move
	if in.FileDetail.Pid != "" {
		err := l.svcCtx.Engine.
			Table("user_repository").
			Where("identity = ? AND user_identity = ?", in.FileDetail.Identity, in.FileDetail.Uid).
			First(fileDetail).Error
		if err != nil && err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR 查询file失败，Pid:%s,Uid: %s,err:%v", in.FileDetail.Pid, in.FileDetail.Uid, err)
		}
		if fileDetail.Id == 0 || err == model.ErrNotFound {
			return nil, errors.Wrapf(ErrFileNoExistsError, "Pid: %s Uid:%s", in.FileDetail.Pid, in.FileDetail.Uid)
		}
		err = l.svcCtx.Engine.
			Table("user_repository").
			Where("identity = ? AND deleted_at IS NULL", in.FileDetail.Identity).
			Update("parent_id", in.FileDetail.Pid).Error
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新user_repository失败")
		}
		return nil, nil
	}

	//rename
	if in.FileDetail.Name != "" {
		var cnt int64
		err := l.svcCtx.Engine.
			Table("user_repository").
			Where("name = ?", in.FileDetail.Name).
			Where("parent_id = (select parent_id from user_repository ur where ur.identity = ?)", in.FileDetail.Identity).
			Count(&cnt).Error
		if err != nil && err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR 查询file失败，Pid:%s,Uid: %s,err:%v", in.FileDetail.Pid, in.FileDetail.Uid, err)
		}
		if fileDetail.Id == 0 || err == model.ErrNotFound {
			return nil, errors.Wrapf(ErrFileNoExistsError, "Identity: %s", in.FileDetail.Identity)
		}
		if cnt > 0 {
			return nil, errors.New("文件名已存在")
		}

		err = l.svcCtx.Engine.
			Table("user_repository").
			Where("identity = ? AND user_identity = ?", in.FileDetail.Identity, in.FileDetail.Uid).
			Update("name", in.FileDetail.Name).Error
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新user_repository失败")
		}
		return nil, nil

	}

	//delete
	//TODO 定时任务
	if in.FileDetail.DelState != 0 {
		err := l.svcCtx.Engine.
			Table("user_repository").
			Where("identity = ? AND user_identity = ?", in.FileDetail.Identity, in.FileDetail.Uid).
			Update("del_state", globalkey.DelStateYes, "deleted_at", time.Now()).Error
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新user_repository失败")
		}
		return nil, nil
	}

	return &pb.UpdateFileResp{}, nil
}
