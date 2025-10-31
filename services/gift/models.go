package gift

import (
	"github.com/gatepay2025/gatepay-sdk-go/services/common"
	"github.com/shopspring/decimal"
)

type MerchantGiftCreateReq struct {
	common.BaseRequest
	Title      string `json:"title" validate:"required"`
	TemplateId string `json:"templateId" validate:"required"`
	Currency   string `json:"currency" validate:"required"`
	Amount     string `json:"amount" validate:"required"` // 单卡面值
	Count      int64  `json:"count"`
}

type GiftCardResp struct {
	CardNum      string          `json:"card_num"`
	CardKey      string          `json:"card_key"`
	Amount       decimal.Decimal `json:"amount"`
	Currency     string          `json:"currency"`
	Status       int             `json:"status"`
	CardTempId   string          `json:"card_temp_id"`
	Creator      int64           `json:"creator"`
	CreatorName  string          `json:"creator_name"`
	ExchangeUid  int64           `json:"exchange_uid"`
	Owner        int64           `json:"owner"`
	GiveCount    int             `json:"give_count"`
	LastGiveTime int64           `json:"last_give_time"`
	CreateTime   int64           `json:"create_time"`
	BatchId      string          `json:"batchId"`
	Reason       string          `json:"reason"`
	ErrMsg       string          `json:"err_msg"`
}

type TempListResp struct {
	CardTempId string `json:"card_temp_id"`
	ImageUrl   string `json:"image_url"`
	TitleEn    string `json:"title_en"`
	TitleCn    string `json:"title_cn"`
	CoverType  string `json:"cover_type"`
}

type MerchantGiftQueryReq struct {
	common.BaseRequest
	CardNumber string `json:"card_number"`
	Key        string `json:"key"`
}

type MerchantGiftQueryResp struct {
	CardNum      string `json:"card_num"`
	Key          string `json:"key"`
	Title        string `json:"title"`
	CreatorName  string `json:"creator_name"`
	Amount       string `json:"amount"`
	Status       int    `json:"status"`
	Currency     string `json:"currency"`
	CreateTime   int64  `json:"create_time"`
	ExchangeUid  int64  `json:"exchange_uid"`
	ExchangeTime int64  `json:"exchange_time"`
}
