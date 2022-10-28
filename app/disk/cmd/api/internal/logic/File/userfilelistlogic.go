package File

import (
	"cloud-disk/app/define"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/common/ctxdata"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest) (resp *types.UserFileListReply, err error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)

	resp = new(types.UserFileListReply)

	// 分页参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}
	page := req.Page
	if page == 0 {
		page = 1
	}
	// offset := (page - 1) * size

	// TODO 按文件名查询
	// id := req.Id =》 pid
	// if id == 0 {
	// 	id = -1
	// }

	//搜出用户的所以文件
	fl, err := l.svcCtx.DiskRpc.ListFile(l.ctx, &pb.ListFileReq{Uid: uid})
	if err != nil {
		return nil, err
	}

	_ = copier.Copy(&resp.List, fl.FileList)

	_ = copier.Copy(&resp.DeletedList, fl.DeletedList)

	// 查询总数
	err = l.svcCtx.Engine.
		Table("user_repository").
		// TODO pid = ? AND
		Where("uid = ? AND deleted_at IS NULL", uid).
		Count(&resp.Count).Error
	if err != nil {
		return
	}

	return
}
