package proto

const UploadClientConfInfoURL = `upload_client_conf_info` //机器配置上报url

type UploadClientConfInfoRequest struct {
	OutMchID    string       `json:"out_mch_id,omitempty"`
	OutSubMchID string       `json:"out_sub_mch_id,omitempty"`
	OutShopID   string       `json:"out_shop_id,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	MachineInfo string       `json:"machine_info,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
}

type UploadClientConfInfoResponseDetail struct {
	NonceStr string `json:"nonce_str,omitempty"`
}

type UploadClientConfInfoResponse struct {
	ResponseStatus
	InternalStatus       uint                                `json:"internal_status,omitempty"`
	LogID                uint                                `json:"log_id,omitempty"`
	UploadClientConfInfo *UploadClientConfInfoResponseDetail `json:"upload_client_conf_info,omitempty"`
}
