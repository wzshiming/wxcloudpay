package proto

type PayPlatform int

const (
	PayPlatformWxpay  PayPlatform = 1
	PayPlatformAlipay PayPlatform = 2
)

type AuthenType int

const (
	AuthenTypeHmacSha256 AuthenType = 1
)

type SignType int

const (
	SignTypeRsassaPss2048Sha256 SignType = 1
)

type EncryptType int

const (
	EncryptTypeAes128GCM EncryptType = 1
)

type CompressType int

const (
	CompressTypeNone CompressType = 0
	CompressTypeZip  CompressType = 1
	CompressTypeRAR  CompressType = 2
	CompressTypeGZip CompressType = 3
)

type AuthorizeTerminalType int

const (
	AuthorizeTerminalTypePC      AuthorizeTerminalType = 1
	AuthorizeTerminalTypeAndroid AuthorizeTerminalType = 2
)

type Interface int

const (
	InterfaceNone                      Interface = 0
	InterfaceSetShop                   Interface = 13
	InterfaceQueryShop                 Interface = 14
	InterfaceMicroPay                  Interface = 100
	InterfaceScanCodePay               Interface = 101
	InterfaceOfficialAccountPay        Interface = 102
	InterfaceCloseOrder                Interface = 103
	InterfaceRefund                    Interface = 104
	InterfaceQueryOrder                Interface = 106
	InterfaceQueryRefundOrder          Interface = 107
	InterfaceReverse                   Interface = 108
	InterfaceDownloadBill              Interface = 201
	InterfaceUploadClientMonitor       Interface = 202
	InterfaceUploadClientConf          Interface = 203
	InterfaceClientAuthorize           Interface = 210
	InterfaceClientTryGetInitializers  Interface = 212
	InterfaceQueryIDentityMapping      Interface = 213
	InterfacePing                      Interface = 50000
	InterfaceOrderDetailQuery          Interface = 221
	InterfaceClientSubMchSalesOverview Interface = 224
	InterfaceApp                       Interface = 100000
)

type WxpayOrderState int

const (
	WxpayOrderStateInit                WxpayOrderState = 1
	WxpayOrderStateMicropaySuccess     WxpayOrderState = 2
	WxpayOrderStateUnifiedorderSuccess WxpayOrderState = 3
	WxpayOrderStateRefund              WxpayOrderState = 4
	WxpayOrderStateMicropayNotpay      WxpayOrderState = 5
	WxpayOrderStateUnifiedorderNotpay  WxpayOrderState = 6
	WxpayOrderStateClosed              WxpayOrderState = 7
	WxpayOrderStateRevoked             WxpayOrderState = 8
	WxpayOrderStateUserpaying          WxpayOrderState = 9
	WxpayOrderStatePayerror            WxpayOrderState = 10
	WxpayOrderStateVoid                WxpayOrderState = 11
)

type AlipayOrderState int

const (
	AlipayOrderStateInit         AlipayOrderState = 1
	AlipayOrderStateSuccess      AlipayOrderState = 2
	AlipayOrderStateWaitBuyerPay AlipayOrderState = 4
	AlipayOrderStateClosed       AlipayOrderState = 5
	AlipayOrderStateFinish       AlipayOrderState = 6
	AlipayOrderStateVoid         AlipayOrderState = 7
)

type AlipayRefundOrderState int

const (
	AlipayRefundOrderStateInit    AlipayRefundOrderState = 1
	AlipayRefundOrderStateSuccess AlipayRefundOrderState = 2
	AlipayRefundOrderStateFail    AlipayRefundOrderState = 3
)

type TradeType int

const (
	TradeTypeMicroPay           TradeType = 1
	TradeTypeScanCodePay        TradeType = 2
	TradeTypeOfficialAccountPay TradeType = 3
	TradeTypeAppPay             TradeType = 4
	TradeTypeWavePay            TradeType = 5
	TradeTypeMobileWebPay       TradeType = 6
	TradeTypeSelfCreated        TradeType = 7
	TradeTypeOneCodePay         TradeType = 8
)

type WxpayRefundOrderState int

const (
	WxpayRefundOrderStateInit       WxpayRefundOrderState = 1
	WxpayRefundOrderStateSuccess    WxpayRefundOrderState = 2
	WxpayRefundOrderStateFail       WxpayRefundOrderState = 3
	WxpayRefundOrderStateProcessing WxpayRefundOrderState = 4
	WxpayRefundOrderStateChange     WxpayRefundOrderState = 5
	WxpayRefundOrderStateVoid       WxpayRefundOrderState = 6
)

type ManagerType int

const (
	ManagerTypeMchManager    ManagerType = 4
	ManagerTypeSubMchManager ManagerType = 1
	ManagerTypeShopManager   ManagerType = 2
	ManagerTypeShopStaff     ManagerType = 3
	ManagerTypeNotManager    ManagerType = 100
)

