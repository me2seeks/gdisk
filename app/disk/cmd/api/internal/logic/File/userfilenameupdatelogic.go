package File

import (
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"context"

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

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	var fileDetail *pb.FileDetail
	fileDetail.Identity = req.Identity
	fileDetail.Pid = req.Name
	fileDetail.Uid = userIdentity

	_, err = l.svcCtx.DiskRpc.UpdateFile(l.ctx, &pb.UpdateFileReq{FileDetail: fileDetail})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
