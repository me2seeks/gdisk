package File

import (
	"context"

	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebsocketMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebsocketMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebsocketMessageLogic {
	return &WebsocketMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebsocketMessageLogic) WebsocketMessage(req *types.WebsocketMessageRequest) (resp *types.WebsocketMessageReply, err error) {
	// todo: add your logic here and delete this line

	return
}
