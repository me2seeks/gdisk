package fileFolder

import (
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatefolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatefolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatefolderLogic {
	return &CreatefolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatefolderLogic) Createfolder(req *types.CreateFolderReq) (resp *types.CreateFolderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
