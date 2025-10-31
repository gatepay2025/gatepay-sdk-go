package withdraw

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/services"
	nethttp "net/http"
	neturl "net/url"
)

type APIWithdrawService services.Service

// /v1/pay/withdraw
func (a *APIWithdrawService) WithdrawCreateOrder(ctx context.Context, req BatchWithdrawRequest) (resp *BatchWithdrawResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/withdraw"
	localVarPostBody = req

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &BatchWithdrawResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}

// /v1/pay/withdraw/query
func (a *APIWithdrawService) WithdrawQuery(ctx context.Context, req *BatchWithdrawQueryReq) (resp *BatchWithdrawQueryResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/withdraw/query"
	localVarPostBody = req

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &BatchWithdrawQueryResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}

// /v1/pay/wallet/currency_chains
func (a *APIWithdrawService) GetCurrencyChains(ctx context.Context, req WalletCurrency) (resp *[]CurrencyChainItem, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/wallet/currency_chains"
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("currency", req.Currency)

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &[]CurrencyChainItem{}
	err = core.UnMarshalResponseWithdraw(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}

// /v1/pay/wallet/total_balance
func (a *APIWithdrawService) GetTotalBalance(ctx context.Context, req WalletCurrency) (resp *TotalBanlanceResp, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/wallet/total_balance"
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("currency", req.Currency)

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &TotalBanlanceResp{}
	err = core.UnMarshalResponseWithdraw(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}

// /v1/pay/wallet/withdraw_status  //非GatePay接口
func (a *APIWithdrawService) GetWithdrawStatus(ctx context.Context, req WalletCurrency) (resp *[]WithdrawStatusItem, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/wallet/withdraw_status"
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("currency", req.Currency)

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &[]WithdrawStatusItem{}
	err = core.UnMarshalResponseWithdraw(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}
