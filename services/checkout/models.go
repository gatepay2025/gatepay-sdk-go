package checkout

import (
	"github.com/gatepay2025/gatepay-sdk-go/services/common"
	"github.com/shopspring/decimal"
)

type EnvRequest struct {
	TerminalType string `json:"terminalType" validator:"nonemptyString"`
	Scene        string `json:"scene" validator:"scene"`
}

type GoodsRequest struct {
	GoodsName   string `json:"goodsName"`
	GoodsDetail string `json:"goodsDetail"`
}

type CreateOrderRequest struct {
	common.BaseRequest
	MerchantTradeNo string          `json:"merchantTradeNo"`
	Currency        string          `json:"currency"`
	OrderAmount     decimal.Decimal `json:"orderAmount"`
	PayCurrency     string          `json:"payCurrency"`
	ActualCurrency  string          `json:"actualCurrency"`
	Env             EnvRequest      `json:"env"`
	Goods           GoodsRequest    `json:"goods"`
	OrderExpireTime int64           `json:"orderExpireTime"`
	ReturnURL       string          `json:"returnUrl"`
	CancelURL       string          `json:"cancelUrl"`
	MerchantUserId  int64           `json:"merchantUserId"`
	Chain           string          `json:"chain"`
	FullCurrType    string          `json:"fullCurrType"`
	ChannelId       string          `json:"channelId"`
}

type Chain struct {
	ChainType    string `json:"chain_type"`
	Address      string `json:"address"`
	FullCurrType string `json:"fullCurrType,omitempty"`
}

type CreateOrderResponseV2 struct {
	PrepayID     string `json:"prepayId"`
	TerminalType string `json:"terminalType"`
	ExpireTime   int64  `json:"expireTime"`
	QrContent    string `json:"qrContent"`
	Location     string `json:"location"`
	PayCurrency  string `json:"payCurrency"`
	PayAmount    string `json:"payAmount"`
	ChainInfo    Chain  `json:"chain"`
	AppName      string `json:"appName"`
	AppLogo      string `json:"appLogo"`
	GoodsName    string `json:"goodsName"`
	InUsdt       string `json:"inUsdt"`
}

type CreateCheckStandRefundResponse struct {
	RefundRequestID   string          `json:"refundRequestId"`
	PrepayID          string          `json:"prepayId"`
	OrderCurrency     string          `json:"orderCurrency"`
	OrderAmount       decimal.Decimal `json:"orderAmount"`
	RefundOrderAmount decimal.Decimal `json:"refundOrderAmount"`
	PayCurrency       string          `json:"payCurrency"`
	PayAmount         decimal.Decimal `json:"payAmount"`
	RefundPayAmount   decimal.Decimal `json:"refundPayAmount"`
}

type CreateCheckStandRefundRequest struct {
	common.BaseRequest
	RefundRequestID     string          `json:"refundRequestId"`
	PrepayID            string          `json:"prepayId"`
	RefundOrderCurrency string          `json:"refundOrderCurrency"`
	RefundOrderAmount   decimal.Decimal `json:"refundOrderAmount"`
	RefundPayCurrency   string          `json:"refundPayCurrency"`
	RefundPayAmount     decimal.Decimal `json:"refundPayAmount"`
	RefundReason        string          `json:"refundReason"`
	ReceiverId          int64           `json:"receiverId"`
}
