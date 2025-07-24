package convert

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/core/stringutillib"
	"log"
	"testing"
)

// /v1/pay/convert/order
func TestCreateConvertOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiConvertService{Client: client}

	req := ConvertOrderRequest{
		SellCurrency: "GT",
		BuyCurrency:  "USDT",
		BuyAmount:    "0.5",
		SellAmount:   "0.0801",
		QuoteId:      "PAY-e7b34a76",
		ClientReqId:  "1111",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.PayConvert(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/convert/currency
func TestGetConvertCurrency(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiConvertService{Client: client}

	req := ConvertSupportCurrencyRequest{
		Side: "sell",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.GetConvertCurrency(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/convert/pair
func TestGetConvertPair(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiConvertService{Client: client}

	req := ConvertPairRequest{
		Side:     "sell",
		Currency: "GT",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.GetConvertPair(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/convert/preview
func TestGetConvertPreview(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiConvertService{Client: client}

	req := ConvertPreviewRequest{
		SellCurrency: "GT",
		BuyCurrency:  "USDT",
		BuyAmount:    "1",
		SellAmount:   "0.0401",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.ConvertPreView(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/convert/order
func TestGetConvertOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiConvertService{Client: client}

	req := QueryConvertOrderReq{
		OrderID: "326850433152987136",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")
	resp, result, err := service.GetConvertOrder(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}
