package proto

const OrderDetailURL = `client_order_detail` //客户端订单查询 url

type OrderDetailRequest struct {
	PayPlatform      int    `json:"pay_platform,omitempty"`
	OutSubMchID      string `json:"out_sub_mch_id,omitempty"`
	OutShopID        string `json:"out_shop_id,omitempty"`
	StaffID          string `json:"staff_id,omitempty"`
	QueryOrderType   int    `json:"query_order_type,omitempty"`
	TradeType        int    `json:"trade_type,omitempty"`
	StartTime        int    `json:"start_time,omitempty"`
	EndTime          int    `json:"end_time,omitempty"`
	PageNum          uint   `json:"page_num,omitempty"`
	PageSize         uint   `json:"page_size,omitempty"`
	EnableOverview   bool   `json:"enable_overview,omitempty"`
	DeviceID         string `json:"device_id,omitempty"`
	QueryCashierType int    `json:"query_cashier_type,omitempty"`
}

type WxpayOrderMchExt struct {
	MchID          string `json:"mch_id,omitempty"`
	SubMchID       string `json:"sub_mch_id,omitempty"`
	ShopID         string `json:"shop_id,omitempty"`
	AppID          string `json:"app_id,omitempty"`
	SubAppID       string `json:"sub_app_id,omitempty"`
	OpenID         string `json:"open_id,omitempty"`
	SubOpenID      string `json:"sub_open_id,omitempty"`
	IsSubscribe    bool   `json:"is_subscribe,omitempty"`
	SubIsSubscribe bool   `json:"sub_is_subscribe,omitempty"`
}

type AlipayOrderMchExt struct {
	AppID    string `json:"app_id,omitempty"`
	SubAppID string `json:"sub_app_id,omitempty"`
	StoreID  string `json:"store_id,omitempty"`
	UserID   string `json:"user_id,omitempty"`
}

type OrderMch struct {
	PayPlatform   PayPlatform `json:"pay_platform,omitempty"`
	OutMchID      string      `json:"out_mch_id,omitempty"`
	OutChannelID  string      `json:"out_channel_id,omitempty"`
	OutSubMchID   string      `json:"out_sub_mch_id,omitempty"`
	OutShopID     string      `json:"out_shop_id,omitempty"`
	MchID         string      `json:"mch_id,omitempty"`
	SubMchID      string      `json:"sub_mch_id,omitempty"`
	SubMchPayInfo string      `json:"sub_mch_pay_info,omitempty"`
	MchUin        string      `json:"mch_uin,omitempty"`
	MchSubUin     string      `json:"mch_sub_uin,omitempty"`
	// // 第三方支付平台特有信息
	// oneof order_mch_ext {
	WxpayOrderMchExt  *WxpayOrderMchExt  `json:"wxpay_order_mch_ext,omitempty"`
	AlipayOrderMchExt *AlipayOrderMchExt `json:"alipay_order_mch_ext,omitempty"`
}

type Order struct {
	OrderMch     *OrderMch             `json:"order_mch,omitempty"`
	OrderContent *OrderContent         `json:"order_content,omitempty"`
	OrderClient  *OrderClient          `json:"order_client,omitempty"`
	AuthenInfo   *AuthenInfo           `json:"authen_info,omitempty"`
	Originator   TransactionOriginator `json:"originator,omitempty"`
}

type RefundOrder struct {
	RefundOrderMch     *OrderMch             `json:"refund_order_mch,omitempty"`
	RefundOrderContent *RefundOrderContent   `json:"refund_order_content,omitempty"`
	OrderClient        *OrderClient          `json:"order_client,omitempty"`
	AuthenInfo         *AuthenInfo           `json:"authen_info,omitempty"`
	Originator         TransactionOriginator `json:"originator,omitempty"`
}

type Receipt struct {
	// Types that are valid to be assigned to R:
	//	*Receipt_Order
	//	*Receipt_RefundOrder
	R isReceipt_R
}

type isReceipt_R interface {
	isReceipt_R()
}

type Receipt_Order struct {
	Order *Order
}

type Receipt_RefundOrder struct {
	RefundOrder *RefundOrder
}

func (*Receipt_Order) isReceipt_R() {}

func (*Receipt_RefundOrder) isReceipt_R() {}

//

type OrderDetailResponseDetail struct {
	TotalCount   uint                                     `json:"total_count,omitempty"`
	OrderDetails []*OrderDetailResponseDetail_OrderDetail `json:"order_details,omitempty"`
	Overview     *OrderStatClientInfo                     `json:"overview,omitempty"`
}

type OrderDetailResponseDetail_OrderDetail struct {
	ShopInfo         *ShopInfo  `json:"shop_info,omitempty"`
	ShopStaffInfo    *StaffInfo `json:"shop_staff_info,omitempty"`
	Receipt          *Receipt   `json:"receipt,omitempty"`
	DisplayTimeField string     `json:"display_time_field,omitempty"`
}

type OrderDetailResponse struct {
	ResponseStatus
	InternalStatus   uint                       `json:"internal_status,omitempty"`
	LogID            uint                       `json:"log_id,omitempty"`
	OrderDetailQuery *OrderDetailResponseDetail `json:"order_detail_query,omitempty"`
}
