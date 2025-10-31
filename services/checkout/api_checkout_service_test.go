package checkout

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/core/stringutillib"
	"github.com/shopspring/decimal"
	"log"
	"testing"
)

func TestCreateCheckoutOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiCheckOutService{Client: client}

	req := CreateOrderRequest{
		MerchantTradeNo: "merchant-trade-no",
		OrderExpireTime: 1761714000000,
		OrderAmount:     decimal.NewFromInt(1),
		Currency:        "USDT",
		CancelURL:       "cancel-url",
		ChannelId:       "channel-id",

		Env: EnvRequest{
			TerminalType: "APP",
		},
		Goods: GoodsRequest{
			GoodsDetail: "goods",
			GoodsName:   "goods",
		},
		Chain:          "ETH",
		FullCurrType:   "USDT_ETH",
		MerchantUserId: 0,
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.CreateCheckOutOrder(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestCreateCheckoutCreateRefund(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiCheckOutService{Client: client}

	req := CreateCheckStandRefundRequest{
		PrepayID:            "order-id",
		RefundRequestID:     "request-id",
		RefundOrderCurrency: "USDT",
		RefundOrderAmount:   decimal.NewFromFloat(1),
		RefundPayCurrency:   "USDT",
		RefundPayAmount:     decimal.NewFromFloat(1),
		RefundReason:        "test",
		ReceiverId:          0,
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.CreateCheckoutRefund(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}

}
