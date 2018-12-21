package proto

const ClientTryGetInitializersURL = `client_try_get_initializers`

type ClientTryGetInitializersRequest struct {
	OutSubMchID string `json:"out_sub_mch_id,omitempty"`
	Token       string `json:"token,omitempty"`
}

type CipherDesc struct {
	Algorithm EncryptType `json:"algorithm,omitempty"`
	Key       string      `json:"key,omitempty"`
}

type ClientMchInfo struct {
	CompanyName  string `json:"company_name,omitempty"`
	OutMchID     string `json:"out_mch_id,omitempty"`
	MerchantName string `json:"merchant_name,omitempty"`
	Logo         []byte `json:"logo,omitempty"`
}

type ClientSubMchInfo struct {
	CompanyName                string        `json:"company_name,omitempty"`
	CloudCashierID             string        `json:"cloud_cashier_id,omitempty"`
	OutSubMchID                string        `json:"out_sub_mch_id,omitempty"`
	MerchantName               string        `json:"merchant_name,omitempty"`
	ClientConfigurationCiphers []*CipherDesc `json:"client_configuration_ciphers,omitempty"`
	Logo                       []byte        `json:"logo,omitempty"`
}

type ClientShopDeviceInfo struct {
	DeviceID   string `json:"device_id,omitempty"`
	DeviceType int    `json:"device_type,omitempty"`
}

type ClientShopStaffInfo struct {
	StaffID    string `json:"staff_id,omitempty"`
	StaffName  string `json:"staff_name,omitempty"`
	RefundAuth bool   `json:"refund_auth,omitempty"`
}

type ClientShopInfo struct {
	OutShopID string                  `json:"out_shop_id,omitempty"`
	ShopName  string                  `json:"shop_name,omitempty"`
	Devices   []*ClientShopDeviceInfo `json:"devices,omitempty"`
	Staffs    []*ClientShopStaffInfo  `json:"staffs,omitempty"`
}

type ManagerTypeEntityIDPair struct {
	ManagerType int    `json:"manager_type,omitempty"`
	OutID       string `json:"out_id,omitempty"`
	StaffID     string `json:"staff_id,omitempty"`
}

type ClientTryGetInitializersResponseDetail struct {
	InitializationGranted bool                       `json:"initialization_granted,omitempty"`
	Identity              string                     `json:"identity,omitempty"`
	StaffID               string                     `json:"staff_id,omitempty"`
	ManagerTypes          []int                      `json:"manager_types,omitempty"`
	ManagerEntities       []*ManagerTypeEntityIDPair `json:"manager_entities,omitempty"`
	MchInfo               *ClientMchInfo             `json:"mch_info,omitempty"`
	SubMchInfo            *ClientSubMchInfo          `json:"sub_mch_info,omitempty"`
	ShopInfos             []*ClientShopInfo          `json:"shop_infos,omitempty"`
}

type ClientTryGetInitializersResponse struct {
	ResponseStatus
	InternalStatus           uint                                    `json:"internal_status,omitempty"`
	LogID                    uint                                    `json:"log_id,omitempty"`
	ClientTryGetInitializers *ClientTryGetInitializersResponseDetail `json:"client_try_get_initializers,omitempty"`
}
