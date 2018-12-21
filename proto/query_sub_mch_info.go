package proto

const QuerySubMchInfoURL = `sdk_query_sub_mch_info`

type QuerySubMchInfoRequest struct {
	OutSubMchID string `json:"out_sub_mch_id,omitempty"`
	NonceStr    string `json:"nonce_str,omitempty"`
	PageNum     int    `json:"page_num,omitempty"`
	PageSize    int    `json:"page_size,omitempty"`
}

//  This structure is generated from data.
type QuerySubMchInfoResponse struct {
	Description     string                         `json:"description,omitempty"`
	InternalStatus  int                            `json:"internal_status,omitempty"`
	LogID           int                            `json:"log_id,omitempty"`
	QuerySubMchInfo *QuerySubMchInfoResponseDetail `json:"query_sub_mch_info,omitempty"`
	Status          int                            `json:"status,omitempty"`
}

// QuerySubMchInfoResponse This structure is generated from data.
type QuerySubMchInfoResponseDetail struct {
	NonceStr    string                        `json:"nonce_str,omitempty"`
	SubMchInfos []*QuerySubMchInfoSubMchInfos `json:"sub_mch_infos,omitempty"`
	TotalCount  int                           `json:"total_count,omitempty"`
}

// QuerySubMchInfoSubMchInfos This structure is generated from data.
type QuerySubMchInfoSubMchInfos struct {
	CompanyName       string                                   `json:"company_name,omitempty"`
	MchID             string                                   `json:"mch_id,omitempty"`
	MchKey            *QuerySubMchInfoSubMchInfosMchKey        `json:"mch_key,omitempty"`
	MchSubCompanyName string                                   `json:"mch_sub_company_name,omitempty"`
	MchSubUin         string                                   `json:"mch_sub_uin,omitempty"`
	OutMchID          string                                   `json:"out_mch_id,omitempty"`
	PayPlatform       int                                      `json:"pay_platform,omitempty"`
	SubMchInfos       []*QuerySubMchInfoSubMchInfosSubMchInfos `json:"sub_mch_infos,omitempty"`
}

// QuerySubMchInfoSubMchInfosMchKey This structure is generated from data.
type QuerySubMchInfoSubMchInfosMchKey struct {
	MchID          string                                          `json:"mch_id,omitempty"`
	PayPlatform    int                                             `json:"pay_platform,omitempty"`
	WxpayMchKeyExt *QuerySubMchInfoSubMchInfosMchKeyWxpayMchKeyExt `json:"wxpay_mch_key_ext,omitempty"`
}

// QuerySubMchInfoSubMchInfosMchKeyWxpayMchKeyExt This structure is generated from data.
type QuerySubMchInfoSubMchInfosMchKeyWxpayMchKeyExt struct {
	AppID string `json:"app_id,omitempty"`
}

