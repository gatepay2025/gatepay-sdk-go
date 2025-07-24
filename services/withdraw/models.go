package withdraw

import (
	"github.com/gatepay2025/gatepay-sdk-go/services/common"
	"github.com/shopspring/decimal"
)

type BatchWithdrawRequest struct {
	common.BaseRequest
	BatchID      string          `json:"batch_id"`
	WithdrawList []*WithdrawItem `json:"withdraw_list"`
	ChannelId    string          `json:"channel_id"`
}

type WithdrawItem struct {
	MerchantWithdrawId string `json:"merchant_withdraw_id"`
	Currency           string `json:"currency"`
	Amount             string `json:"amount"`
	Chain              string `json:"chain"`
	Address            string `json:"address"`
	Memo               string `json:"memo"`
	FeeType            int8   `json:"fee_type"`
}

type BatchWithdrawResponse struct {
	BatchId string `json:"batch_id"` //由GatePay侧生成
}

type BatchWithdrawQueryReq struct {
	common.BaseRequest
	BatchId string `json:"batch_id"`
	Status  string `json:"detail_status"`
}

type BatchWithdrawQueryResponse struct {
	BatchId      string                       `json:"batch_id"`
	MerchantId   int64                        `json:"merchant_id"`
	ClientID     string                       `json:"client_id"`
	Status       string                       `json:"status"`        //总单的状态
	CreateTime   int64                        `json:"create_time"`   //总单创建时间
	WithdrawList []*BatchWithdrawSuborderItem `json:"withdraw_list"` //子单信息
	ChannelId    string                       `json:"channel_id"`    // 客户渠道名称
}

type BatchWithdrawSuborderItem struct {
	ID                   int64           `json:"id,omitempty"`
	BatchID              string          `json:"batch_id,omitempty"`
	MerchantID           int64           `json:"merchant_id"`
	ChannelID            string          `json:"channel_id"`
	SuborderID           string          `json:"suborder_id"`
	WithdrawID           string          `json:"withdraw_id,omitempty"`
	Chain                string          `json:"chain"`
	Address              string          `json:"address"`
	Currency             string          `json:"currency"`
	Amount               decimal.Decimal `json:"amount"`
	Fee                  decimal.Decimal `json:"fee"`
	TxID                 string          `json:"tx_id"`
	Timestamp            int64           `json:"timestamp,omitempty"`
	Memo                 string          `json:"memo"`
	Status               string          `json:"status"`
	MerchantWithdrawId   string          `json:"merchant_withdraw_id"`
	ErrMsg               string          `json:"err_msg,omitempty"`
	ClientID             string          `json:"client_id,omitempty"`
	CreateTime           int64           `json:"create_time"`
	UpdateTime           int64           `json:"update_time"`
	FeeType              int8            `json:"fee_type"`
	BatchWithdrawId      string          `json:"batch_withdraw_id"`
	Desc                 string          `json:"desc"`
	ReconciliationStatus int8            `json:"reconciliation_status"`
	IsPlaced             int             `json:"is_placed"`
	FinishTime           int64           `json:"finish_time,omitempty"`
	SubAmount            decimal.Decimal `json:"sub_amount,omitempty"`
	DoneAmount           decimal.Decimal `json:"done_amount,omitempty"`
}

type WalletCurrency struct {
	common.BaseRequest
	Currency string `json:"currency"`
}

type CurrencyChainItem struct {
	Chain              string `json:"chain"`
	NameCn             string `json:"name_cn"`
	NameEn             string `json:"name_en"`
	ContractAddress    string `json:"contract_address"`
	IsDisabled         int    `json:"is_disabled"`
	IsDepositDisabled  int    `json:"is_deposit_disabled"`
	IsWithdrawDisabled int    `json:"is_withdraw_disabled"`
}

type TotalBanlanceResp struct {
	Details struct {
		CrossMargin struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"cross_margin"`
		Spot struct {
			Currency string `json:"currency"`
			Amount   string `json:"amount"`
		} `json:"spot"`
		Finance struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"finance"`
		Margin struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
			Borrowed string `json:"borrowed"`
		} `json:"margin"`
		Quant struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"quant"`
		Futures struct {
			Amount        string `json:"amount"`
			Currency      string `json:"currency"`
			UnrealisedPnl string `json:"unrealised_pnl"`
		} `json:"futures"`
		Delivery struct {
			Currency      string `json:"currency"`
			Amount        string `json:"amount"`
			UnrealisedPnl string `json:"unrealised_pnl"`
		} `json:"delivery"`
		Warrant struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"warrant"`
		Cbbc struct {
			Currency string `json:"currency"`
			Amount   string `json:"amount"`
		} `json:"cbbc"`
	} `json:"details"`
	Total struct {
		Currency      string `json:"currency"`
		Amount        string `json:"amount"`
		UnrealisedPnl string `json:"unrealised_pnl"`
		Borrowed      string `json:"borrowed"`
	} `json:"total"`
}

type WithdrawStatusItem struct {
	Currency                string            `json:"currency"`
	Name                    string            `json:"name"`
	NameCn                  string            `json:"name_cn"`
	Deposit                 string            `json:"deposit"`
	WithdrawPercent         string            `json:"withdraw_percent"`
	WithdrawFix             string            `json:"withdraw_fix"`
	WithdrawDayLimit        string            `json:"withdraw_day_limit"`
	WithdrawDayLimitRemain  string            `json:"withdraw_day_limit_remain"`
	WithdrawAmountMini      string            `json:"withdraw_amount_mini"`
	WithdrawEachtimeLimit   string            `json:"withdraw_eachtime_limit"`
	WithdrawFixOnChains     map[string]string `json:"withdraw_fix_on_chains"`
	WithdrawPercentOnChains map[string]string `json:"withdraw_percent_on_chains"`
}

type WithdrawalsReq struct {
	common.BaseRequest
	Currency        string `json:"currency" query:"currency"`
	WithdrawId      string `json:"withdraw_id" query:"withdraw_id"`
	AssetClass      string `json:"asset_class" query:"asset_class"`
	WithdrawOrderId string `json:"withdraw_order_id" query:"withdraw_order_id"`
	From            int64  `json:"from" query:"from"`
	To              int64  `json:"to" query:"to"`
	Limit           int    `json:"limit" query:"limit"`
	Offset          int    `json:"offset" query:"offset"`
}
