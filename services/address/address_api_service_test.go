package address

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/core/stringutillib"
	"github.com/gatepay2025/gatepay-sdk-go/services/common"
	"github.com/shopspring/decimal"
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
	service := &AddressApiService{Client: client}

	req := ChainsRequest{Currency: "USDT", BaseRequest: &common.BaseRequest{}}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")

	resp, result, err := service.GetAddressChains(ctx, req)
	if err != nil {
		log.Printf("call GetAddressChains err:%s", err)
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleChains() {
	//Create Configuration
	cfg := core.NewConfig()
	//Setting up the payment secret key
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}

	//create request parameters
	req := ChainsRequest{Currency: "USDT"}
	//Call the Get Address Chain Information function
	resp, result, err := service.GetAddressChains(ctx, req)
	//Processing Response
	if err != nil {
		log.Printf("call GetAddressChains err:%s", err)
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleCreateAddressOrder() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}

	req := CreateOrderRequest{
		MerchantTradeNo: "merchant-trade-no",
		OrderExpireTime: 1750834613000,
		OrderAmount:     decimal.NewFromInt(1),
		Currency:        "USDT",
		CancelURL:       "cancel-url",
		ChannelId:       "",

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
	resp, result, err := service.CreateAddress(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestCreateAddressOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}

	orderAmount, _ := decimal.NewFromString("1.2")

	req := CreateOrderRequest{
		MerchantTradeNo: "merchant-trade-no",
		OrderExpireTime: 1762710400000,
		OrderAmount:     orderAmount,
		Currency:        "USDT",
		CancelURL:       "cancel-url",
		ChannelId:       "",

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
	resp, result, err := service.CreateAddress(ctx, req)
	if err != nil {
		log.Printf("call CreateAddress err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleQueryAddressOrder() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}
	req := QueryAddressOrderRequest{MerchantTradeNo: "merchant-trade-no", PrepayID: "order-id"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")

	resp, result, err := service.QueryAddressOrder(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestQueryAddressOrder(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}
	req := QueryAddressOrderRequest{MerchantTradeNo: "merchant-trade-no", PrepayID: "order-id"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")

	resp, result, err := service.QueryAddressOrder(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestGetAddressCurrencies(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}
	//接口没有参数，可以传common.BaseRequest,这样可以在BaseRequest传入client_id用于签名
	req := common.BaseRequest{}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")

	resp, result, err := service.GetAddressCurrencies(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleAddressCurrencies() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}
	//接口没有参数，可以传common.BaseRequest,这样可以在BaseRequest传入client_id用于签名
	req := common.BaseRequest{}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")

	resp, result, err := service.GetAddressCurrencies(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestSupportedConvertCurrencies(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}

	req := SupportedConvertCurrenciesReq{Currency: "ETH"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetAddressSupportedConvertCurrencies(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func ExampleSupportedConvertCurrencies() {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}

	req := SupportedConvertCurrenciesReq{Currency: "ETH"}
	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.GetAddressSupportedConvertCurrencies(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

func TestQueryAddressTransactionDetail(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("secret-key")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &AddressApiService{Client: client}
	req := TransactionDetailReq{PrepayId: "order-id"}

	req.AddHeader("X-GatePay-Certificate-ClientId", "client-id")
	resp, result, err := service.QueryAddressTransactionDetail(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}
