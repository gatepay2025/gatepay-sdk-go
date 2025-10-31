package convert

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/services"
	nethttp "net/http"
	neturl "net/url"
)

type ApiConvertService services.Service

// /v1/pay/convert/currency
func (a *ApiConvertService) GetConvertCurrency(ctx context.Context, req ConvertSupportCurrencyRequest) (resp *ConvertSupportCurrencyResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/convert/currency"
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("side", req.Side)
	localVarHTTPContentTypes := core.ApplicationJSON

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &ConvertSupportCurrencyResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/convert/pair
func (a *ApiConvertService) GetConvertPair(ctx context.Context, req ConvertPairRequest) (resp *[]ConvertSupportCurrencyPairResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/convert/pair"
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("side", req.Side)
	localVarQueryParams.Add("currency", req.Currency)
	localVarHTTPContentTypes := core.ApplicationJSON

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &[]ConvertSupportCurrencyPairResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/convert/preview
func (a *ApiConvertService) ConvertPreView(ctx context.Context, req ConvertPreviewRequest) (resp *ConvertPreviewResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarPath := core.DefaultEndpoint + "/v1/pay/convert/preview"
	localVarPostBody = req
	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &ConvertPreviewResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/convert
func (a *ApiConvertService) PayConvert(ctx context.Context, req ConvertOrderRequest) (resp *ConvertOrderResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarPath := core.DefaultEndpoint + "/v1/pay/convert"
	localVarPostBody = req
	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &ConvertOrderResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil

}

// /v1/pay/convert/order
func (a *ApiConvertService) GetConvertOrder(ctx context.Context, req QueryConvertOrderReq) (resp *QueryConvertOrderResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/convert/order"
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("orderId", req.OrderID)
	localVarHTTPContentTypes := core.ApplicationJSON

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &QueryConvertOrderResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
