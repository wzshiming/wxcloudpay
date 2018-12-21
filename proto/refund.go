package proto

const RefundURL = `refund` //申请退款url

type RefundRequest struct {
	PayMchKey     *PayMchKey     `json:"pay_mch_key,omitempty"`
	RefundContent *RefundContent `json:"refund_content,omitempty"`
	OrderClient   *OrderClient   `json:"order_client,omitempty"`
	NonceStr      string         `json:"nonce_str,omitempty"`
}

type RefundOrderContent struct {
	OutRefundNo            string    `json:"out_refund_no,omitempty"`
	RefundID               string    `json:"refund_id,omitempty"`
	OutTradeNo             string    `json:"out_trade_no,omitempty"`
	TradeType              TradeType `json:"trade_type,omitempty"`
	NonceStr               string    `json:"nonce_str,omitempty"`
	CreateTime             int       `json:"create_time,omitempty"`
	LastUpdateTime         int       `json:"last_update_time,omitempty"`
	IsTransforming         bool      `json:"is_transforming,omitempty"`
	TotalFee               int       `json:"total_fee,omitempty"`
	RefundFee              int       `json:"refund_fee,omitempty"`
	RefundFeeType          string    `json:"refund_fee_type,omitempty"`
	RefundSuccessTime      int       `json:"refund_success_time,omitempty"`
	RefundReason           string    `json:"refund_reason,omitempty"`
	InnerRefundSuccessTime int       `json:"inner_refund_success_time,omitempty"`
	// 第三方支付平台特有信息
	// oneof refund_order_content_ext
	WxpayRefundOrderContentExt  *WxpayRefundOrderContentExt  `json:"wxpay_refund_order_content_ext,omitempty"`
	AlipayRefundOrderContentExt *AlipayRefundOrderContentExt `json:"alipay_refund_order_content_ext,omitempty"`
}

type AlipayRefundOrderContentExt struct {
	FundChange           string                 `json:"fund_change,omitempty"`
	GmtRefundPay         int                    `json:"gmt_refund_pay,omitempty"`
	RefundDetailItemList []*AlipayFundBill      `json:"refund_detail_item_list,omitempty"`
	RefundStatus         AlipayRefundOrderState `json:"refund_status,omitempty"`
}

type RefundResponseDetail struct {
	NonceStr           string              `json:"nonce_str,omitempty"`
	PayMchKey          *PayMchKey          `json:"pay_mch_key,omitempty"`
	RefundOrderContent *RefundOrderContent `json:"refund_order_content,omitempty"`
}

type RefundResponse struct {
	ResponseStatus
	InternalStatus uint                  `json:"internal_status,omitempty"`
	LogID          uint                  `json:"log_id,omitempty"`
	Refund         *RefundResponseDetail `json:"refund,omitempty"`
}
