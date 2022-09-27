package File

import (
	"cloud-disk/app/disk/cmd/rpc/disk"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (*types.FileUploadPrepareReply, error) {
	fileUploadPrepareReply, err := l.svcCtx.DiskRpc.FileUploadPrepare(l.ctx, &disk.FileUploadPrepareRep{
		Md5:  req.Md5,
		Name: req.Name,
		Ext:  req.Ext,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "disk.FileUploadPrepareRep Failed req: %+v", req)
	}
	var resp types.FileUploadPrepareReply
	_ = copier.Copy(&resp, fileUploadPrepareReply)

	return &resp, nil
}
