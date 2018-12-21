package proto

const QueryShopInfoURL = `query_sub_mch_shop_info`

type QueryShopInfoRequest struct {
	OutMchID    string `json:"out_mch_id,omitempty"`
	OutSubMchID string `json:"out_sub_mch_id,omitempty"`
	NonceStr    string `json:"nonce_str,omitempty"`
	PageNum     int    `json:"page_num,omitempty"`
	PageSize    int    `json:"page_size,omitempty"`
}

type QueryShopInfoResponse struct {
	ResponseStatus
	InternalStatus int                          `json:"internal_status,omitempty"`
	LogID          int                          `json:"log_id,omitempty"`
	QueryShopInfo  *QueryShopInfoResponseDetail `json:"query_shop_info,omitempty"`
}

// QueryShopInfo This structure is generated from data.
type QueryShopInfoResponseDetail struct {
	NonceStr   string                    `json:"nonce_str,omitempty"`
	ShopInfos  []*QueryShopInfoShopInfos `json:"shop_infos,omitempty"`
	TotalCount int                       `json:"total_count,omitempty"`
}

// QueryShopInfoShopInfos This structure is generated from data.
type QueryShopInfoShopInfos struct {
	Address                string                                   `json:"address,omitempty"`
	AlipayShopInfoExt      *QueryShopInfoShopInfosAlipayShopInfoExt `json:"alipay_shop_info_ext,omitempty"`
	BranchName             string                                   `json:"branch_name,omitempty"`
	City                   string                                   `json:"city,omitempty"`
	CoordinateType         int                                      `json:"coordinate_type,omitempty"`
	CreateTime             int                                      `json:"create_time,omitempty"`
	District               string                                   `json:"district,omitempty"`
	FeeType                string                                   `json:"fee_type,omitempty"`
	GoodsTag               string                                   `json:"goods_tag,omitempty"`
	Height                 string                                   `json:"height,omitempty"`
	Latitude               string                                   `json:"latitude,omitempty"`
	Longitude              string                                   `json:"longitude,omitempty"`
	OutShopID              string                                   `json:"out_shop_id,omitempty"`
	OutShopIDURL           string                                   `json:"out_shop_id_url,omitempty"`
	PayPageMode            int                                      `json:"pay_page_mode,omitempty"`
	Phone                  string                                   `json:"phone,omitempty"`
	Province               string                                   `json:"province,omitempty"`
	ReceiveNotifyAfterDuty bool                                     `json:"receive_notify_after_duty,omitempty"`
	ShopID                 string                                   `json:"shop_id,omitempty"`
	ShopName               string                                   `json:"shop_name,omitempty"`
	ShopType               int                                      `json:"shop_type,omitempty"`
	UpdateTime             int                                      `json:"update_time,omitempty"`
	WxpayShopInfoExt       *QueryShopInfoShopInfosWxpayShopInfoExt  `json:"wxpay_shop_info_ext,omitempty"`
}

// QueryShopInfoShopInfosAlipayShopInfoExt This structure is generated from data.
type QueryShopInfoShopInfosAlipayShopInfoExt struct {
	AliAuthorizationURL string `json:"ali_authorization_url,omitempty"`
}

// QueryShopInfoShopInfosWxpayShopInfoExt This structure is generated from data.
type QueryShopInfoShopInfosWxpayShopInfoExt struct {
	ExtSubMchID string `json:"ext_sub_mch_id,omitempty"`
}
