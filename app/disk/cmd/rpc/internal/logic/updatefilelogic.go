package logic

import (
	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/app/disk/model"
	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"time"

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
	resp := pb.UpdateFileResp{}
	//	move
	if in.FileDetail.ParentId != "" {
		err := l.svcCtx.Engine.
			Table("user_repository").
			Where("identity = ? AND uid = ?", in.FileDetail.Identity, in.FileDetail.Uid).
			First(fileDetail).Error
		if err != nil && err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR 查询file失败，Pid:%s,Uid: %s,err:%v", in.FileDetail.ParentId, in.FileDetail.Uid, err)
		}
		if fileDetail.Id == 0 || err == model.ErrNotFound {
			return nil, errors.Wrapf(ErrFileNoExistsError, "Pid: %s Uid:%s", in.FileDetail.ParentId, in.FileDetail.Uid)
		}
		err = l.svcCtx.Engine.
			Table("user_repository").
			Where("identity = ? AND del_state = ?", in.FileDetail.Identity, globalkey.DelStateNo).
			Update("parent_id", in.FileDetail.ParentId).Error
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新user_repository失败")
		}
		return &resp, nil
	}

	//rename
	if in.FileDetail.Name != "" {
		var cnt int64
		err := l.svcCtx.Engine.
			Table("user_repository").
			Where("name = ?", in.FileDetail.Name).
			Where("parent_id = (select parent_id from user_repository  ur where ur.identity = ?)", in.FileDetail.Identity).
			Count(&cnt).Error
		if cnt > 0 {
			return nil, errors.New("文件名已存在")
		}

		println(in.FileDetail.Name)

		err = l.svcCtx.Engine.
			Table("user_repository").
			Where("identity = ? AND uid = ?", in.FileDetail.Identity, in.FileDetail.Uid).
			Update("name", in.FileDetail.Name).Error
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新user_repository失败")
		}
		return &resp, nil

	}

	//delete
	//TODO 定时任务
	if in.FileDetail.DelState != 0 {
		err := l.svcCtx.Engine.
			Table("user_repository").
			Where("identity = ? AND uid = ?", in.FileDetail.Identity, in.FileDetail.Uid).
			Update("del_state", globalkey.DelStateYes).
			Update("deleted_at", time.Now()).Error
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "更新user_repository失败")
		}
		return &resp, nil
	}

	return &resp, nil
}
