package proto

type RefundSucceedCbRequest struct {
	NonceStr           string                `json:"nonce_str,omitempty"`
	RefundOrderContent []*RefundOrderContent `json:"refund_order_content,omitempty"`
	PayMchKey          *PayMchKey            `json:"pay_mch_key,omitempty"`
}

type RefundSucceedCbResponseDetail struct {
	Status      int    `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
}

type RefundSucceedCbResponse struct {
	ResponseContent *RefundSucceedCbResponseDetail `json:"response_content,omitempty"`
	AuthenInfo      *AuthenInfo                    `json:"authen_info,omitempty"`
}
