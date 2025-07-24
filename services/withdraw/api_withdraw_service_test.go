package withdraw

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/core/stringutillib"
	"log"
	"testing"
)

func TestCreateWithDrawOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &APIWithdrawService{Client: client}

	/*
		{
		  "batch_id" : "237394559478075555",
		  "reason": "Payment for services",
		  "withdraw_list": [
		    {
		      "currency": "USDT",
		      "amount": "1",
		      "chain": "ETH",
		      "address": "0x1234567890abcdef"
		    },
		    {
		      "currency": "USDT",
		      "amount": "0.001",
		      "chain": "ETH",
		      "address": "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
		    }
		  ]
		}
	*/
	req := BatchWithdrawRequest{
		BatchID:   "237394559478075557",
		ChannelId: "",
		WithdrawList: []*WithdrawItem{
			{
				Currency:           "USDT",
				Amount:             "1",
				Chain:              "ETH",
				Address:            "0x1234567890abcdef",
				MerchantWithdrawId: "12345600",
			},
			{
				Currency:           "USDT",
				Amount:             "0.001",
				Chain:              "ETH",
				Address:            "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
				MerchantWithdrawId: "12345601",
			},
		},
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.WithdrawCreateOrder(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestWithdrawQuery(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &APIWithdrawService{Client: client}
	//接口没有参数，可以传common.BaseRequest,这样可以在BaseRequest传入client_id用于签名
	req := &BatchWithdrawQueryReq{
		BatchId: "237394559478075556",
		Status:  "ALL",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")

	resp, result, err := service.WithdrawQuery(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestGetCurrencyChains(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &APIWithdrawService{Client: client}

	req := WalletCurrency{
		Currency: "USDT",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")

	resp, result, err := service.GetCurrencyChains(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}
