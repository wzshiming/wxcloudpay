package proto

const MicroPayURL = `micro_pay` //刷卡支付的url

type MicroPayRequest struct {
	PayMchKey   *PayMchKey   `json:"pay_mch_key,omitempty"`
	PayContent  *PayContent  `json:"pay_content,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
}

type MicroPayResponseDetail struct {
	NonceStr     string        `json:"nonce_str,omitempty"`
	PayMchKey    *PayMchKey    `json:"pay_mch_key,omitempty"`
	OrderContent *OrderContent `json:"order_content,omitempty"`
}

type MicroPayResponse struct {
	ResponseStatus
	InternalStatus uint                    `json:"internal_status,omitempty"`
	LogID          uint                    `json:"log_id,omitempty"`
	MicroPay       *MicroPayResponseDetail `json:"micro_pay,omitempty"`
}
