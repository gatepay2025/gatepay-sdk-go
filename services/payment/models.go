package payment

import (
	"github.com/gatepay2025/gatepay-sdk-go/services/common"
	"github.com/shopspring/decimal"
)

type DetailStatus string

type OperateOrderRequest struct {
	common.BaseRequest
	PrepayID        string `json:"prepayId"`
	MerchantTradeNo string `json:"merchantTradeNo"`
}

type QueryOrderResponse struct {
	PrepayID        string `json:"prepayId"`
	MerchantID      int64  `json:"merchantId"`
	MerchantTradeNo string `json:"merchantTradeNo"`
	TransactionID   string `json:"transactionId"`
	GoodsName       string `json:"goodsName"`
	Currency        string `json:"currency"`
	OrderAmount     string `json:"orderAmount"`
	Status          string `json:"status"`
	CreateTime      int64  `json:"createTime"`
	ExpireTime      int64  `json:"expireTime"`
	TransactTime    int64  `json:"transactTime"`
	OrderName       string `json:"order_name"`
	PayCurrency     string `json:"pay_currency"`
	PayAmount       string `json:"pay_amount"`
	ExpectCurrency  string `json:"expectCurrency,omitempty"`
	ActualCurrency  string `json:"actualCurrency,omitempty"`
	ActualAmount    string `json:"actualAmount,omitempty"`
	Rate            string `json:"rate"`
	PayChan         string `json:"channel_type"`
	PayAccount      string `json:"pay_account"`
	AppName         string `json:"appName"`
	AppLogo         string `json:"appLogo"`
	InUsdt          string `json:"inUsdt"`
	ChannelId       string `json:"channelId"`
}

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
	Currency        string          `json:"currency"`       // order_currency
	OrderAmount     decimal.Decimal `json:"orderAmount"`    // order_amount default zero
	PayCurrency     string          `json:"payCurrency"`    // pay_currency 非地址支付PayCurrency 在实际付款确定 ，地址支付在下单时候确定
	ActualCurrency  string          `json:"actualCurrency"` // merchant actual currency
	Env             EnvRequest      `json:"env"`
	Goods           GoodsRequest    `json:"goods"`
	OrderExpireTime int64           `json:"orderExpireTime"`
	ReturnURL       string          `json:"returnUrl"`
	CancelURL       string          `json:"cancelUrl"`
	MerchantUserId  int64           `json:"merchantUserId"` // userId in merchants' system
	Chain           string          `json:"chain"`
	FullCurrType    string          `json:"fullCurrType"`
	ChannelId       string          `json:"channelId"` // 客户渠道名称
}

type CreateOrderResponse struct {
	PrepayID     string `json:"prepayId"`
	TerminalType string `json:"terminalType"`
	ExpireTime   int64  `json:"expireTime"`
}

type OperateOrderResponse struct {
	Result string `json:"result"`
}

type CreateRefundRequest struct {
	common.BaseRequest
	RefundRequestID string          `json:"refundRequestId"`
	PrepayID        string          `json:"prepayId"`
	RefundAmount    decimal.Decimal `json:"refundAmount"`
	RefundReason    string          `json:"refundReason"`
	RefundGateId    string          `json:"RefundGateId"`
}

type CreateRefundResponse struct {
	RefundRequestID string `json:"refundRequestId"`
	PrepayID        string `json:"prepayId"`
	OrderAmount     string `json:"orderAmount"`
	RefundAmount    string `json:"refundAmount"`
}

type QueryRefundRequest struct {
	common.BaseRequest
	RefundRequestID string `json:"refundRequestId"`
}

type QueryRefundResponse struct {
	RefundEntry
}

type RefundEntry struct {
	RefundRequestID string `json:"refundRequestId"`
	PrepayID        string `json:"prepayId"`
	OrderAmount     string `json:"orderAmount"`
	RefundAmount    string `json:"refundAmount"`
	RefundStatus    string `json:"refundStatus"`
	ChannelId       string `json:"channelId"`
}

type BizSceneType string

const (
	BizDirectTransfer  BizSceneType = "DIRECT_TRANSFER" // 支付付款
	BizRewards                      = "REWARDS"         //奖励
	BizReimburseMent                = "REIMBURSEMENT"   //退款
	BizMerchantPayment              = "MERCHANTPAYMENT" //商户付款
	BizOthersPayment                = "OTHERSPAYMENT"   //其他支付
	BizBatchGiftCard                = "BATCHGIFTCARD"   // 批量创建礼品卡
	BizGiftExchange                 = "GIFTEXCHANGE"    // 礼品卡兑现
)

type BatchOrderList struct {
	UserId   int64  `json:"user_id"`
	Amount   string `json:"amount"`
	RewardId string
}

