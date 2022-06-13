// Code generated by goctl. DO NOT EDIT!
// Source: disk.proto

package payment

import (
	"context"

	"cloud-disk/app/disk/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	JudgeReq  = pb.JudgeReq
	JudgeResp = pb.JudgeResp

	Payment interface {
		// 判断是否能存入
		JudgeAllowed(ctx context.Context, in *JudgeReq, opts ...grpc.CallOption) (*JudgeResp, error)
	}

	defaultPayment struct {
		cli zrpc.Client
	}
)

func NewPayment(cli zrpc.Client) Payment {
	return &defaultPayment{
		cli: cli,
	}
}

// 判断是否能存入
func (m *defaultPayment) JudgeAllowed(ctx context.Context, in *JudgeReq, opts ...grpc.CallOption) (*JudgeResp, error) {
	client := pb.NewPaymentClient(m.cli.Conn())
	return client.JudgeAllowed(ctx, in, opts...)
}