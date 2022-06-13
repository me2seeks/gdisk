package logic

import (
	"context"

	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeAllowedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeAllowedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeAllowedLogic {
	return &JudgeAllowedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 判断是否能存入
func (l *JudgeAllowedLogic) JudgeAllowed(in *pb.JudgeReq) (*pb.JudgeResp, error) {
	// todo: add your logic here and delete this line

	return &pb.JudgeResp{}, nil
}
