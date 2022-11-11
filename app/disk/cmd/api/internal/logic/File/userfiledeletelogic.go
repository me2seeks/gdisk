package File

import (
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/globalkey"
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileDeleteLogic {
	return &UserFileDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileDeleteLogic) UserFileDelete(req *types.UserFileDeleteRequest) (resp *types.UserFileDeleteReply, err error) {
	u := ctxdata.GetUidFromCtx(l.ctx)

	var fileDetail *pb.FileDetail
	fileDetail.Identity = req.Identity
	fileDetail.DelState = globalkey.DelStateYes
	fileDetail.Uid = u

	_, err = l.svcCtx.DiskRpc.UpdateFile(l.ctx, &pb.UpdateFileReq{FileDetail: fileDetail})
	if err != nil {
		return nil, err
	}
	return nil, nil

}