type BatchTransfer struct {
	common.BaseRequest
	BatchId         string            `json:"batchid"`
	MerchantBatchNo string            `json:"merchant_batch_no" query:"merchant_batch_no"`
	BizScene        BizSceneType      `json:"bizscene" query:"bizscene"`
	MerchantID      string            `json:"merchant_id"`
	ClientId        string            `json:"client_id"`
	Currency        string            `json:"currency" query:"currency"`
	Name            string            `json:"name " query:"name"`
	Description     string            `json:"description " query:"description"`
	OrderList       []*BatchOrderList `json:"batchorderList" query:"batchorderList"`
	ChannelId       string            `json:"channelId" query:"channelId"`
}

type BatchTransferResponse struct {
	MerchantBatchNo string `json:"merchant_batch_no"`
	BatchId         string `json:"batch_id"`
}

type BatchTransferQueryReq struct {
	common.BaseRequest
	BatchId         string       `json:"batch_id"`
	MerchantBatchNo string       `json:"merchant_batch_no"`
	Status          DetailStatus `json:"detail_status"`
}

type BatchTransferQueryResponse struct {
	BatchId         string                  `json:"batch_id"`
	MerchantId      int64                   `json:"merchant_id"`
	MerchantBatchNo string                  `json:"merchant_batch_no"`
	Status          string                  `json:"status"`
	Currency        string                  `json:"currency,omitempty"`
	ChannelId       string                  `json:"channel_id"`
	OrdersList      []*BatchOrderListNotify `json:"orders_list"`
}

type BatchOrderListNotify struct {
	UserId        int64  `gorm:"column:receiver_id" json:"receiver_id"`
	Amount        string `gorm:"column:amount" json:"amount"`
	Currency      string `gorm:"column:currency" json:"currency,omitempty"`
	Status        string `gorm:"column:status" json:"status"`
	RewardId      string `gorm:"column:reward_id" json:"reward_id"`
	TransactionId string `gorm:"column:transaction_id" json:"transaction_id"`
	CreateTime    int64  `gorm:"column:create_time" json:"create_time"`
	ChannelId     string `gorm:"column:channel_id" json:"channel_id"`
}

type QrOrderResponse struct {
	PrepayID     string `json:"prepayId"`
	TerminalType string `json:"terminalType"`
	ExpireTime   int64  `json:"expireTime"`
	QrContent    string `json:"qrContent"`
	Location     string `json:"location"`
}

type BalanceHistoryItem struct {
	TransactionType int    `json:"transactionType"`
	TimeType        int64  `json:"timeType"`
	StartTime       int64  `json:"startTime"`
	EndTime         int64  `json:"endTime"`
	Count           int    `json:"count"`
	Page            int    `json:"page"`
	Status          string `json:"status"`
	Currency        string `json:"currency"`
	OrderIdNo       string `json:"orderIdNo"`
	FinancialType   string `json:"financialType"`
	SortType        string `json:"sortType"`
	OrderType       int    `json:"orderType"`
	TransactID      string `json:"transactId"`
	TransactTime    int64  `json:"transactTime"`
	OrderId         string `json:"orderId"`
	MerchantTradeNo string `json:"merchantTradeNo"`
	PayAmount       string `json:"payAmount"`
	Balance         string `json:"Balance"`
	BalanceCurrency string `json:"BalanceCurrency"`
	Payer           int64  `json:"payer"`
	Buyer           string `json:"buyer"`
	FullChain       string `json:"fullChain"`
	Hash            string `json:"hash"`
	Address         string `json:"address"`
	PayChannel      string `json:"payChannel"`
	BillType        uint8  `json:"billType"`
	RefundGateId    string `json:"refund_gate_id"`
	Business        string `json:"business"`
}

type BalanceHistoryListReq struct {
	common.BaseRequest
	TransactionType int    `json:"transactionType"`
	TimeType        int64  `json:"timeType"`
	StartTime       int64  `json:"startTime"`
	EndTime         int64  `json:"endTime"`
	Count           int    `json:"count"`
	Page            int    `json:"page"`
	Status          string `json:"status"`
	Currency        string `json:"currency"`
	OrderType       int    `json:"orderType"`
	OrderIdNo       string `json:"orderIdNo"`
	FinancialType   string `json:"financialType"`
	SortType        string `json:"sortType"`
}

type BalanceHistoryListResp struct {
	MerchantId             int64                 `json:"merchant_id"`
	Total                  int64                 `json:"total"`
	HasNext                bool                  `json:"hasNext"`
	NextPage               int                   `json:"nextPage"`
	BalanceHistoryItemList []*BalanceHistoryItem `json:"balance_history_item_list"`
}
