package channelmanage

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/core/stringutillib"
	"log"
	"testing"
)

// /v1/pay/channelmanage/save
func TestChannelManageSave(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &ApiChannelManagementService{Client: client}

	mc := MerchantChannel{
		ChannelId:    "abccc123",
		Desc:         "astratest123111",
		ChannelType:  "1",
		Address:      "1234",
		Currency:     "USDT",
		Chain:        "ETH2",
		CustomFields: []*CustomField{},
	}
	req := MerchantChannelReq{
		ChannelId:   "astratest1230aastratest1230aastratest1230aastrates123",
		ChannelType: "1",
	}
	req.MerchantChannelList = append(req.MerchantChannelList, &mc)
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")

	resp, result, err := service.ChannelManageSave(ctx, req)
	if err != nil {
		log.Printf("call TestChannelManageSave err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/channelmanage/list
func TestChannelManageList(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &ApiChannelManagementService{Client: client}
	req := MerchantChannelReq{
		Page:  1,
		Count: 2,
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")

	resp, result, err := service.ChannelManageList(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/channelmanage/update
func TestChannelManageUpdate(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &ApiChannelManagementService{Client: client}
	mc := MerchantChannel{
		ChannelId:    "abccc123",
		Desc:         "astratest123111",
		ChannelType:  "1",
		Address:      "1234",
		Currency:     "USDT",
		Chain:        "ETH2",
		CustomFields: []*CustomField{},
	}
	req := MerchantChannelReq{
		ChannelId:   "astratest1230aastratest1230aastratest1230aastrates123",
		ChannelType: "1",
	}
	req.MerchantChannelList = append(req.MerchantChannelList, &mc)
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")

	resp, result, err := service.ChannelManageUpdate(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}

// /v1/pay/channelmanage/delete
func TestChannelManageDelete(t *testing.T) {
	cfg := core.NewConfig()
	credentials := core.NewCredentials("Mz6M_q4AkDnZCSoTDo03A6OtWzN5ut8_Uix3jyVjxAU=")
	client, err := core.NewClient(cfg, credentials)
	if err != nil {
		return
	}

	ctx := context.Background()
	service := &ApiChannelManagementService{Client: client}
	req := MerchantChannelReq{
		ChannelId: "astra1229",
	}
	req.AddHeader("X-GatePay-Certificate-ClientId", "mZ96D37oKk-HrWJc")

	resp, result, err := service.ChannelManageDelete(ctx, req)
	if err != nil {
		log.Printf("call QueryAddressOrder err:%s", err.Error())
	} else {
		log.Printf("status=%d resp=%v", result.Response.StatusCode, stringutillib.ObjToJsonStr(resp))
	}
}
