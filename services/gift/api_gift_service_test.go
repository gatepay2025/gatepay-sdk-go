package gift

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/core/stringutillib"
	"github.com/gatepay2025/gatepay-sdk-go/services/common"
	"log"
	"testing"
)

// /v1/pay/gift/create
func TestCreateGiftOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiGiftService{Client: client}

	req := MerchantGiftCreateReq{
		Title:      "title",
		TemplateId: "template-id",
		Currency:   "USDT",
		Amount:     "1",
		Count:      1,
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.CreateGiftCard(ctx, req)
	if err != nil {
		log.Printf("call TestTransferOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}

}

func ExampleCreateGiftOrder() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiGiftService{Client: client}

	req := MerchantGiftCreateReq{
		Title:      "title",
		TemplateId: "template-id",
		Currency:   "USDT",
		Amount:     "1",
		Count:      1,
	}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.CreateGiftCard(ctx, req)
	if err != nil {
		log.Printf("call TestTransferOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/gift/query
func TestGiftQuery(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiGiftService{Client: client}

	req := MerchantGiftQueryReq{
		CardNumber: "card-number",
		Key:        "",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetGiftQuery(ctx, req)
	if err != nil {
		log.Printf("call TestGiftQuery err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleGiftQuery() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiGiftService{Client: client}

	req := MerchantGiftQueryReq{
		CardNumber: "card-number",
		Key:        "",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetGiftQuery(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/gift/temp/list
func TestGiftTempList(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiGiftService{Client: client}

	req := common.BaseRequest{}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetGiftTempList(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleGiftTempList() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiGiftService{Client: client}

	req := common.BaseRequest{}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetGiftTempList(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/balance
func TestPayBalance(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiGiftService{Client: client}

	req := common.BaseRequest{}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetPayBalance(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExamplePayBalance() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &ApiGiftService{Client: client}

	req := common.BaseRequest{}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetPayBalance(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}
