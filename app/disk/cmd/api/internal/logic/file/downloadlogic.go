package file

import (
	"cloud-disk/common/oss"
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadLogic {
	return &DownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadLogic) Download(req *types.DownloadCertificateReq) (*types.DownloadCertificateResp, error) {

	fileDetail, err := l.svcCtx.FileModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	downloadUrl := oss.DownloadUrl(fileDetail.FileHash)
	return &types.DownloadCertificateResp{
		DownloadUrl: downloadUrl,
	}, nil
}
