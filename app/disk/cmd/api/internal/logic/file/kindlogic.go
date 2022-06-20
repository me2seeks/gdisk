package file

import (
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/globalkey"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/mr"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KindLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKindLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KindLogic {
	return &KindLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KindLogic) Kind(req *types.KindReq) (*types.KindResp, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	postfix := globalkey.Postfix[req.Kind]

	var fileDetails []types.FileDetail
	err := mr.MapReduceVoid(func(source chan<- interface{}) {
		l := len(postfix)
		for i := 0; i < l; i++ {
			source <- postfix[i]
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		builder := l.svcCtx.FileModel.RowBuilder().Where(squirrel.Eq{
			"user_id":   uid,
			"postfix":   item.(string),
			"del_state": globalkey.DelStateNo,
		})
		Files, err := l.svcCtx.FileModel.FindAll(l.ctx, builder, "desc")
		if err != nil {
			logx.WithContext(l.ctx).Errorf("获取种类文件失败，postfix: %s err: %v", item.(string), err)
			return
		}
		writer.Write(Files)
	}, func(pipe <-chan interface{}, cancel func(error)) {
		for item := range pipe {
			var fileDetail types.FileDetail
			_ = copier.Copy(fileDetail, item)

			fileDetails = append(fileDetails, fileDetail)
		}
	})
	if err != nil {
		return nil, err
	}
	return &types.KindResp{
		Files: fileDetails,
	}, nil
}
