package logic

import (
	"cloud-disk/app/payment/model"
	"cloud-disk/common/xerr"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"cloud-disk/app/payment/cmd/rpc/internal/svc"
	"cloud-disk/app/payment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentSuccessRefundByOrderSnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentSuccessRefundByOrderSnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentSuccessRefundByOrderSnLogic {
	return &GetPaymentSuccessRefundByOrderSnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPaymentSuccessRefundByOrderSn 根据订单sn查询流水记录
func (l *GetPaymentSuccessRefundByOrderSnLogic) GetPaymentSuccessRefundByOrderSn(in *pb.GetPaymentSuccessRefundByOrderSnReq) (*pb.GetPaymentSuccessRefundByOrderSnResp, error) {
	whereBuilder := l.svcCtx.ThirdPaymentModel.RowBuilder().Where(
		"order_sn = ? and (trade_state = ? or trade_state = ? )",
		in.OrderSn, model.ThirdPaymentPayTradeStateSuccess, model.ThirdPaymentPayTradeStateRefund,
	)
	thirdPayment, err := l.svcCtx.ThirdPaymentModel.FindOneByQuery(l.ctx, whereBuilder)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get payment record fail"), "get payment record fail FindOneByQuery  err : %v , orderSn:%s", err, in.OrderSn)
	}

	var resp pb.PaymentDetail
	if thirdPayment != nil {

		_ = copier.Copy(&resp, thirdPayment)
		resp.CreateTime = thirdPayment.CreateTime.Unix()
		resp.UpdateTime = thirdPayment.UpdateTime.Unix()
		resp.PayTime = thirdPayment.PayTime.Unix()

	}

	return &pb.GetPaymentSuccessRefundByOrderSnResp{
		PaymentDetail: &resp,
	}, nil
}
