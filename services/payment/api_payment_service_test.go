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
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := OperateOrderRequest{PrepayID: "382552424963100672", MerchantTradeNo: ""}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.GetOrder(ctx, req)
	if err != nil {
		log.Printf("call GetAddressChains err:%s", err)
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestOrderList(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
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
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
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
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}
	req := OperateOrderRequest{MerchantTradeNo: "j4058308409230424822343104", PrepayID: "372425901225357312"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")

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
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}
	req := CreateOrderRequest{
		MerchantTradeNo: "j4058308409230424822343668",
		OrderExpireTime: 1753158265000,
		OrderAmount:     decimal.NewFromInt(1),
		Currency:        "USDT",
		CancelURL:       "https://www.baidu.com",
		ChannelId:       "",

		//Env: EnvRequest{
		//	Scene:        "APP",
		//	TerminalType: "IOS",
		//},
		Goods: GoodsRequest{
			GoodsDetail: "goods",
			GoodsName:   "goods",
		},
		//Chain:          "ETH",
		//FullCurrType:   "USDT_ETH",
		MerchantUserId: 6790011,
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
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
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}
	req := QueryRefundRequest{RefundRequestID: "4839024899928418051"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")

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
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}
	req := CreateRefundRequest{
		RefundRequestID: "35018077417507388",
		PrepayID:        "35018077417504835",
		RefundAmount:    decimal.NewFromInt(1),
		RefundReason:    "abc",
		RefundGateId:    "",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
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
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	/*
		{
		    "merchant_batch_no": "1221234567826035",
		    "merchant_id": "10002",
		    "currency": "GT",
		    "name ": "batch-transfer",
		    "description ": "batch-transfer",
		    "bizscene" : "REWARDS",
		    "batchorderList": [
		        {
		            "user_id": 10000,
		            "amount": "1.33"
		        }
		    ]
		}
	*/

	req := BatchTransfer{
		MerchantID:      "10002",
		MerchantBatchNo: "1221234567826061",
		Currency:        "GT",
		Name:            "batch-transfer",
		Description:     "batch-transfer",
		BizScene:        BizSceneType(BizRewards),
		OrderList: []*BatchOrderList{
			{
				UserId: 10000,
				Amount: "1.33",
			},
		},
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
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
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := BatchTransferQueryReq{
		BatchId:         "382211210518921216",
		MerchantBatchNo: "1221234567826061",
		Status:          "ALL",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
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
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := CreateOrderRequest{
		MerchantTradeNo: "j4058308409230424822343666",
		OrderExpireTime: 1753090549000,
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
	resp, result, err := service.CreateNativeOrder(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}
