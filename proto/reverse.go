package proto

const ReverseURL = `reverse` //撤单的url

type ReverseRequest struct {
	PayMchKey   *PayMchKey   `json:"pay_mch_key,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	OutTradeNo  string       `json:"out_trade_no,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
}

type ReverseResponseDetail struct {
	NonceStr string `json:"nonce_str,omitempty"`
}

type ReverseResponse struct {
	ResponseStatus
	InternalStatus uint                   `json:"internal_status,omitempty"`
	LogID          uint                   `json:"log_id,omitempty"`
	Reverse        *ReverseResponseDetail `json:"reverse,omitempty"`
}
