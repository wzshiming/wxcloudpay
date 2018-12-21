package proto

const CloseOrderURL = `close_order` //关单url

type CloseOrderRequest struct {
	PayMchKey   *PayMchKey   `json:"pay_mch_key,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	OutTradeNo  string       `json:"out_trade_no,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
	TradeType   TradeType    `json:"trade_type,omitempty"`
}

type CloseOrderResponseDetail struct {
	NonceStr string `json:"nonce_str,omitempty"`
}

type CloseOrderResponse struct {
	ResponseStatus
	InternalStatus uint                      `json:"internal_status,omitempty"`
	LogID          uint                      `json:"log_id,omitempty"`
	CloseOrder     *CloseOrderResponseDetail `json:"close_order,omitempty"`
}