type QueryOrderType int

const (
	QueryOrderTypeOrder       QueryOrderType = 1
	QueryOrderTypeRefundOrder QueryOrderType = 2
	QueryOrderTypeAll         QueryOrderType = 3
)

type TransactionOriginator int

const (
	TransactionOriginatorChannel TransactionOriginator = 1
	TransactionOriginatorSubMch  TransactionOriginator = 2
)

type StatType int

const (
	StatTypeDaily   StatType = 1
	StatTypeMonthly StatType = 2
	StatTypeWeekly  StatType = 3
	StatTypeYearly  StatType = 4
)

type WxpayPayMchKeyExt struct {
	AppID     string `json:"app_id,omitempty"`
	SubAppID  string `json:"sub_app_id,omitempty"`
	OpenID    string `json:"open_id,omitempty"`
	SubOpenID string `json:"sub_open_id,omitempty"`
}

type PayMchKey struct {
	PayPlatform   PayPlatform `json:"pay_platform,omitempty"`
	OutMchID      string      `json:"out_mch_id,omitempty"`
	OutSubMchID   string      `json:"out_sub_mch_id,omitempty"`
	SubMchPayInfo string      `json:"sub_mch_pay_info,omitempty"`
	OutShopID     string      `json:"out_shop_id,omitempty"`
}

type GoodsDetail struct {
	GoodsID   string `json:"goods_id,omitempty"`
	GoodsName string `json:"goods_name,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
	Price     int64  `json:"price,omitempty"`
}

type Detail struct {
	CostPrice   int64          `json:"cost_price,omitempty"`
	ReceiptID   string         `json:"receipt_id,omitempty"`
	GoodsDetail []*GoodsDetail `json:"goods_detail,omitempty"`
}

// detail的格式为上面的Detail
type WxpayPayContentExt struct {
	Attach    string `json:"attach,omitempty"`
	GoodsTag  string `json:"goods_tag,omitempty"`
	ProductID string `json:"product_id,omitempty"`
	LimitPay  string `json:"limit_pay,omitempty"`
}

// 支付宝订单内容扩展信息
type AlipayPayContentExt struct {
	DiscountableAmount   int64  `json:"discountable_amount,omitempty"`
	UndiscountableAmount int64  `json:"undiscountable_amount,omitempty"`
	ProductCode          string `json:"product_code,omitempty"`
	RoyaltyInfo          string `json:"royalty_info,omitempty"`
	ExtendParams         string `json:"extend_params,omitempty"`
	DisablePayChannels   string `json:"disable_pay_channels,omitempty"`
	TimeoutExpress       string `json:"timeout_express,omitempty"`
}

type PayContent struct {
	OutTradeNo string  `json:"out_trade_no,omitempty"`
	AuthorCode string  `json:"author_code,omitempty"`
	TimeExpire *uint64 `json:"time_expire,omitempty"`
	TotalFee   int64   `json:"total_fee,omitempty"`
	FeeType    string  `json:"fee_type,omitempty"`
	Body       string  `json:"body,omitempty"`
	Detail     string  `json:"detail,omitempty"`
	// 第三方支付平台特有信息，二选一
	// oneof order_content_ext {
	WxpayPayContentExt  *WxpayPayContentExt  `json:"wxpay_pay_content_ext,omitempty"`
	AlipayPayContentExt *AlipayPayContentExt `json:"alipay_pay_content_ext,omitempty"`
}

type OrderClient struct {
	ShopID          string `json:"shop_id,omitempty"`
	DeviceID        string `json:"device_id,omitempty"`
	StaffID         string `json:"staff_id,omitempty"`
	TerminalType    uint   `json:"terminal_type,omitempty"`
	SubTerminalType uint   `json:"sub_terminal_type,omitempty"`
	MachineNo       string `json:"machine_no,omitempty"`
	SdkVersion      string `json:"sdk_version,omitempty"`
	SpbillCreateIp  string `json:"spbill_create_ip,omitempty"`
}

type WxpayCouponInfo struct {
	CouponBatchID string `json:"coupon_batch_id,omitempty"`
	CouponID      string `json:"coupon_id,omitempty"`
	CouponFee     int64  `json:"coupon_fee,omitempty"`
	CouponType    string `json:"coupon_type,omitempty"`
}

type WxpayOrderContentExt struct {
	CurrentTradeState WxpayOrderState    `json:"current_trade_state,omitempty"`
	ExpectTradeState  WxpayOrderState    `json:"expect_trade_state,omitempty"`
	Attach            string             `json:"attach,omitempty"`
	BankType          string             `json:"bank_type,omitempty"`
	GoodsTag          string             `json:"goods_tag,omitempty"`
	CouponFee         int64              `json:"coupon_fee,omitempty"`
	CouponCount       int64              `json:"coupon_count,omitempty"`
	CouponInfos       []*WxpayCouponInfo `json:"coupon_infos,omitempty"`
	ProductID         string             `json:"product_id,omitempty"`
	PrepareID         string             `json:"prepare_id,omitempty"`
	TradeStateDesc    string             `json:"trade_state_desc,omitempty"`
	LimitPay          string             `json:"limit_pay,omitempty"`
}

type OrderContent struct {
	OutTradeNo         string    `json:"out_trade_no,omitempty"`
	TransactionID      string    `json:"transaction_id,omitempty"`
	TradeType          TradeType `json:"trade_type,omitempty"`
	AuthorCode         string    `json:"author_code,omitempty"`
	CodeUrl            string    `json:"code_url,omitempty"`
	TimeExpire         *uint64   `json:"time_expire,omitempty"`
	TimeEnd            *uint64   `json:"time_end,omitempty"`
	NotifyUrl          string    `json:"notify_url,omitempty"`
	NonceStr           string    `json:"nonce_str,omitempty"`
	CreateTime         *uint64   `json:"create_time,omitempty"`
	LastUpdateTime     *uint64   `json:"last_update_time,omitempty"`
	IsTransforming     *bool     `json:"is_transforming,omitempty"`
	TotalFee           int64     `json:"total_fee,omitempty"`
	FeeType            string    `json:"fee_type,omitempty"`
	CashFee            int64     `json:"cash_fee,omitempty"`
	CashFeeType        string    `json:"cash_fee_type,omitempty"`
	SettlementTotalFee int64     `json:"settlement_total_fee,omitempty"`
	Body               string    `json:"body,omitempty"`
	Detail             string    `json:"detail,omitempty"`
	// oneof order_content_ext {
	WxpayOrderContentExt  *WxpayOrderContentExt  `json:"wxpay_order_content_ext,omitempty"`
	AlipayOrderContentExt *AlipayOrderContentExt `json:"alipay_order_content_ext,omitempty"`
}

type AlipayFundBill struct {
	FundChannel string `json:"fund_channel,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	RealAmount  int64  `json:"real_amount,omitempty"`
}

