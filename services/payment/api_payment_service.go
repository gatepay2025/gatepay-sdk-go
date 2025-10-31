package payment

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/services"
	nethttp "net/http"
	neturl "net/url"
	"strconv"
)

type PayApiService services.Service

// /v1/pay/order/query
func (a *PayApiService) GetOrder(ctx context.Context, req OperateOrderRequest) (resp *QueryOrderResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/order/query"
	localVarPostBody = req

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("prepayId", req.PrepayID)
	localVarQueryParams.Add("merchantTradeNo", req.MerchantTradeNo)
	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &QueryOrderResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil

}

// /v1/pay/order 创建非地址支付支付订单
func (a *PayApiService) CreateOrder(ctx context.Context, req CreateOrderRequest) (resp *CreateOrderResponse, result *core.APIResult, err error) {
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

	localVarPath := core.DefaultEndpoint + "/v1/pay/order"
	localVarPostBody = req
	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &CreateOrderResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/order/close 主动关闭订单
func (a *PayApiService) CloseOrder(ctx context.Context, req OperateOrderRequest) (resp *OperateOrderResponse, result *core.APIResult, err error) {
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

	localVarPath := core.DefaultEndpoint + "/v1/pay/order/close"
	localVarPostBody = req
	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &OperateOrderResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/order/refund 创建退款单
func (a *PayApiService) CreateRefundOder(ctx context.Context, req CreateRefundRequest) (resp *CreateRefundResponse, result *core.APIResult, err error) {
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

	localVarPath := core.DefaultEndpoint + "/v1/pay/order/refund"
	localVarPostBody = req
	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &CreateRefundResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/order/refund/query
func (a *PayApiService) QueryRefund(ctx context.Context, req QueryRefundRequest) (resp *QueryRefundResponse, result *core.APIResult, err error) {
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

	localVarPath := core.DefaultEndpoint + "/v1/pay/order/refund/query"
	localVarPostBody = req
	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &QueryRefundResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/batch/transfer 批量转账
func (a *PayApiService) CreateTransfer(ctx context.Context, req BatchTransfer) (resp *BatchTransferResponse, result *core.APIResult, err error) {
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

	localVarPath := core.DefaultEndpoint + "/v1/pay/batch/transfer"
	localVarPostBody = req
	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &BatchTransferResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/batch/transfer/query
func (a *PayApiService) BatchTransferQuery(ctx context.Context, req BatchTransferQueryReq) (resp *BatchTransferQueryResponse, result *core.APIResult, err error) {
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

	localVarPath := core.DefaultEndpoint + "/v1/pay/batch/transfer/query"
	localVarPostBody = req
	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &BatchTransferQueryResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/transactions/native
func (a *PayApiService) CreateNativeOrder(ctx context.Context, req CreateOrderRequest) (resp *QrOrderResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarPath := core.DefaultEndpoint + "/v1/pay/transactions/native"
	localVarPostBody = req

	localVarHTTPContentTypes := core.ApplicationJSON

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &QrOrderResponse{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// v1/pay/bill/orderlist
func (a *PayApiService) GetOrderList(ctx context.Context, req BalanceHistoryListReq) (resp *BalanceHistoryListResp, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/bill/orderlist"
	localVarQueryParams = neturl.Values{}

	tt := req.TransactionType
	localVarQueryParams.Add("transactionType", strconv.Itoa(tt))
	timeType := req.TimeType
	localVarQueryParams.Add("timeType", strconv.FormatInt(timeType, 10))
	st := strconv.FormatInt(req.StartTime, 10)
	localVarQueryParams.Add("startTime", st)
	et := strconv.FormatInt(req.EndTime, 10)
	localVarQueryParams.Add("endTime", et)
	ct := strconv.Itoa(req.Count)
	localVarQueryParams.Add("count", ct)
	pg := strconv.Itoa(req.Page)
	localVarQueryParams.Add("page", pg)
	status := req.Status
	if status != "" {
		localVarQueryParams.Add("status", status)
	}
	cur := req.Currency
	if cur != "" {
		localVarQueryParams.Add("currency", cur)
	}
	ot := req.OrderType
	localVarQueryParams.Add("orderType", strconv.Itoa(ot))
	orderIdNo := req.OrderIdNo
	if orderIdNo != "" {
		localVarQueryParams.Add("orderIdNo", orderIdNo)
	}
	financialType := req.FinancialType
	if financialType != "" {
		localVarQueryParams.Add("financialType", financialType)
	}
	sortType := req.SortType
	if sortType != "" {
		localVarQueryParams.Add("sortType", sortType)
	}

	localVarHTTPContentTypes := core.ApplicationJSON

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &BalanceHistoryListResp{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
