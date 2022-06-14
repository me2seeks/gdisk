// Code generated by goctl. DO NOT EDIT!
// Source: disk.proto

package server

import (
	"context"

	"cloud-disk/app/disk/cmd/rpc/internal/logic"
	"cloud-disk/app/disk/cmd/rpc/internal/svc"
	"cloud-disk/app/disk/cmd/rpc/pb"
)

type StoreServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedStoreServer
}

func NewStoreServer(svcCtx *svc.ServiceContext) *StoreServer {
	return &StoreServer{
		svcCtx: svcCtx,
	}
}

// store 详情
func (s *StoreServer) DetailStore(ctx context.Context, in *pb.StoreDetailReq) (*pb.StoreDetailResp, error) {
	l := logic.NewDetailStoreLogic(ctx, s.svcCtx)
	return l.DetailStore(in)
}

// 改变store大小
func (s *StoreServer) ChangeStore(ctx context.Context, in *pb.ChangeStoreReq) (*pb.ChangeStoreResp, error) {
	l := logic.NewChangeStoreLogic(ctx, s.svcCtx)
	return l.ChangeStore(in)
}
