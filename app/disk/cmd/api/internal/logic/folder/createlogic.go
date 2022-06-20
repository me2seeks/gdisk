package folder

import (
	"cloud-disk/app/disk/cmd/rpc/disk"
	"cloud-disk/app/disk/model"
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"time"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateFolderReq) (*types.CreateFolderResp, error) {
	uId := ctxdata.GetUidFromCtx(l.ctx)
	storeDetail, err := l.svcCtx.DiskRpc.DetailStore(l.ctx, &disk.StoreDetailReq{
		Uid: uId,
	})
	if err != nil {
		return nil, err
	}

	var folder = new(model.Folder)
	folder.ParentFolderId = req.Pid
	folder.CreateTime = time.Now()
	folder.FolderName = req.FolderName
	folder.StoreId = storeDetail.Store.Id

	_, err = l.svcCtx.FolderModel.Insert(l.ctx, nil, folder)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR: Failed to FolderModel.Insert err: %v", err)
	}
	return &types.CreateFolderResp{
		CreateTime: folder.CreateTime.Unix(),
	}, nil
}