type AlipayOrderContentExt struct {
	CurrentTradeState    AlipayOrderState       `json:"current_trade_state,omitempty"`
	ExpectTradeState     AlipayOrderState       `json:"expect_trade_state,omitempty"`
	VoucherDetailList    []*AlipayVoucherDetail `json:"voucher_detail_list,omitempty"`
	FundBillList         []*AlipayFundBill      `json:"fund_bill_list,omitempty"`
	DiscountableAmount   int64                  `json:"discountable_amount,omitempty"`
	UndiscountableAmount int64                  `json:"undiscountable_amount,omitempty"`
	PointAmount          int64                  `json:"point_amount,omitempty"`
	InvoiceAmount        int64                  `json:"invoice_amount,omitempty"`
	ProductCode          string                 `json:"product_code,omitempty"`
	Subject              string                 `json:"subject,omitempty"`
	RoyaltyInfo          string                 `json:"royalty_info,omitempty"`
	SendPayDate          int64                  `json:"send_pay_date,omitempty"`
	IndustrySepcDetail   string                 `json:"industry_sepc_detail,omitempty"`
	ExtendParams         string                 `json:"extend_params,omitempty"`
	EnablePayChannels    string                 `json:"enable_pay_channels,omitempty"`
	// 只透传不处理
	DisablePayChannels  string  `json:"disable_pay_channels,omitempty"`
	DiscountGoodsDetail string  `json:"discount_goods_detail,omitempty"`
	BuyerLogonID        string  `json:"buyer_logon_id,omitempty"`
	SellerID            string  `json:"seller_id,omitempty"`
	SellerEmail         string  `json:"seller_email,omitempty"`
	GmtRefund           *uint64 `json:"gmt_refund,omitempty"`
	GmtClose            *uint64 `json:"gmt_close,omitempty"`
	RefundFee           int64   `json:"refund_fee,omitempty"`
	OutBizNo            string  `json:"out_biz_no,omitempty"`
}

type Authen struct {
	AuthenType AuthenType `json:"authen_type,omitempty"`
	AuthenCode string     `json:"authen_code,omitempty"`
}

type Signature struct {
	SignType SignType `json:"sign_type,omitempty"`
	Sign     string   `json:"sign,omitempty"`
}

type AuthenInfo struct {
	A *Authen    `json:"a,omitempty"`
	S *Signature `json:"s,omitempty"`
}

