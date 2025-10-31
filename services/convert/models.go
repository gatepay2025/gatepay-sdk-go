package convert

import "github.com/gatepay2025/gatepay-sdk-go/services/common"

type ConvertSupportCurrencyRequest struct {
	common.BaseRequest
	Side string `json:"side"`
}

type ConvertSupportCurrencyResponse struct {
	Currency []string `json:"currency"`
}

type ConvertPairRequest struct {
	common.BaseRequest
	Side     string `query:"side"`
	Currency string `query:"currency"`
}

type ConvertSupportCurrencyPairResponse struct {
	Pair            string `json:"pair"`
	SellCurrency    string `json:"sellCurrency"`
	SellCurrencyMax string `json:"sellCurrencyMax"`
	SellCurrencyMin string `json:"sellCurrencyMin"`
	BuyCurrency     string `json:"buyCurrency"`
	BuyCurrencyMax  string `json:"buyCurrencyMax"`
	BuyCurrencyMin  string `json:"buyCurrencyMin"`
}

type ConvertPreviewRequest struct {
	common.BaseRequest
	BuyCurrency  string `json:"buyCurrency"`
	BuyAmount    string `json:"buyAmount" `
	SellCurrency string `json:"sellCurrency"`
	SellAmount   string `json:"sellAmount" `
}

type ConvertPreviewResponse struct {
	SellCurrency string `json:"sellCurrency"`
	BuyCurrency  string `json:"buyCurrency"`
	BuyAmount    string `json:"buyAmount"`
	SellAmount   string `json:"sellAmount"`
	Rate         string `json:"price"`
	QuoteId      string `json:"quoteId"`
}

type ConvertOrderRequest struct {
	common.BaseRequest
	ClientReqId  string `json:"clientReqId"`
	QuoteId      string `json:"quoteId"`
	BuyAmount    string `json:"buyAmount"`
	BuyCurrency  string `json:"buyCurrency"`
	SellAmount   string `json:"sellAmount"`
	SellCurrency string `json:"sellCurrency"`
}

type ConvertOrder struct {
	OrderID      string  `json:"order_id"`
	UserID       int64   `json:"userId"`
	SellCurrency string  `json:"sellCurrency"`
	BuyCurrency  string  `json:"buyCurrency"`
	SellAmount   string  `json:"sellAmount"`
	BuyAmount    string  `json:"buyAmount"`
	Status       int64   `json:"status"`
	Rate         float64 `json:"rate"`
	QuoteID      string  `json:"quoteId"`
	CreateTime   int64   `json:"createTime"`
}

type ConvertOrderResponse struct {
	ConvertOrder
}

type QueryConvertOrderReq struct {
	common.BaseRequest
	OrderID string `query:"orderId"`
}

type QueryConvertOrderResponse struct {
	ConvertOrder
}
