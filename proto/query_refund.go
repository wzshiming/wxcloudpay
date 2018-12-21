package proto

const QueryRefundURL = `query_refund_order` //退款查询url

type QueryRefundRequest struct {
	PayMchKey   *PayMchKey   `json:"pay_mch_key,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	OutTradeNo  string       `json:"out_trade_no,omitempty"`
	OutRefundNo string       `json:"out_refund_no,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
}

type QueryRefundResponseDetail struct {
	NonceStr           string                `json:"nonce_str,omitempty"`
	PayMchKey          *PayMchKey            `json:"pay_mch_key,omitempty"`
	RefundOrderContent []*RefundOrderContent `json:"refund_order_content,omitempty"`
}

type QueryRefundResponse struct {
	ResponseStatus
	InternalStatus   uint                       `json:"internal_status,omitempty"`
	LogID            uint                       `json:"log_id,omitempty"`
	QueryRefundOrder *QueryRefundResponseDetail `json:"query_refund_order,omitempty"`
}
