package proto

const QueryIdentityMappingURL = `query_identity_mapping`

type QueryIdentityMappingRequest struct {
	// Types that are valid to be assigned to Condition:
	//	*QueryIdentityMappingRequest_Mch_
	//	*QueryIdentityMappingRequest_Out_
	Condition isQueryIdentityMappingRequest_Condition
}

type isQueryIdentityMappingRequest_Condition interface {
	isQueryIdentityMappingRequest_Condition()
}

type QueryIdentityMappingRequest_Mch_ struct {
	Mch *QueryIdentityMappingRequest_Mch
}

type QueryIdentityMappingRequest_Out_ struct {
	Out *QueryIdentityMappingRequest_Out
}

func (*QueryIdentityMappingRequest_Mch_) isQueryIdentityMappingRequest_Condition() {}

func (*QueryIdentityMappingRequest_Out_) isQueryIdentityMappingRequest_Condition() {}

//

type QueryIdentityMappingRequest_Mch struct {
	MchID    string `json:"mch_id,omitempty"`
	SubMchID string `json:"sub_mch_id,omitempty"`
	ShopID   string `json:"shop_id,omitempty"`
}

type QueryIdentityMappingRequest_Out struct {
	OutMchID    string `json:"out_mch_id,omitempty"`
	OutSubMchID string `json:"out_sub_mch_id,omitempty"`
	OutShopID   string `json:"out_shop_id,omitempty"`
}

type QueryIdentityMappingResponseDetail struct {
	OutMchID       string   `json:"out_mch_id,omitempty"`
	OutSubMchID    string   `json:"out_sub_mch_id,omitempty"`
	OutShopID      string   `json:"out_shop_id,omitempty"`
	NonceStr       string   `json:"nonce_str,omitempty"`
	WxAuthorCodes  []string `json:"wx_author_codes,omitempty"`
	AliAuthorCodes []string `json:"ali_author_codes,omitempty"`
}

type QueryIdentityMappingResponse struct {
	ResponseStatus
	InternalStatus       uint                                `json:"internal_status,omitempty"`
	LogID                uint                                `json:"log_id,omitempty"`
	QueryIdentityMapping *QueryIdentityMappingResponseDetail `json:"query_identity_mapping,omitempty"`
}
