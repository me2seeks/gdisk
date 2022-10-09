package logic

import (
	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/app/disk/model"
	"context"

	"cloud-disk/common/oss"
	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadPrepareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(in *pb.FileUploadPrepareRep) (*pb.FileUploadPrepareResp, error) {
	rp := new(model.RepositoryPool)
	resp := new(pb.FileUploadPrepareResp)

	l.svcCtx.Engine.
		Where("hash = ?", in.Md5).
		First(&rp)

	if rp.Id != 0 {
		// 文件已存在，秒传成功
		resp.Identity = rp.Identity
	} else {
		// 文件不存在，获取文件的 UploadID、key，执行分片上传
		key, uploadId, err := oss.CosInitPart(in.Ext)
		if err != nil {
			return resp, err
		}
		resp.Key = key
		resp.UploadId = uploadId
	}

	return resp, nil
}
