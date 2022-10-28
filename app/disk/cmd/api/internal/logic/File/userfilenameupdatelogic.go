package File

import (
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/globalkey"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest) (resp *types.UserFileNameUpdateReply, err error) {
	userIdentity := ctxdata.GetUidFromCtx(l.ctx)

	resp = new(types.UserFileNameUpdateReply)
	if req.Name == "" {
		return nil, errors.New("文件名为空")
	}

	// 判断当前文件名在该层级下是否已存在
	var cnt int64
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("name = ?", req.Name).
		Where("parent_id = (select parent_id from user_repository ur where ur.identity = ?)", req.Identity).
		Where("del_state", globalkey.DelStateYes).
		Count(&cnt).Error

	if err != nil || cnt > 0 {
		return nil, errors.New("文件名已存在")
	}

	var fileDetail *pb.FileDetail
	fileDetail.Identity = req.Identity
	fileDetail.Name = req.Name
	fileDetail.Uid = userIdentity

	_, err = l.svcCtx.DiskRpc.UpdateFile(l.ctx, &pb.UpdateFileReq{FileDetail: fileDetail})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
