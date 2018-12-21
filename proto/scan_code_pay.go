package proto

const ScanCodePayURL = `scan_code_pay` //扫码支付url

type ScanCodePayRequest struct {
	PayMchKey   *PayMchKey   `json:"pay_mch_key,omitempty"`
	PayContent  *PayContent  `json:"pay_content,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
}

type ScanCodePayResponseDetail struct {
	NonceStr string `json:"nonce_str,omitempty"`
	CodeUrl  string `json:"code_url,omitempty"`
}

type ScanCodePayResponse struct {
	ResponseStatus
	InternalStatus uint                       `json:"internal_status,omitempty"`
	LogID          uint                       `json:"log_id,omitempty"`
	ScanCodePay    *ScanCodePayResponseDetail `json:"scan_code_pay,omitempty"`
}
