package File

import (
	"cloud-disk/app/define"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
	"cloud-disk/app/disk/cmd/rpc/pb"
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

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	usrFile := make([]*types.UserFile, 0)
	deletedFile := make([]*types.DeletedUserFile, 0)
	var cnt int64
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
	file, err := l.svcCtx.DiskRpc.ListFile(l.ctx, &pb.ListFileReq{Uid: userIdentity})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&usrFile, file)
	
	err = l.svcCtx.Engine.
		Table("user_repository").
		Select("user_repository.id, user_repository.pid, user_repository.identity, "+
			"user_repository.repository_identity, user_repository.ext, user_repository.deleted_at,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Where("uid = ? ", userIdentity).
		Where("user_repository.deleted_at IS NOT NULL").
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Find(&deletedFile).Error

	if err != nil {
		return
	}

	// 查询总数
	err = l.svcCtx.Engine.
		Table("user_repository").
		// TODO pid = ? AND
		Where("uid = ? AND deleted_at IS NULL", userIdentity).
		Count(&cnt).Error
	if err != nil {
		return
	}

	resp.List = usrFile
	resp.DeletedList = deletedFile
	resp.Count = cnt

	return
}
