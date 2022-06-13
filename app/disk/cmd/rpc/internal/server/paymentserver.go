// Code generated by goctl. DO NOT EDIT!
// Source: disk.proto

package server

import (
	"context"

	"cloud-disk/app/disk/cmd/rpc/internal/logic"
	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"
)

type PaymentServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedPaymentServer
}

func NewPaymentServer(svcCtx *svc.ServiceContext) *PaymentServer {
	return &PaymentServer{
		svcCtx: svcCtx,
	}
}

// 判断是否能存入
func (s *PaymentServer) JudgeAllowed(ctx context.Context, in *pb.JudgeReq) (*pb.JudgeResp, error) {
	l := logic.NewJudgeAllowedLogic(ctx, s.svcCtx)
	return l.JudgeAllowed(in)
}