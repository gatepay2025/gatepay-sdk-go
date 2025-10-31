package channelmanage

import (
	"context"
	"github.com/gatepay2025/gatepay-sdk-go/core"
	"github.com/gatepay2025/gatepay-sdk-go/services"
	nethttp "net/http"
	neturl "net/url"
	"strconv"
)

type ApiChannelManagementService services.Service

// /v1/pay/channelmanage/save 新增客户
func (a *ApiChannelManagementService) ChannelManageSave(ctx context.Context, req MerchantChannelReq) (resp map[string]string, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/channelmanage/save"
	localVarPostBody = req

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = emptyData
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}

// /v1/pay/channelmanage/list
func (a *ApiChannelManagementService) ChannelManageList(ctx context.Context, req MerchantChannelReq) (resp *MerchantChannelSearchRes, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/channelmanage/list"
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("channelId", req.ChannelId)
	localVarQueryParams.Add("desc", req.Desc)
	localVarQueryParams.Add("page", strconv.Itoa(req.Page))
	localVarQueryParams.Add("Count", strconv.Itoa(req.Count))
	localVarQueryParams.Add("channelType", req.ChannelType)
	localVarHTTPContentTypes := core.ApplicationJSON

	//set用户设置的Header
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = &MerchantChannelSearchRes{}
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// /v1/pay/channelmanage/update
func (a *ApiChannelManagementService) ChannelManageUpdate(ctx context.Context, req MerchantChannelReq) (resp map[string]string, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPut
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/channelmanage/update"
	localVarPostBody = req

	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = emptyData
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}

// /v1/pay/channelmanage/delete
func (a *ApiChannelManagementService) ChannelManageDelete(ctx context.Context, req MerchantChannelReq) (resp map[string]string, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	localVarPath := core.DefaultEndpoint + "/v1/pay/channelmanage/delete"
	localVarPostBody = req
	for k, v := range req.GetHeaders() {
		localVarHeaderParams.Set(k, v)
	}

	localVarHTTPContentTypes := core.ApplicationJSON
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentTypes)
	if err != nil {
		return nil, result, err
	}

	resp = emptyData
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}

	return resp, result, nil
}
