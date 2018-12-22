package proto

type PaySucceedCbRequest struct {
	NonceStr     string        `json:"nonce_str,omitempty"`
	OrderClient  *OrderClient  `json:"order_client,omitempty"`
	OrderContent *OrderContent `json:"order_content,omitempty"`
	PayMchKey    *PayMchKey    `json:"pay_mch_key,omitempty"`
}

type PaySucceedCbResponseDetail struct {
	Status      int    `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
}

type PaySucceedCbResponse struct {
	ResponseContent *PaySucceedCbResponseDetail `json:"response_content,omitempty"`
	AuthenInfo      *AuthenInfo                 `json:"authen_info,omitempty"`
}
