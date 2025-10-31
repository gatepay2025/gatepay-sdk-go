package payment

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/core/stringutillib"
	"github.com/shopspring/decimal"
	"log"
	"testing"
)

func TestGetOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := OperateOrderRequest{PrepayID: "order-id", MerchantTradeNo: "merchant-trade-no"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetOrder(ctx, req)
	if err != nil {
		log.Printf("call GetAddressChains err:%s", err)
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleGetOrder() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := OperateOrderRequest{PrepayID: "order-id", MerchantTradeNo: "merchant-trade-no"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetOrder(ctx, req)
	if err != nil {
		log.Printf("call GetAddressChains err:%s", err)
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestOrderList(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := BalanceHistoryListReq{
		StartTime: 1760514000000,
		EndTime:   1760514000000 + 6000000,
		Currency:  "USDT",
		Count:     3,
		Page:      1,
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetOrderList(ctx, req)
	if err != nil {
		log.Printf("call GetAddressChains err:%s", err)
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleOrderList() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := BalanceHistoryListReq{
		StartTime: 1746758276000,
		EndTime:   1746772913000,
		Count:     3,
		Page:      1,
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetOrderList(ctx, req)
	if err != nil {
		log.Printf("call GetAddressChains err:%s", err)
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/order/query
func TestOrderQuery(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}
	req := OperateOrderRequest{PrepayID: "order-id"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")

	resp, result, err := service.GetOrder(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/order
func TestCreateOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}
	req := CreateOrderRequest{
		MerchantTradeNo: "merchant-trade-no",
		OrderExpireTime: 1761714000000,
		OrderAmount:     decimal.NewFromInt(1),
		Currency:        "USDT",
		CancelURL:       "cancel-url",
		ChannelId:       "channel-id",

		Goods: GoodsRequest{
			GoodsDetail: "goods",
			GoodsName:   "goods",
		},
		MerchantUserId: 0,
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.CreateOrder(ctx, req)
	if err != nil {
		log.Printf("call TestCreateOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleCreateOrder() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}
	req := CreateOrderRequest{
		MerchantTradeNo: "merchant-trade-no",
		OrderExpireTime: 1753439448000,
		OrderAmount:     decimal.NewFromInt(1),
		Currency:        "USDT",
		CancelURL:       "cancel-url",
		ChannelId:       "channel-id",

		Goods: GoodsRequest{
			GoodsDetail: "goods",
			GoodsName:   "goods",
		},
		MerchantUserId: 0,
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.CreateOrder(ctx, req)
	if err != nil {
		log.Printf("call TestCreateOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/order/refund/query
func TestRefundQuery(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}
	req := QueryRefundRequest{RefundRequestID: "request-id"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")

	resp, result, err := service.QueryRefund(ctx, req)
	if err != nil {
		log.Printf("call TestRefundQuery err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleRefundQuery() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}
	req := QueryRefundRequest{RefundRequestID: "request-id"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")

	resp, result, err := service.QueryRefund(ctx, req)
	if err != nil {
		log.Printf("call TestRefundQuery err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/order/refund
func TestCreateRefundOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}
	req := CreateRefundRequest{
		RefundRequestID: "request-id",
		PrepayID:        "order-id",
		RefundAmount:    decimal.NewFromInt(1),
		RefundReason:    "abc",
		RefundGateId:    "",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.CreateRefundOder(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/batch/transfer
func TestBatchTransferOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := BatchTransfer{
		BatchId:         "batch-id",
		MerchantID:      "merchant-id",
		MerchantBatchNo: "merchant-trade-no",
		Currency:        "GT",
		Name:            "batch-transfer",
		Description:     "batch-transfer",
		BizScene:        BizDirectTransfer,
		OrderList: []*BatchOrderList{
			{
				UserId: 0,
				Amount: "1.33",
			},
		},
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.CreateTransfer(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleBatchTransferOrder() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := BatchTransfer{
		MerchantID:      "merchant-id",
		MerchantBatchNo: "merchant-batch-no",
		Currency:        "GT",
		Name:            "batch-transfer",
		Description:     "batch-transfer",
		BizScene:        BizSceneType(BizRewards),
		OrderList: []*BatchOrderList{
			{
				UserId: 0,
				Amount: "1.33",
			},
		},
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.CreateTransfer(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/batch/transfer/query
func TestTransferOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := BatchTransferQueryReq{
		BatchId:         "batch-id",
		MerchantBatchNo: "merchant-batch-no",
		Status:          "ALL",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.BatchTransferQuery(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleTransferOrder() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := BatchTransferQueryReq{
		BatchId:         "batch-id",
		MerchantBatchNo: "merchant-batch-no",
		Status:          "ALL",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.BatchTransferQuery(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/transactions/native
func TestNativeOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := CreateOrderRequest{
		MerchantTradeNo: "merchant-trade-no",
		OrderExpireTime: 1761717600000,
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
	resp, result, err := service.CreateNativeOrder(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleNativeOrder() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := CreateOrderRequest{
		MerchantTradeNo: "merchant-trade-no",
		OrderExpireTime: 1753443048000,
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
	resp, result, err := service.CreateNativeOrder(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}