// QuerySubMchInfoSubMchInfosSubMchInfos This structure is generated from data.
type QuerySubMchInfoSubMchInfosSubMchInfos struct {
	AdPageURL                 string                                                    `json:"ad_page_url,omitempty"`
	AdminEmail                string                                                    `json:"admin_email,omitempty"`
	AlipaySubMchInfoExt       *QuerySubMchInfoSubMchInfosSubMchInfosAlipaySubMchInfoExt `json:"alipay_sub_mch_info_ext,omitempty"`
	AuthenKey                 string                                                    `json:"authen_key,omitempty"`
	AuthenType                int                                                       `json:"authen_type,omitempty"`
	BankRate                  int                                                       `json:"bank_rate,omitempty"`
	CloudCashierID            string                                                    `json:"cloud_cashier_id,omitempty"`
	CompanyName               string                                                    `json:"company_name,omitempty"`
	CreateTime                string                                                    `json:"create_time,omitempty"`
	DefaultOrderBody          string                                                    `json:"default_order_body,omitempty"`
	Desc                      string                                                    `json:"desc,omitempty"`
	Direct                    bool                                                      `json:"direct,omitempty"`
	DisableCardOnlineRecharge bool                                                      `json:"disable_card_online_recharge,omitempty"`
	EncrytKey                 string                                                    `json:"encryt_key,omitempty"`
	EncrytType                int                                                       `json:"encryt_type,omitempty"`
	IsUseCpayShopSystem       bool                                                      `json:"is_use_cpay_shop_system,omitempty"`
	LastUpdateTime            string                                                    `json:"last_update_time,omitempty"`
	Logo                      string                                                    `json:"logo,omitempty"`
	MerchantName              string                                                    `json:"merchant_name,omitempty"`
	OneCodePayAdInfo          *QuerySubMchInfoSubMchInfosSubMchInfosOneCodePayAdInfo    `json:"one_code_pay_ad_info,omitempty"`
	OutSubMchID               string                                                    `json:"out_sub_mch_id,omitempty"`
	OutSubMchIDURL            string                                                    `json:"out_sub_mch_id_url,omitempty"`
	Phone                     string                                                    `json:"phone,omitempty"`
	PublicKey                 string                                                    `json:"public_key,omitempty"`
	SignType                  int                                                       `json:"sign_type,omitempty"`
	SubMchAdminInfos          []*QuerySubMchInfoSubMchInfosSubMchInfosSubMchAdminInfos  `json:"sub_mch_admin_infos,omitempty"`
	SubMchID                  string                                                    `json:"sub_mch_id,omitempty"`
	SubMchSource              int                                                       `json:"sub_mch_source,omitempty"`
	TypedCashierIds           []*QuerySubMchInfoSubMchInfosSubMchInfosTypedCashierIds   `json:"typed_cashier_ids,omitempty"`
	UpstreamCompanyName       string                                                    `json:"upstream_company_name,omitempty"`
	UpstreamOutChannelID      string                                                    `json:"upstream_out_channel_id,omitempty"`
	WxpaySubMchInfoExt        *QuerySubMchInfoSubMchInfosSubMchInfosWxpaySubMchInfoExt  `json:"wxpay_sub_mch_info_ext,omitempty"`
}

// QuerySubMchInfoSubMchInfosSubMchInfosAlipaySubMchInfoExt This structure is generated from data.
type QuerySubMchInfoSubMchInfosSubMchInfosAlipaySubMchInfoExt struct {
	AliAuthorizationURL string `json:"ali_authorization_url,omitempty"`
	Direct              bool   `json:"direct,omitempty"`
	SubAppID            string `json:"sub_app_id,omitempty"`
	SubMchUserID        string `json:"sub_mch_user_id,omitempty"`
}

// QuerySubMchInfoSubMchInfosSubMchInfosOneCodePayAdInfo This structure is generated from data.
type QuerySubMchInfoSubMchInfosSubMchInfosOneCodePayAdInfo struct {
	Picture string `json:"picture,omitempty"`
	URL     string `json:"url,omitempty"`
}

// QuerySubMchInfoSubMchInfosSubMchInfosSubMchAdminInfos This structure is generated from data.
type QuerySubMchInfoSubMchInfosSubMchInfosSubMchAdminInfos struct {
	CloudPayOpenID          string `json:"cloud_pay_open_id,omitempty"`
	Duty                    string `json:"duty,omitempty"`
	MchOpenID               string `json:"mch_open_id,omitempty"`
	Name                    string `json:"name,omitempty"`
	ReceiveOneCodePayNotify bool   `json:"receive_one_code_pay_notify,omitempty"`
}

// QuerySubMchInfoSubMchInfosSubMchInfosTypedCashierIds This structure is generated from data.
type QuerySubMchInfoSubMchInfosSubMchInfosTypedCashierIds struct {
	ID   string `json:"id,omitempty"`
	Type int    `json:"type,omitempty"`
}

// QuerySubMchInfoSubMchInfosSubMchInfosWxpaySubMchInfoExt This structure is generated from data.
type QuerySubMchInfoSubMchInfosSubMchInfosWxpaySubMchInfoExt struct {
	Direct   bool   `json:"direct,omitempty"`
	IsMicro  bool   `json:"is_micro,omitempty"`
	SubAppID string `json:"sub_app_id,omitempty"`
}
