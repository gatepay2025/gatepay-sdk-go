package payment

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/core/stringutillib"
	"log"
	"testing"
)

func TestChains(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &PayApiService{Client: client}

	req := OperateOrderRequest{PrepayID: "prepay-order-id", MerchantTradeNo: ""}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetOrder(ctx, req)
	if err != nil {
		log.Printf("call GetAddressChains err:%s", err)
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}
