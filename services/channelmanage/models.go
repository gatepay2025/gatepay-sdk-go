package channelmanage

import "github.com/gatepay2025/gatepay-sdk-go/services/common"

var emptyData = map[string]string{}

type CustomField struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MerchantChannelReq struct {
	common.BaseRequest
	ChannelId           string             `json:"channelId"`
	Desc                string             `json:"desc"`
	ChannelType         string             `json:"channelType"`
	MerchantChannelList []*MerchantChannel `json:"merchantChannelList"`
	Page                int                `json:"page"`  //开始页编号
	Count               int                `json:"count"` //查询记录数
}

type MerchantChannel struct {
	ChannelId    string         `json:"channelId"`
	Desc         string         `json:"desc"`
	ChannelType  string         `json:"channelType"`
	Chain        string         `json:"chain"`
	Currency     string         `json:"currency"`
	Address      string         `json:"address"`
	CreateTime   string         `json:"createTime"`
	UpdateTime   string         `json:"updateTime"`
	CustomFields []*CustomField `json:"customFields"`
	Result       string         `json:"result"`
}

type MerchantChannelSearchRes struct {
	Total               int64              `json:"total"`
	MerchantChannelList []*MerchantChannel `json:"merchantChannelList"`
}

type MerchantChannelUpdateReq struct {
	common.BaseRequest
	ChannelId           string             `json:"channelId"`
	Desc                string             `json:"desc"`
	ChannelType         string             `json:"channelType"`
	MerchantChannelList []*MerchantChannel `json:"merchantChannelList"`
	Page                int                `json:"page"`
	Count               int                `json:"count"`
}
