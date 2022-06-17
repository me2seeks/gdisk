package folder

import (
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/globalkey"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
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

	var folder *pb.FolderDetail

	folder.UserId = uId

	if req.IsFile {
		var file *pb.FileDetail
		file.UserId = uId
		file.Id = req.Id

		switch req.ToPid {
		//重命名
		case globalkey.RenameTypeMove:
			file.FileName = req.Name
			_, err := l.svcCtx.DiskRpc.UpdateFile(l.ctx, &pb.UpdateFileReq{
				File: file,
			})
			if err != nil {
				return nil, err
			}
		//删除
		case globalkey.DeleteTypeMove:
			file.DeleteState = globalkey.DelStateYes
			_, err := l.svcCtx.DiskRpc.UpdateFile(l.ctx, &pb.UpdateFileReq{
				File: file,
			})
			if err != nil {
				return nil, err
			}
		//移动
		default:
			file.ParentFolderId = req.ToPid
			_, err := l.svcCtx.DiskRpc.UpdateFile(l.ctx, &pb.UpdateFileReq{
				File: file,
			})
			if err != nil {
				return nil, err
			}
		}
	}

	return &types.MovedResp{}, nil
}