type ClientIntResult struct {
	Interface Interface `json:"interface,omitempty"`
	ResponseStatus
	LogID      uint     `json:"log_id,omitempty"`
	TimeCost   uint     `json:"time_cost,omitempty"`
	StartTime  *uint64  `json:"start_time,omitempty"`
	DomainName string   `json:"domain_name,omitempty"`
	Logs       []string `json:"logs,omitempty"`
	Level      int      `json:"level,omitempty"`
	Channel    int      `json:"channel,omitempty"`
}

type CloudPayMachineInfo struct {
	Ip                string `json:"ip,omitempty"`
	CpuNum            int    `json:"cpu_num,omitempty"`
	TotalMem          int64  `json:"total_mem,omitempty"`
	MemUsage          int    `json:"mem_usage,omitempty"`
	TaskMemleakageCnt int    `json:"task_memleakage_cnt,omitempty"`
}

type UncompressedMonitorInfo struct {
	ClientIntResults []*ClientIntResult `json:"client_int_results,omitempty"`
	MachineInfo      string             `json:"machine_info,omitempty"`
}

type MicroPayRequest struct {
	PayMchKey   *PayMchKey   `json:"pay_mch_key,omitempty"`
	PayContent  *PayContent  `json:"pay_content,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
}

type MicroPayResponseDetail struct {
	NonceStr     string        `json:"nonce_str,omitempty"`
	PayMchKey    *PayMchKey    `json:"pay_mch_key,omitempty"`
	OrderContent *OrderContent `json:"order_content,omitempty"`
}

type MicroPayResponse struct {
	ResponseStatus
	InternalStatus uint                    `json:"internal_status,omitempty"`
	LogID          uint                    `json:"log_id,omitempty"`
	MicroPay       *MicroPayResponseDetail `json:"micro_pay,omitempty"`
}

type QueryOrderRequest struct {
	PayMchKey   *PayMchKey   `json:"pay_mch_key,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	OutTradeNo  string       `json:"out_trade_no,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
	TradeType   TradeType    `json:"trade_type,omitempty"`
}

type QueryOrderResponseDetail struct {
	NonceStr     string        `json:"nonce_str,omitempty"`
	PayMchKey    *PayMchKey    `json:"pay_mch_key,omitempty"`
	OrderContent *OrderContent `json:"order_content,omitempty"`
}

type QueryOrderResponse struct {
	ResponseStatus
	InternalStatus uint                      `json:"internal_status,omitempty"`
	LogID          uint                      `json:"log_id,omitempty"`
	QueryOrder     *QueryOrderResponseDetail `json:"query_order,omitempty"`
}

type ReverseRequest struct {
	PayMchKey   *PayMchKey   `json:"pay_mch_key,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	OutTradeNo  string       `json:"out_trade_no,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
}

type ReverseResponseDetail struct {
	NonceStr string `json:"nonce_str,omitempty"`
}

type ReverseResponse struct {
	ResponseStatus
	InternalStatus uint                   `json:"internal_status,omitempty"`
	LogID          uint                   `json:"log_id,omitempty"`
	Reverse        *ReverseResponseDetail `json:"reverse,omitempty"`
}

type MonitorReportRequest struct {
	OutMchID    string       `json:"out_mch_id,omitempty"`
	OutSubMchID string       `json:"out_sub_mch_id,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	Interval    uint         `json:"interval,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
	IsCompress  *bool        `json:"is_compress,omitempty"`
	// 二选一
	CompressType            CompressType             `json:"compress_type,omitempty"`
	CompressedMonitorInfo   string                   `json:"compressed_monitor_info,omitempty"`
	UncompressedMonitorInfo *UncompressedMonitorInfo `json:"uncompressed_monitor_info,omitempty"`
	OutShopID               string                   `json:"out_shop_id,omitempty"`
}

type MonitorReportResponseDetail struct {
	NonceStr string `json:"nonce_str,omitempty"`
	Interval uint   `json:"interval,omitempty"`
}

type MonitorReportResponse struct {
	ResponseStatus
	InternalStatus          uint                         `json:"internal_status,omitempty"`
	LogID                   uint                         `json:"log_id,omitempty"`
	UploadClientMonitorInfo *MonitorReportResponseDetail `json:"upload_client_monitor_info,omitempty"`
}

type MachineConfReportRequest struct {
	OutMchID    string       `json:"out_mch_id,omitempty"`
	OutSubMchID string       `json:"out_sub_mch_id,omitempty"`
	OutShopID   string       `json:"out_shop_id,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	MachineInfo string       `json:"machine_info,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
}

type MachineConfResponseDetail struct {
	NonceStr string `json:"nonce_str,omitempty"`
}

type MachineConfResponse struct {
	ResponseStatus
	InternalStatus       uint                       `json:"internal_status,omitempty"`
	LogID                uint                       `json:"log_id,omitempty"`
	UploadClientConfInfo *MachineConfResponseDetail `json:"upload_client_conf_info,omitempty"`
}

