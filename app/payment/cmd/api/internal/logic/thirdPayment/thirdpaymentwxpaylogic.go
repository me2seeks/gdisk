package thirdPayment

import (
	"cloud-disk/app/payment/cmd/rpc/payment"
	"cloud-disk/app/payment/model"
	"cloud-disk/app/usercenter/cmd/rpc/usercenter"
	usercenterModel "cloud-disk/app/usercenter/model"
	"cloud-disk/common/ctxdata"
	"cloud-disk/common/xerr"
	"context"
	"github.com/pkg/errors"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"

	"cloud-disk/app/payment/cmd/api/internal/svc"
	"cloud-disk/app/payment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrWxPayError = xerr.NewErrMsg("wechat pay fail")

type ThirdPaymentWxPayLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewThirdPaymentWxPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThirdPaymentWxPayLogic {
	return &ThirdPaymentWxPayLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ThirdPaymentWxPayLogic) ThirdPaymentWxPay(req *types.ThirdPaymentWxPayReq) (*types.ThirdPaymentWxPayResp, error) {
	var totalPrice int64
	var decription string
	switch req.ServiceType {
	case model.ThirdPaymentServiceTypeHomestayOrder:
		homestayTotalPrice, homestayDecription, err := l.getHomestayPriceOrder(req.OrderSn)
		if err != nil {
			return nil, errors.Wrapf(ErrWxPayError, "getHomestayPriceOrder err: %v req: %+v", err, req)
		}
		totalPrice = homestayTotalPrice
		decription = homestayDecription
	default:
		return nil, errors.Wrapf(xerr.NewErrMsg("Payment for this business type is not supported"), "Payment for this business type is not supported req: %+v", req)
	}
	//创建预处理订单
	wechatPrepayRsp, err := l.createWxPayOrder(totalPrice, req.OrderSn, req.ServiceType, decription)
	if err != nil {
		return nil, err
	}

	return &types.ThirdPaymentWxPayResp{
		Appid:     l.svcCtx.Config.WxMiniConf.AppId,
		NonceStr:  *wechatPrepayRsp.NonceStr,
		PaySign:   *wechatPrepayRsp.PaySign,
		Package:   *wechatPrepayRsp.Package,
		Timestamp: *wechatPrepayRsp.TimeStamp,
		SignType:  *wechatPrepayRsp.SignType,
	}, nil
}

func (l *ThirdPaymentWxPayLogic) getHomestayPriceOrder(orderSn string) (int64, string, error) {
	description := "homestay pay"
	//获取订单详情
	resp, err := l.svcCtx.OrderRpc.HomestayOrderDetail(l.ctx, &order.HomestayOrderDetailReq{
		Sn: orderSn,
	})
	if err != nil {
		return 0, description, errors.Wrapf(ErrWxPayError, "HomestayOrderDetail err: %v orderSn: %s", err, orderSn)
	}
	if resp.HomestayOrder == nil || resp.HomestayOrder.Id == 0 {
		return 0, description, errors.Wrapf(xerr.NewErrMsg("order no exists"), "WeChat payment order does not exist orderSn : %s", orderSn)
	}

	return resp.HomestayOrder.OrderTotalPrice, description, nil
}

func (l *ThirdPaymentWxPayLogic) createWxPayOrder(totalPrice int64, orderSn, serviceType, description string) (*jsapi.PrepayWithRequestPaymentResponse, error) {
	//获取用户openid
	userId := ctxdata.GetUidFromCtx(l.ctx)
	userResp, err := l.svcCtx.UsercenterRpc.GetUserAuthByUserId(l.ctx, &usercenter.GetUserAuthByUserIdReq{
		UserId:   userId,
		AuthType: usercenterModel.UserAuthTypeSmallWX,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrWxPayError, "Get user wechat openid err : %v , userId: %d , orderSn:%s", err, userId, orderSn)
	}
	if userResp.UserAuth == nil || userResp.UserAuth.UserId == 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("Get user wechat openid fail，Please pay before authorization by weChat"), "Get user WeChat openid does not exist  userId: %d , orderSn:%s", userId, orderSn)
	}
	openId := userResp.UserAuth.AuthKey

	//创建第三方支付本地记录
	createPaymentResp, err := l.svcCtx.PaymentRpc.CreatePayment(l.ctx, &payment.CreatePaymentReq{
		UserId:      userId,
		PayModel:    model.ThirdPaymentPayModelWechatPay,
		PayTotal:    totalPrice,
		OrderSn:     orderSn,
		ServiceType: serviceType,
	})
	if err != nil || createPaymentResp.Sn == "" {
		return nil, errors.Wrapf(ErrWxPayError,
			"create local third payment record fail : err: %v , userId: %d,totalPrice: %d , orderSn: %s",
			err, userId, totalPrice, orderSn)
	}

	//创建微信支付预付单
	wxPayClient, err := svc.NewWxPayClientV3(l.svcCtx.Config)
	if err != nil {
		return nil, err
	}
	jsApiSvc := jsapi.JsapiApiService{
		Client: wxPayClient,
	}
	//获取prepay_id，以及调用支付所需的参数和签名
	resp, _, err := jsApiSvc.PrepayWithRequestPayment(l.ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(l.svcCtx.Config.WxMiniConf.AppId),
			Mchid:       core.String(l.svcCtx.Config.WxPayConf.MchId),
			Description: core.String(description),
			OutTradeNo:  core.String(createPaymentResp.Sn),
			Attach:      core.String(description),
			NotifyUrl:   core.String(l.svcCtx.Config.WxPayConf.NotifyUrl),
			Amount: &jsapi.Amount{
				Total: core.Int64(totalPrice),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(openId),
			},
		})
	if err != nil {
		return nil, errors.Wrapf(ErrWxPayError, "Failed to initiate WeChat payment pre-order err : %v , userId: %d , orderSn:%s", err, userId, orderSn)
	}

	return resp, nil
}
