package gift

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/services"
	"github.com/gatepay2025/gatepay-sdk-go/services/common"
	nethttp "net/http"
	neturl "net/url"
)

type ApiGiftService services.Service

// /v1/pay/gift/create
func (a *ApiGiftService) CreateGiftCard(ctx context.Context, req MerchantGiftCreateReq) (resp *GiftCardResp, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/gift/create"
	localVarPostBody = req

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &GiftCardResp{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}

// /v1/pay/gift/temp/list
func (a *ApiGiftService) GetGiftTempList(ctx context.Context, req common.BaseRequest) (resp []TempListResp, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/gift/temp/list"
	localVarQueryParams = neturl.Values{}

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = []TempListResp{}
	err = core.UnMarshalResponse(result.Response, &resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}

// /v1/pay/balance
func (a *ApiGiftService) GetPayBalance(ctx context.Context, req common.BaseRequest) (resp map[string]string, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/balance"
	localVarQueryParams = neturl.Values{}

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = make(map[string]string)

	//result转resp
	err = core.UnMarshalResponse(result.Response, &resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/gift/query
func (a *ApiGiftService) GetGiftQuery(ctx context.Context, req MerchantGiftQueryReq) (resp *MerchantGiftQueryResp, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/gift/query"
	localVarQueryParams = neturl.Values{}
	localVarPostBody = req
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &MerchantGiftQueryResp{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