type StaffInfo struct {
	StaffID   string `json:"staff_id,omitempty"`
	StaffName string `json:"staff_name,omitempty"`
	Remark    string `json:"remark,omitempty"`
}

type DeviceInfo struct {
	DeviceID   string `json:"device_id,omitempty"`
	DeviceType int    `json:"device_type,omitempty"`
	Remark     string `json:"remark,omitempty"`
}

type ShopInfo struct {
	ShopID         string        `json:"shop_id,omitempty"`
	ShopName       string        `json:"shop_name,omitempty"`
	BranchName     string        `json:"branch_name,omitempty"`
	Province       string        `json:"province,omitempty"`
	City           string        `json:"city,omitempty"`
	District       string        `json:"district,omitempty"`
	Address        string        `json:"address,omitempty"`
	CoordinateType int           `json:"coordinate_type,omitempty"`
	Longitude      string        `json:"longitude,omitempty"`
	Latitude       string        `json:"latitude,omitempty"`
	Height         string        `json:"height,omitempty"`
	Phone          string        `json:"phone,omitempty"`
	DeviceInfos    []*DeviceInfo `json:"device_infos,omitempty"`
	StaffInfos     []*StaffInfo  `json:"staff_infos,omitempty"`
	OutShopID      string        `json:"out_shop_id,omitempty"`
	FeeType        string        `json:"fee_type,omitempty"`
}

type SetShopInfoRequest struct {
	OutMchID    string      `json:"out_mch_id,omitempty"`
	OutSubMchID string      `json:"out_sub_mch_id,omitempty"`
	NonceStr    string      `json:"nonce_str,omitempty"`
	ShopInfos   []*ShopInfo `json:"shop_infos,omitempty"`
	IsAll       *bool       `json:"is_all,omitempty"`
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

type QueryShopInfoRequest struct {
	OutMchID    string `json:"out_mch_id,omitempty"`
	OutSubMchID string `json:"out_sub_mch_id,omitempty"`
	OutShopID   string `json:"out_shop_id,omitempty"`
	NonceStr    string `json:"nonce_str,omitempty"`
	PageNum     int    `json:"page_num,omitempty"`
	PageSize    int    `json:"page_size,omitempty"`
}

type QueryShopInfoResponseDetail struct {
	NonceStr   string      `json:"nonce_str,omitempty"`
	ShopInfos  []*ShopInfo `json:"shop_infos,omitempty"`
	TotalCount int         `json:"total_count,omitempty"`
}

type QueryShopInfoResponse struct {
	ResponseStatus
	InternalStatus uint                         `json:"internal_status,omitempty"`
	LogID          uint                         `json:"log_id,omitempty"`
	QueryShopInfo  *QueryShopInfoResponseDetail `json:"query_shop_info,omitempty"`
}

type Sdk2CloudPayRequest struct {
	RequestContent string      `json:"request_content,omitempty"`
	AuthenInfo     *AuthenInfo `json:"authen_info,omitempty"`
	SignInfo       *AuthenInfo `json:"sign_info,omitempty"`
}

type CloudPay2SdkResponse struct {
	ResponseContent string      `json:"response_content,omitempty"`
	AuthenInfo      *AuthenInfo `json:"authen_info,omitempty"`
}

type PingResponse struct {
	ResponseStatus
}

type WxpayRefundOrderContentExt struct {
	State               WxpayRefundOrderState `json:"state,omitempty"`
	CashRefundFee       int64                 `json:"cash_refund_fee,omitempty"`
	SettlementRefundFee int64                 `json:"settlement_refund_fee,omitempty"`
	CouponRefundFee     int64                 `json:"coupon_refund_fee,omitempty"`
	CouponRefundCount   int64                 `json:"coupon_refund_count,omitempty"`
	CouponRefundInfos   []*WxpayCouponInfo    `json:"coupon_refund_infos,omitempty"`
	RefundAccount       string                `json:"refund_account,omitempty"`
	RefundChannel       string                `json:"refund_channel,omitempty"`
	RefundRecvAccount   string                `json:"refund_recv_account,omitempty"`
}

type RefundContent struct {
	OutTradeNo    string `json:"out_trade_no,omitempty"`
	OutRefundNo   string `json:"out_refund_no,omitempty"`
	TotalFee      int64  `json:"total_fee,omitempty"`
	RefundFee     int64  `json:"refund_fee,omitempty"`
	RefundFeeType string `json:"refund_fee_type,omitempty"`
	RefundReason  string `json:"refund_reason,omitempty"`
	// 第三方支付平台特有信息
	// oneof refund_order_content_ext {
	WxpayRefundOrderContentExt  *WxpayRefundOrderContentExt  `json:"wxpay_refund_order_content_ext,omitempty"`
	AlipayRefundOrderContentExt *AlipayRefundOrderContentExt `json:"alipay_refund_order_content_ext,omitempty"`
}

type AlipayVoucherDetail struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Type           string `json:"type,omitempty"`
	Amount         int64  `json:"amount,omitempty"`
	MerchantAmount int64  `json:"merchant_amount,omitempty"`
	OtherAmount    int64  `json:"other_amount,omitempty"`
	Memo           string `json:"memo,omitempty"`
}

