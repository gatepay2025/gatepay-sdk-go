package checkout

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/services"
	nethttp "net/http"
	neturl "net/url"
)

type ApiCheckOutService services.Service

// /v1/pay/checkout/order 创建收银台订单
func (a *ApiCheckOutService) CreateCheckOutOrder(ctx context.Context, req CreateOrderRequest) (resp *CreateOrderResponseV2, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/checkout/order"
	localVarPostBody = req
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &CreateOrderResponseV2{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}

// /v1/pay/checkout/refund
func (a *ApiCheckOutService) CreateCheckoutRefund(ctx context.Context, req CreateCheckStandRefundRequest) (resp *CreateCheckStandRefundResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/checkout/refund"
	localVarPostBody = req

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &CreateCheckStandRefundResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}
