package proto

const QueryOrderURL = `query_order` //查单的url

type QueryOrderRequest struct {
	PayMchKey   *PayMchKey   `json:"pay_mch_key,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	OutTradeNo  string       `json:"out_trade_no,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
	TradeType   TradeType    `json:"trade_type,omitempty"`
}

type QueryOrderResponseDetail struct {
	NonceStr     string        `json:"nonce_str,omitempty"`
	PayMchKey    *PayMchKey    `json:"pay_mch_key,omitempty"`
	OrderContent *OrderContent `json:"order_content,omitempty"`
}

type QueryOrderResponse struct {
	ResponseStatus
	InternalStatus uint                      `json:"internal_status,omitempty"`
	LogID          uint                      `json:"log_id,omitempty"`
	QueryOrder     *QueryOrderResponseDetail `json:"query_order,omitempty"`
}