type RefundRequest struct {
	PayMchKey     *PayMchKey     `json:"pay_mch_key,omitempty"`
	RefundContent *RefundContent `json:"refund_content,omitempty"`
	OrderClient   *OrderClient   `json:"order_client,omitempty"`
	NonceStr      string         `json:"nonce_str,omitempty"`
}

type RefundOrderContent struct {
	OutRefundNo            string    `json:"out_refund_no,omitempty"`
	RefundID               string    `json:"refund_id,omitempty"`
	OutTradeNo             string    `json:"out_trade_no,omitempty"`
	TradeType              TradeType `json:"trade_type,omitempty"`
	NonceStr               string    `json:"nonce_str,omitempty"`
	CreateTime             *uint64   `json:"create_time,omitempty"`
	LastUpdateTime         *uint64   `json:"last_update_time,omitempty"`
	IsTransforming         *bool     `json:"is_transforming,omitempty"`
	TotalFee               int64     `json:"total_fee,omitempty"`
	RefundFee              int64     `json:"refund_fee,omitempty"`
	RefundFeeType          string    `json:"refund_fee_type,omitempty"`
	RefundSuccessTime      *uint64   `json:"refund_success_time,omitempty"`
	RefundReason           string    `json:"refund_reason,omitempty"`
	InnerRefundSuccessTime *uint64   `json:"inner_refund_success_time,omitempty"`
	// 第三方支付平台特有信息
	// oneof refund_order_content_ext
	WxpayRefundOrderContentExt  *WxpayRefundOrderContentExt  `json:"wxpay_refund_order_content_ext,omitempty"`
	AlipayRefundOrderContentExt *AlipayRefundOrderContentExt `json:"alipay_refund_order_content_ext,omitempty"`
}

type AlipayRefundOrderContentExt struct {
	FundChange           string                 `json:"fund_change,omitempty"`
	GmtRefundPay         *uint64                `json:"gmt_refund_pay,omitempty"`
	RefundDetailItemList []*AlipayFundBill      `json:"refund_detail_item_list,omitempty"`
	RefundStatus         AlipayRefundOrderState `json:"refund_status,omitempty"`
}

type RefundResponseDetail struct {
	NonceStr           string              `json:"nonce_str,omitempty"`
	PayMchKey          *PayMchKey          `json:"pay_mch_key,omitempty"`
	RefundOrderContent *RefundOrderContent `json:"refund_order_content,omitempty"`
}

type RefundResponse struct {
	ResponseStatus
	InternalStatus uint                  `json:"internal_status,omitempty"`
	LogID          uint                  `json:"log_id,omitempty"`
	Refund         *RefundResponseDetail `json:"refund,omitempty"`
}

type QueryRefundRequest struct {
	PayMchKey   *PayMchKey   `json:"pay_mch_key,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	OutTradeNo  string       `json:"out_trade_no,omitempty"`
	OutRefundNo string       `json:"out_refund_no,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
}

type QueryRefundResponseDetail struct {
	NonceStr           string                `json:"nonce_str,omitempty"`
	PayMchKey          *PayMchKey            `json:"pay_mch_key,omitempty"`
	RefundOrderContent []*RefundOrderContent `json:"refund_order_content,omitempty"`
}

type QueryRefundResponse struct {
	ResponseStatus
	InternalStatus   uint                       `json:"internal_status,omitempty"`
	LogID            uint                       `json:"log_id,omitempty"`
	QueryRefundOrder *QueryRefundResponseDetail `json:"query_refund_order,omitempty"`
}

type ScanCodePayRequest struct {
	PayMchKey   *PayMchKey   `json:"pay_mch_key,omitempty"`
	PayContent  *PayContent  `json:"pay_content,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
}

type ScanCodePayResponseDetail struct {
	NonceStr string `json:"nonce_str,omitempty"`
	CodeUrl  string `json:"code_url,omitempty"`
}

type ScanCodePayResponse struct {
	ResponseStatus
	InternalStatus uint                       `json:"internal_status,omitempty"`
	LogID          uint                       `json:"log_id,omitempty"`
	ScanCodePay    *ScanCodePayResponseDetail `json:"scan_code_pay,omitempty"`
}

type CloseOrderRequest struct {
	PayMchKey   *PayMchKey   `json:"pay_mch_key,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	OutTradeNo  string       `json:"out_trade_no,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
	TradeType   TradeType    `json:"trade_type,omitempty"`
}

