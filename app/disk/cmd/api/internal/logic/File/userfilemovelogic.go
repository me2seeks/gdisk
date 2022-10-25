package File

import (
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/common/ctxdata"
	"context"

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

	var fileDetail *pb.FileDetail
	fileDetail.Identity = req.Identity
	fileDetail.ParentId = req.ParentIdentity
	fileDetail.Uid = userIdentity

	_, err = l.svcCtx.DiskRpc.UpdateFile(l.ctx, &pb.UpdateFileReq{FileDetail: fileDetail})
	if err != nil {
		return nil, err
	}
	return nil, nil

}
