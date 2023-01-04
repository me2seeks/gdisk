package File

import (
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/app/disk/model"
	"cloud-disk/common/ctxdata"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest) (resp *types.UserFileMoveReply, err error) {
	userIdentity := ctxdata.GetUidFromCtx(l.ctx)

	parentData := new(model.UserRepository)
	err = l.svcCtx.Engine.
		Table("user_repository").
		Where("identity = ? ", req.ParentIdentity).
		First(parentData).Error

	if err != nil || parentData.Id == 0 {
		println(err.Error())
		return nil, errors.New("文件夹不存在")
	}

	var fileDetail pb.FileDetail

	fileDetail.ParentId = parentData.Id
	fileDetail.Identity = req.Identity
	fileDetail.Uid = userIdentity

	_, err = l.svcCtx.DiskRpc.UpdateFile(l.ctx, &pb.UpdateFileReq{FileDetail: &fileDetail})

	if err != nil {
		return nil, err
	}
	return nil, nil

}