type CloseOrderResponseDetail struct {
	NonceStr string `json:"nonce_str,omitempty"`
}

type CloseOrderResponse struct {
	ResponseStatus
	InternalStatus uint                      `json:"internal_status,omitempty"`
	LogID          uint                      `json:"log_id,omitempty"`
	CloseOrder     *CloseOrderResponseDetail `json:"close_order,omitempty"`
}

type ClientInitializeRequest struct {
	OutSubMchID           string                `json:"out_sub_mch_id,omitempty"`
	OutShopID             string                `json:"out_shop_id,omitempty"`
	AuthorizeTerminalType AuthorizeTerminalType `json:"authorize_terminal_type,omitempty"`
}

type ClientInitializeResponseDetail struct {
	Activator string  `json:"activator,omitempty"`
	Token     string  `json:"token,omitempty"`
	QrCodeUrl string  `json:"qr_code_url,omitempty"`
	Ttl       *uint64 `json:"ttl,omitempty"`
}

type ClientInitializeResponse struct {
	ResponseStatus
	InternalStatus   uint                            `json:"internal_status,omitempty"`
	LogID            uint                            `json:"log_id,omitempty"`
	ClientInitialize *ClientInitializeResponseDetail `json:"client_initialize,omitempty"`
}

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
	RefundAuth *bool  `json:"refund_auth,omitempty"`
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
	InitializationGranted *bool                      `json:"initialization_granted,omitempty"`
	IDentity              string                     `json:"identity,omitempty"`
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

type QueryIDentityMappingRequest struct {
	// Types that are valid to be assigned to Condition:
	//	*QueryIDentityMappingRequest_Mch_
	//	*QueryIDentityMappingRequest_Out_
	Condition isQueryIDentityMappingRequest_Condition
}

type isQueryIDentityMappingRequest_Condition interface {
	isQueryIDentityMappingRequest_Condition()
}

type QueryIDentityMappingRequest_Mch_ struct {
	Mch *QueryIDentityMappingRequest_Mch `protobuf:"bytes,1,opt,name=mch,oneof"`
}

type QueryIDentityMappingRequest_Out_ struct {
	Out *QueryIDentityMappingRequest_Out `protobuf:"bytes,2,opt,name=out,oneof"`
}

func (*QueryIDentityMappingRequest_Mch_) isQueryIDentityMappingRequest_Condition() {}

func (*QueryIDentityMappingRequest_Out_) isQueryIDentityMappingRequest_Condition() {}

//

type QueryIDentityMappingRequest_Mch struct {
	MchID    string `json:"mch_id,omitempty"`
	SubMchID string `json:"sub_mch_id,omitempty"`
	ShopID   string `json:"shop_id,omitempty"`
}

type QueryIDentityMappingRequest_Out struct {
	OutMchID    string `json:"out_mch_id,omitempty"`
	OutSubMchID string `json:"out_sub_mch_id,omitempty"`
	OutShopID   string `json:"out_shop_id,omitempty"`
}

type QueryIDentityMappingResponseDetail struct {
	OutMchID       string   `json:"out_mch_id,omitempty"`
	OutSubMchID    string   `json:"out_sub_mch_id,omitempty"`
	OutShopID      string   `json:"out_shop_id,omitempty"`
	NonceStr       string   `json:"nonce_str,omitempty"`
	WxAuthorCodes  []string `json:"wx_author_codes,omitempty"`
	AliAuthorCodes []string `json:"ali_author_codes,omitempty"`
}

type QueryIDentityMappingResponse struct {
	ResponseStatus
	InternalStatus       uint                                `json:"internal_status,omitempty"`
	LogID                uint                                `json:"log_id,omitempty"`
	QueryIDentityMapping *QueryIDentityMappingResponseDetail `json:"query_identity_mapping,omitempty"`
}

type OrderDetailRequest struct {
	PayPlatform      int     `json:"pay_platform,omitempty"`
	OutSubMchID      string  `json:"out_sub_mch_id,omitempty"`
	OutShopID        string  `json:"out_shop_id,omitempty"`
	StaffID          string  `json:"staff_id,omitempty"`
	QueryOrderType   int     `json:"query_order_type,omitempty"`
	TradeType        int     `json:"trade_type,omitempty"`
	StartTime        *uint64 `json:"start_time,omitempty"`
	EndTime          *uint64 `json:"end_time,omitempty"`
	PageNum          uint    `json:"page_num,omitempty"`
	PageSize         uint    `json:"page_size,omitempty"`
	EnableOverview   *bool   `json:"enable_overview,omitempty"`
	DeviceID         string  `json:"device_id,omitempty"`
	QueryCashierType int     `json:"query_cashier_type,omitempty"`
}

