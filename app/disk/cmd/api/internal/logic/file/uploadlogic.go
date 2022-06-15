package file

import (
	"cloud-disk/app/disk/cmd/rpc/pb"
	"cloud-disk/app/disk/cmd/rpc/store"
	"cloud-disk/app/disk/model"
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/tool"
	"cloud-disk/common/upload"
	"cloud-disk/common/xerr"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"path"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *types.UploadCertificateReq) (*types.UploadCertificateResp, error) {
	var needSize int64 = 0
	uId := ctxdata.GetUidFromCtx(l.ctx)

	storeDetail, err := l.svcCtx.StoreRpc.DetailStore(l.ctx, &store.StoreDetailReq{
		Uid: uId,
	})
	if err != nil {
		return nil, err
	}

	for _, file := range req.Files {
		needSize += file.Size
	}
	if needSize > (storeDetail.Store.MaxSize - storeDetail.Store.CurrentSize) {
		return nil, errors.Wrapf(xerr.NewErrMsg("容量不够"), "store 不够  uid: %d", uId)
	}

	if err = l.svcCtx.FileModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		for _, file := range req.Files {
			//检验该路径下是否有同名并处理
			builder := l.svcCtx.FileModel.RowBuilder().Where(squirrel.Eq{
				"file_name": file.FileName,
				"file_path": file.Path,
			})
			_, err := l.svcCtx.FileModel.FindOneByQuery(l.ctx, builder)
			switch err {
			case nil:
				file.FileName = tool.SameName(file.FileName)
			case model.ErrNotFound:
			default:
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR FileModel.FindOneByQuery err: %v", err)
			}

			//更改用户容量
			_, err = l.svcCtx.StoreRpc.ChangeStore(l.ctx, &pb.ChangeStoreReq{
				Uid:         uId,
				CurrentSize: storeDetail.Store.CurrentSize + needSize,
			})
			if err != nil {
				return err
			}

			//插入数据库
			_, err = l.svcCtx.FileModel.Insert(l.ctx, session, &model.File{
				UserId:   uId,
				FileName: file.FileName,
				StoreId:  storeDetail.Store.Id,
				FilePath: file.Path,
				Size:     file.Size,
				Postfix:  path.Ext(file.FileName),
			})
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ERROR  FileModel.Insert err: %v", err)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	//延迟删除

	//上传凭证
	certificate := upload.UploadCertificate(needSize)
	return &types.UploadCertificateResp{
		Certificate: certificate,
	}, nil
}
