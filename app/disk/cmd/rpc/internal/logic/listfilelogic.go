package logic

import (
	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/app/disk/model"
	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFileLogic {
	return &ListFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListFileLogic) ListFile(in *pb.ListFileReq) (*pb.ListFileResp, error) {
	resp := pb.ListFileResp{}
	if in.Uid != "" {
		err := l.svcCtx.Engine.
			Table("user_repository").
			Select("user_repository.id, user_repository.parent_id, user_repository.identity, "+
				"user_repository.repository_identity, user_repository.ext, user_repository.updated_at,"+
				"user_repository.name, repository_pool.path, repository_pool.size").
			Where("uid= ? ", in.Uid).
			Where("user_repository.del_state", globalkey.DelStateNo).
			Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
			Find(&resp.FileList).Error
		//Limit(size).
		//Offset(offset).
		if err != nil && err != model.ErrNotFound {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR 查询file list失败，Uid: %s,err:%v", in.Uid, err)
		}
	}
	err := l.svcCtx.Engine.
		Table("user_repository").
		Select("user_repository.id, user_repository.parent_id, user_repository.identity, "+
			"user_repository.repository_identity, user_repository.ext, user_repository.updated_at,"+
			"user_repository.name, repository_pool.path, repository_pool.size").
		Where("uid= ? ", in.Uid).
		Where("user_repository.del_state", globalkey.DelStateYes).
		Joins("left join repository_pool on user_repository.repository_identity = repository_pool.identity").
		Find(&resp.DeletedList).Error
	//Limit(size).

	//Offset(offset).
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR 查询repository_pool list失败，Uid: %s,err:%v", in.Uid, err)
	}

	return &resp, nil
}
