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
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiCheckOutService{Client: client}

	req := CreateOrderRequest{
		MerchantTradeNo: "j4058308409230424822343119",
		OrderExpireTime: 1753158265000,
		OrderAmount:     decimal.NewFromInt(1),
		Currency:        "USDT",
		CancelURL:       "https://www.baidu.com",
		ChannelId:       "",

		Env: EnvRequest{
			Scene:        "APP",
			TerminalType: "IOS",
		},
		Goods: GoodsRequest{
			GoodsDetail: "goods",
			GoodsName:   "goods",
		},
		Chain:          "ETH",
		FullCurrType:   "USDT_ETH",
		MerchantUserId: 6790011,
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.CreateCheckOutOrder(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestCreateCheckoutCreateRefnund(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiCheckOutService{Client: client}

	req := CreateCheckStandRefundRequest{
		/*
			RefundRequestID     string          `json:"refundRequestId"`
			PrepayID            string          `json:"prepayId"`
			RefundOrderCurrency string          `json:"refundOrderCurrency"`
			RefundOrderAmount   decimal.Decimal `json:"refundOrderAmount"`
			RefundPayCurrency   string          `json:"refundPayCurrency"`
			RefundPayAmount     decimal.Decimal `json:"refundPayAmount"`
			RefundReason        string          `json:"refundReason"`
			ReceiverId          int64           `json:"receiverId"`
		*/
		PrepayID:            "380667380540264448",
		RefundOrderCurrency: "USDT",
		RefundOrderAmount:   decimal.NewFromFloat(1.00011),
		RefundPayCurrency:   "USDT",
		RefundPayAmount:     decimal.NewFromFloat(1.00011),
		RefundReason:        "test",
		ReceiverId:          10002,
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.CreateCheckoutRefund(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}

}
