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
	credentials := core.NewCredentials("secret-key")
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
		BuyAmount:    "1",
		SellAmount:   "0.0401",
		QuoteId:      "quote-id",
		ClientReqId:  "client-req-id",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
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
	credentials := core.NewCredentials("secret-key")
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
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetConvertCurrency(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleGetConvertCurrency() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
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
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
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
	credentials := core.NewCredentials("secret-key")
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
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetConvertPair(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleGetConvertPair() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
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
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
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
	credentials := core.NewCredentials("secret-key")
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
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.ConvertPreView(ctx, req)
	if err != nil {
		log.Printf("call TestGetConvertPreview err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleGetConvertPreview() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
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
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.ConvertPreView(ctx, req)
	if err != nil {
		log.Printf("call TestGetConvertPreview err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/convert/order
func TestGetConvertOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiConvertService{Client: client}

	req := QueryConvertOrderReq{
		OrderID: "order-id",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetConvertOrder(ctx, req)
	if err != nil {
		log.Printf("call TestGetConvertOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleGetConvertOrder() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiConvertService{Client: client}

	req := QueryConvertOrderReq{
		OrderID: "order-id",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetConvertOrder(ctx, req)
	if err != nil {
		log.Printf("call TestGetConvertOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}
