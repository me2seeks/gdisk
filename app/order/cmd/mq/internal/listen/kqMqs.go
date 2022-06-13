package listen

import (
	"cloud-disk/app/order/cmd/mq/internal/config"
	kqMq "cloud-disk/app/order/cmd/mq/internal/mqs/kq"
	"cloud-disk/app/order/cmd/mq/internal/svc"
	"context"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

//pub sub use kq (kafka)
func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.PaymentUpdateStatusConf, kqMq.NewPaymentUpdateStatusMq(ctx, svcContext)),
		//.....
	}
}
