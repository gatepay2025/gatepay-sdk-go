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
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &APIWithdrawService{Client: client}

	req := BatchWithdrawRequest{
		BatchID:   "batch-id",
		ChannelId: "channel-id",
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
				Address:            "0x1234567890abcdefg",
				MerchantWithdrawId: "12345601",
			},
		},
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.WithdrawCreateOrder(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleCreateWithDrawOrder() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &APIWithdrawService{Client: client}
	req := BatchWithdrawRequest{
		BatchID:   "batch-id",
		ChannelId: "channel-id",
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
				Address:            "0x1234567890abcdefg",
				MerchantWithdrawId: "12345601",
			},
		},
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.WithdrawCreateOrder(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestWithdrawQuery(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &APIWithdrawService{Client: client}
	//接口没有参数，可以传common.BaseRequest,这样可以在BaseRequest传入client_id用于签名
	req := &BatchWithdrawQueryReq{
		BatchId: "batch-id",
		Status:  "PROCESSING",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")

	resp, result, err := service.WithdrawQuery(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleWithdrawQuery() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &APIWithdrawService{Client: client}
	//接口没有参数，可以传common.BaseRequest,这样可以在BaseRequest传入client_id用于签名
	req := &BatchWithdrawQueryReq{
		BatchId: "batch-id",
		Status:  "ALL",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")

	resp, result, err := service.WithdrawQuery(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestGetCurrencyChains(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("")
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
	req.AddHeader("X-GatePay-Certificate-ClientId", "")

	resp, result, err := service.GetCurrencyChains(ctx, req)
	if err != nil {
		log.Printf("call TestGetCurrencyChains err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%s", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleGetCurrencyChains() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
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
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")

	resp, result, err := service.GetCurrencyChains(ctx, req)
	if err != nil {
		log.Printf("call TestGetCurrencyChains err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%s", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestGetTotalBalance(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("")
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
	req.AddHeader("X-GatePay-Certificate-ClientId", "")

	resp, result, err := service.GetTotalBalance(ctx, req)
	if err != nil {
		log.Printf("call TestGetCurrencyChains err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%s", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestGetWithdrawStatus(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("")
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
	req.AddHeader("X-GatePay-Certificate-ClientId", "")

	resp, result, err := service.GetWithdrawStatus(ctx, req)
	if err != nil {
		log.Printf("call TestGetCurrencyChains err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%s", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}
