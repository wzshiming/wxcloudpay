package proto

const SetShopInfoURL = `set_sub_mch_shop_info` //设置门店的ur

type SetShopInfoRequest struct {
	OutMchID    string      `json:"out_mch_id,omitempty"`
	OutSubMchID string      `json:"out_sub_mch_id,omitempty"`
	NonceStr    string      `json:"nonce_str,omitempty"`
	ShopInfos   []*ShopInfo `json:"shop_infos,omitempty"`
	IsAll       bool        `json:"is_all,omitempty"`
}

type SetShopInfoResponseDetail struct {
	NonceStr  string      `json:"nonce_str,omitempty"`
	ShopInfos []*ShopInfo `json:"shop_infos,omitempty"`
}

type SetShopInfoResponse struct {
	ResponseStatus
	InternalStatus uint                       `json:"internal_status,omitempty"`
	LogID          uint                       `json:"log_id,omitempty"`
	SetShopInfo    *SetShopInfoResponseDetail `json:"set_shop_info,omitempty"`
}
