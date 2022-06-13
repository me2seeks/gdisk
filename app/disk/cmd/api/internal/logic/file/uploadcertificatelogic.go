package file

import (
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadCertificateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadCertificateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadCertificateLogic {
	return &UploadCertificateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadCertificateLogic) UploadCertificate(req *types.UploadCertificateReq) (resp *types.UploadCertificateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
