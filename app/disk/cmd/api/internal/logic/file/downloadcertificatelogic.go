package file

import (
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownloadCertificateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadCertificateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownloadCertificateLogic {
	return &DownloadCertificateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadCertificateLogic) DownloadCertificate(req *types.DownloadCertificateReq) (resp *types.DownloadCertificateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
