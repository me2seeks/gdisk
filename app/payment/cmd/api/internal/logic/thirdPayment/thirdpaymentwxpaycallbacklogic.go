package thirdPayment

import (
	"cloud-disk/app/payment/cmd/rpc/payment"
	"cloud-disk/app/payment/model"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"net/http"

	"cloud-disk/app/payment/cmd/api/internal/svc"
	"cloud-disk/app/payment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

//感觉需要dtm
var ErrWxPayCallbackError = xerr.NewErrMsg("wechat pay callback fail")

type ThirdPaymentWxPayCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThirdPaymentWxPayCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdPaymentWxPayCallbackLogic {
	return &ThirdPaymentWxPayCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdPaymentWxPayCallbackLogic) ThirdPaymentWxPayCallback(rw http.ResponseWriter, req *http.Request) (*types.ThirdPaymentWxPayCallbackResp, error) {
	//检索本地商户证书私钥。
	_, err := svc.NewWxPayClientV3(l.svcCtx.Config)
	if err != nil {
		return nil, err
	}

	// Get platform certificate accessor
	certVisitor := downloader.MgrInstance().GetCertificateVisitor(l.svcCtx.Config.WxPayConf.MchId)
	handler := notify.NewNotifyHandler(l.svcCtx.Config.WxPayConf.APIv3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))
	//Verifying signatures, parsing data
	transaction := new(payments.Transaction)
	_, err = handler.ParseNotifyRequest(context.Background(), req, transaction)
	if err != nil {
		return nil, errors.Wrapf(ErrWxPayCallbackError, "Failed to parse data ,err:%v", err)
	}

	returnCode := "SUCCESS"
	err = l.verifyAndUpdateState(transaction)
	if err != nil {
		returnCode = "FAIL"
	}

	return &types.ThirdPaymentWxPayCallbackResp{
		ReturnCode: returnCode,
	}, err
}

//Verify and update relevant flow data
func (l *ThirdPaymentWxPayCallbackLogic) verifyAndUpdateState(notifyTrasaction *payments.Transaction) error {
	paymentResp, err := l.svcCtx.PaymentRpc.GetPaymentBySn(l.ctx, &payment.GetPaymentBySnReq{
		Sn: *notifyTrasaction.OutTradeNo,
	})
	if err != nil || paymentResp.PaymentDetail.Id != 0 {
		return errors.Wrapf(ErrWxPayCallbackError, "Failed to get payment flow record err:%v ,notifyTrasaction:%+v ", err, notifyTrasaction)
	}

	//	比对金额
	notifyPayTotal := *notifyTrasaction.Amount.PayerTotal
	if paymentResp.PaymentDetail.PayTotal != notifyPayTotal {
		return errors.Wrapf(ErrWxPayCallbackError, "Order amount exception  notifyPayTotal:%v , notifyTrasaction:%v ", notifyPayTotal, notifyTrasaction)
	}
	//判断状态
	payStatus := l.getPayStatusByWXPayTradeState(*notifyTrasaction.TradeState)
	if payStatus == model.ThirdPaymentPayTradeStateSuccess {
		//付款通知

		if paymentResp.PaymentDetail.PayStatus != model.ThirdPaymentPayTradeStateWait {
			return nil
		}

		//	修改状态
		if _, err := l.svcCtx.PaymentRpc.UpdateTradeState(l.ctx, &payment.UpdateTradeStateReq{
			Sn:             *notifyTrasaction.OutTradeNo,
			TradeState:     *notifyTrasaction.TradeState,
			TransactionId:  *notifyTrasaction.TransactionId,
			TradeType:      *notifyTrasaction.TradeType,
			TradeStateDesc: *notifyTrasaction.TradeStateDesc,
			PayStatus:      l.getPayStatusByWXPayTradeState(*notifyTrasaction.TradeState),
		}); err != nil {
			return errors.Wrapf(ErrWxPayCallbackError, "更新流水状态失败  err:%v , notifyTrasaction:%v ", err, notifyTrasaction)
		}
	} else if payStatus == model.ThirdPaymentPayTradeStateWait {
		//Refund notification @todo to be done later, not needed at this time
	}
	return nil
}

const (
	SUCCESS    = "SUCCESS"    //支付成功
	REFUND     = "REFUND"     //转入退款
	NOTPAY     = "NOTPAY"     //未支付
	CLOSED     = "CLOSED"     //已关闭
	REVOKED    = "REVOKED"    //已撤销（付款码支付）
	USERPAYING = "USERPAYING" //用户支付中（付款码支付）
	PAYERROR   = "PAYERROR"   //支付失败(其他原因，如银行返回失败)
)

func (l *ThirdPaymentWxPayCallbackLogic) getPayStatusByWXPayTradeState(wxPayTradeState string) int64 {

	switch wxPayTradeState {
	case SUCCESS: //支付成功
		return model.ThirdPaymentPayTradeStateSuccess
	case USERPAYING: //支付中
		return model.ThirdPaymentPayTradeStateWait
	case REFUND: //已退款
		return model.ThirdPaymentPayTradeStateWait
	default:
		return model.ThirdPaymentPayTradeStateFAIL
	}
}
