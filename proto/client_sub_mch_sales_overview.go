package proto

const ClientSubMchSalesOverviewURL = `client_sub_mch_sales_overview`

type ClientSubMchSalesOverviewRequest struct {
	OutMchID    string                                         `json:"out_mch_id,omitempty"`
	OutSubMchID string                                         `json:"out_sub_mch_id,omitempty"`
	Shops       []*ClientSubMchSalesOverviewRequest_ShopEntity `json:"shops,omitempty"`
	PayPlatform int                                            `json:"pay_platform,omitempty"`
	TradeType   int                                            `json:"trade_type,omitempty"`
	StatType    int                                            `json:"stat_type,omitempty"`
	StartTime   int                                            `json:"start_time,omitempty"`
	EndTime     int                                            `json:"end_time,omitempty"`
	PageNum     uint                                           `json:"page_num,omitempty"`
	PageSize    uint                                           `json:"page_size,omitempty"`
	StatRole    int                                            `json:"stat_role,omitempty"`
}

type ClientSubMchSalesOverviewRequest_ShopEntity struct {
	OutShopID string   `json:"out_shop_id,omitempty"`
	StaffIDs  []string `json:"staff_ids,omitempty"`
}

type OrderStatClientInfo struct {
	Date               int    `json:"date,omitempty"`
	MchUin             string `json:"mch_uin,omitempty"`
	OutMchID           string `json:"out_mch_id,omitempty"`
	MchID              string `json:"mch_id,omitempty"`
	MchSubUin          string `json:"mch_sub_uin,omitempty"`
	OutSubMchID        string `json:"out_sub_mch_id,omitempty"`
	SubMchID           string `json:"sub_mch_id,omitempty"`
	PayPlatform        int    `json:"pay_platform,omitempty"`
	OutShopID          string `json:"out_shop_id,omitempty"`
	ShopID             string `json:"shop_id,omitempty"`
	StaffID            string `json:"staff_id,omitempty"`
	TotalCount         int    `json:"total_count,omitempty"`
	TotalAmount        int    `json:"total_amount,omitempty"`
	SuccessCount       int    `json:"success_count,omitempty"`
	SuccessAmount      int    `json:"success_amount,omitempty"`
	SettleAmount       int    `json:"settle_amount,omitempty"`
	RefundCount        int    `json:"refund_count,omitempty"`
	RefundAmount       int    `json:"refund_amount,omitempty"`
	NetCount           int    `json:"net_count,omitempty"`
	NetAmount          int    `json:"net_amount,omitempty"`
	DistinctPayerCount int    `json:"distinct_payer_count,omitempty"`
	PaySettleAmount    int    `json:"pay_settle_amount,omitempty"`
	RevokeSettleAmount int    `json:"revoke_settle_amount,omitempty"`
	RefundSettleAmount int    `json:"refund_settle_amount,omitempty"`
}

type ClientSubMchSalesOverviewResponseDetail struct {
	TotalCount int                    `json:"total_count,omitempty"`
	Overview   *OrderStatClientInfo   `json:"overview,omitempty"`
	StatInfos  []*OrderStatClientInfo `json:"stat_infos,omitempty"`
}

type ClientSubMchSalesOverviewResponse struct {
	ResponseStatus
	InternalStatus            uint                                     `json:"internal_status,omitempty"`
	LogID                     uint                                     `json:"log_id,omitempty"`
	ClientSubMchSalesOverview *ClientSubMchSalesOverviewResponseDetail `json:"client_sub_mch_sales_overview,omitempty"`
}
