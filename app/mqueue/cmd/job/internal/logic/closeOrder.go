package logic

import (
	"cloud-disk/app/mqueue/cmd/job/internal/svc"
	"cloud-disk/app/mqueue/cmd/job/jobtype"
	"cloud-disk/app/order/cmd/rpc/order"
	"cloud-disk/app/order/model"
	"cloud-disk/common/xerr"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
)

var ErrCloseOrderFal = xerr.NewErrMsg("close order fail")

// CloseHomestayOrderHandler 关闭没有支付的订单
type CloseHomestayOrderHandler struct {
	svcCtx *svc.ServiceContext
}

func NewCloseHomestayOrderHandler(svcCtx *svc.ServiceContext) *CloseHomestayOrderHandler {
	return &CloseHomestayOrderHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask   : if return err != nil , asynq will retry
func (l *CloseHomestayOrderHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	var p jobtype.DeferCloseHomestayOrderPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.Wrapf(ErrCloseOrderFal, "ERROR closeHomestayOrderStateMqHandler payload err:%v, payLoad:%+v", err, t.Payload())
	}

	resp, err := l.svcCtx.OrderRpc.HomestayOrderDetail(ctx, &order.HomestayOrderDetailReq{
		Sn: p.Sn,
	})

	if err != nil || resp.HomestayOrder == nil {
		return errors.Wrapf(ErrCloseOrderFal, "ERROR closeHomestayOrderStateMqHandler  get order fail or order no exists err:%v, sn:%s ,HomestayOrder : %+v", err, p.Sn, resp.HomestayOrder)
	}
	//更改订单状态
	if resp.HomestayOrder.TradeState == model.HomestayOrderTradeStateWaitPay {
		resp, err := l.svcCtx.OrderRpc.UpdateHomestayOrderTradeState(ctx, &order.UpdateHomestayOrderTradeStateReq{
			Sn:         p.Sn,
			TradeState: model.HomestayOrderTradeStateCancel,
		})
		fmt.Println(resp)
		if err != nil {
			return errors.Wrapf(ErrCloseOrderFal, "ERROR CloseHomestayOrderHandler close order fail  err:%v, sn:%s ", err, p.Sn)
		}
	}
	return nil
}