type WxpayOrderMchExt struct {
	MchID          string `json:"mch_id,omitempty"`
	SubMchID       string `json:"sub_mch_id,omitempty"`
	ShopID         string `json:"shop_id,omitempty"`
	AppID          string `json:"app_id,omitempty"`
	SubAppID       string `json:"sub_app_id,omitempty"`
	OpenID         string `json:"open_id,omitempty"`
	SubOpenID      string `json:"sub_open_id,omitempty"`
	IsSubscribe    *bool  `json:"is_subscribe,omitempty"`
	SubIsSubscribe *bool  `json:"sub_is_subscribe,omitempty"`
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
	Order *Order `protobuf:"bytes,1,opt,name=order,oneof"`
}

type Receipt_RefundOrder struct {
	RefundOrder *RefundOrder `protobuf:"bytes,2,opt,name=refund_order,json=refundOrder,oneof"`
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

type AdminInfo struct {
	Name string `json:"name,omitempty"`
}

type UpdateMchInfoss struct {
	MchAdminInfos []*AdminInfo `json:"mch_admin_infos,omitempty"`
}

type UpdateMchInfo struct {
	OutMchID      string           `json:"out_mch_id,omitempty"`
	PayPlatform   PayPlatform      `json:"pay_platform,omitempty"`
	MchID         string           `json:"mch_id,omitempty"`
	MchUpdateInfo *UpdateMchInfoss `json:"mch_update_info,omitempty"`
	PageNum       uint             `json:"page_num,omitempty"`
	PageSize      uint             `json:"page_size,omitempty"`
}

type UpdateMchInfoData struct {
	Data     string `json:"data,omitempty"`
	LoginUin string `json:"login_uin,omitempty"`
}

type ClientSubMchSalesOverviewRequest struct {
	OutMchID    string                                         `json:"out_mch_id,omitempty"`
	OutSubMchID string                                         `json:"out_sub_mch_id,omitempty"`
	Shops       []*ClientSubMchSalesOverviewRequest_ShopEntity `json:"shops,omitempty"`
	PayPlatform int                                            `json:"pay_platform,omitempty"`
	TradeType   int                                            `json:"trade_type,omitempty"`
	StatType    int                                            `json:"stat_type,omitempty"`
	StartTime   *uint64                                        `json:"start_time,omitempty"`
	EndTime     *uint64                                        `json:"end_time,omitempty"`
	PageNum     uint                                           `json:"page_num,omitempty"`
	PageSize    uint                                           `json:"page_size,omitempty"`
	StatRole    int                                            `json:"stat_role,omitempty"`
}

type ClientSubMchSalesOverviewRequest_ShopEntity struct {
	OutShopID string   `json:"out_shop_id,omitempty"`
	StaffIDs  []string `json:"staff_ids,omitempty"`
}

type OrderStatClientInfo struct {
	Date               *uint64 `json:"date,omitempty"`
	MchUin             string  `json:"mch_uin,omitempty"`
	OutMchID           string  `json:"out_mch_id,omitempty"`
	MchID              string  `json:"mch_id,omitempty"`
	MchSubUin          string  `json:"mch_sub_uin,omitempty"`
	OutSubMchID        string  `json:"out_sub_mch_id,omitempty"`
	SubMchID           string  `json:"sub_mch_id,omitempty"`
	PayPlatform        int     `json:"pay_platform,omitempty"`
	OutShopID          string  `json:"out_shop_id,omitempty"`
	ShopID             string  `json:"shop_id,omitempty"`
	StaffID            string  `json:"staff_id,omitempty"`
	TotalCount         int64   `json:"total_count,omitempty"`
	TotalAmount        int64   `json:"total_amount,omitempty"`
	SuccessCount       int64   `json:"success_count,omitempty"`
	SuccessAmount      int64   `json:"success_amount,omitempty"`
	SettleAmount       int64   `json:"settle_amount,omitempty"`
	RefundCount        int64   `json:"refund_count,omitempty"`
	RefundAmount       int64   `json:"refund_amount,omitempty"`
	NetCount           int64   `json:"net_count,omitempty"`
	NetAmount          int64   `json:"net_amount,omitempty"`
	DistinctPayerCount int64   `json:"distinct_payer_count,omitempty"`
	PaySettleAmount    int64   `json:"pay_settle_amount,omitempty"`
	RevokeSettleAmount int64   `json:"revoke_settle_amount,omitempty"`
	RefundSettleAmount int64   `json:"refund_settle_amount,omitempty"`
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

type ResponseStatus struct {
	Status      uint   `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r *ResponseStatus) ResponseStatus() *ResponseStatus {
	return r
}

type IsResponseStatus interface {
	ResponseStatus() *ResponseStatus
}
