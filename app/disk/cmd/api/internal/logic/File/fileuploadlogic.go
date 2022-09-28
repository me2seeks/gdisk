package File

import (
	"cloud-disk/app/disk/model"
	"cloud-disk/common/uuid"
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadReply, err error) {
	rp := &model.RepositoryPool{
		Identity: uuid.UUID(),
		Name:     req.Name,
		Hash:     req.Hash,
		Path:     req.Path,
		Ext:      req.Ext,
		Size:     req.Size,
	}
	resp = new(types.FileUploadReply)
	err = l.svcCtx.Engine.
		Select("identity", "name", "hash", "path", "ext", "size", "created_at", "updated_at").
		Create(rp).Error
	if err != nil {

		return
	}

	resp.Identity = rp.Identity
	resp.Ext = rp.Ext
	resp.Name = rp.Name
	return
}
