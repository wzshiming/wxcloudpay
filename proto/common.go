package proto

type Sdk2CloudPayRequest struct {
	RequestContent string      `json:"request_content,omitempty"`
	AuthenInfo     *AuthenInfo `json:"authen_info,omitempty"`
	SignInfo       *AuthenInfo `json:"sign_info,omitempty"`
}

type CloudPay2SdkResponse struct {
	ResponseContent string      `json:"response_content,omitempty"`
	AuthenInfo      *AuthenInfo `json:"authen_info,omitempty"`
}

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
	Price     int    `json:"price,omitempty"`
}

type Detail struct {
	CostPrice   int            `json:"cost_price,omitempty"`
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
	DiscountableAmount   int    `json:"discountable_amount,omitempty"`
	UndiscountableAmount int    `json:"undiscountable_amount,omitempty"`
	ProductCode          string `json:"product_code,omitempty"`
	RoyaltyInfo          string `json:"royalty_info,omitempty"`
	ExtendParams         string `json:"extend_params,omitempty"`
	DisablePayChannels   string `json:"disable_pay_channels,omitempty"`
	TimeoutExpress       string `json:"timeout_express,omitempty"`
}

type PayContent struct {
	OutTradeNo string `json:"out_trade_no,omitempty"`
	AuthorCode string `json:"author_code,omitempty"`
	TimeExpire int    `json:"time_expire,omitempty"`
	TotalFee   int    `json:"total_fee,omitempty"`
	FeeType    string `json:"fee_type,omitempty"`
	Body       string `json:"body,omitempty"`
	Detail     string `json:"detail,omitempty"`
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
	CouponFee     int    `json:"coupon_fee,omitempty"`
	CouponType    string `json:"coupon_type,omitempty"`
}

type WxpayOrderContentExt struct {
	CurrentTradeState WxpayOrderState    `json:"current_trade_state,omitempty"`
	ExpectTradeState  WxpayOrderState    `json:"expect_trade_state,omitempty"`
	Attach            string             `json:"attach,omitempty"`
	BankType          string             `json:"bank_type,omitempty"`
	GoodsTag          string             `json:"goods_tag,omitempty"`
	CouponFee         int                `json:"coupon_fee,omitempty"`
	CouponCount       int                `json:"coupon_count,omitempty"`
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
	TimeExpire         int       `json:"time_expire,omitempty"`
	TimeEnd            int       `json:"time_end,omitempty"`
	NotifyUrl          string    `json:"notify_url,omitempty"`
	NonceStr           string    `json:"nonce_str,omitempty"`
	CreateTime         int       `json:"create_time,omitempty"`
	LastUpdateTime     int       `json:"last_update_time,omitempty"`
	IsTransforming     bool      `json:"is_transforming,omitempty"`
	TotalFee           int       `json:"total_fee,omitempty"`
	FeeType            string    `json:"fee_type,omitempty"`
	CashFee            int       `json:"cash_fee,omitempty"`
	CashFeeType        string    `json:"cash_fee_type,omitempty"`
	SettlementTotalFee int       `json:"settlement_total_fee,omitempty"`
	Body               string    `json:"body,omitempty"`
	Detail             string    `json:"detail,omitempty"`
	// oneof order_content_ext {
	WxpayOrderContentExt  *WxpayOrderContentExt  `json:"wxpay_order_content_ext,omitempty"`
	AlipayOrderContentExt *AlipayOrderContentExt `json:"alipay_order_content_ext,omitempty"`
}

type AlipayFundBill struct {
	FundChannel string `json:"fund_channel,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	RealAmount  int    `json:"real_amount,omitempty"`
}

type AlipayOrderContentExt struct {
	CurrentTradeState    AlipayOrderState       `json:"current_trade_state,omitempty"`
	ExpectTradeState     AlipayOrderState       `json:"expect_trade_state,omitempty"`
	VoucherDetailList    []*AlipayVoucherDetail `json:"voucher_detail_list,omitempty"`
	FundBillList         []*AlipayFundBill      `json:"fund_bill_list,omitempty"`
	DiscountableAmount   int                    `json:"discountable_amount,omitempty"`
	UndiscountableAmount int                    `json:"undiscountable_amount,omitempty"`
	PointAmount          int                    `json:"point_amount,omitempty"`
	InvoiceAmount        int                    `json:"invoice_amount,omitempty"`
	ProductCode          string                 `json:"product_code,omitempty"`
	Subject              string                 `json:"subject,omitempty"`
	RoyaltyInfo          string                 `json:"royalty_info,omitempty"`
	SendPayDate          int                    `json:"send_pay_date,omitempty"`
	IndustrySepcDetail   string                 `json:"industry_sepc_detail,omitempty"`
	ExtendParams         string                 `json:"extend_params,omitempty"`
	EnablePayChannels    string                 `json:"enable_pay_channels,omitempty"`
	// 只透传不处理
	DisablePayChannels  string `json:"disable_pay_channels,omitempty"`
	DiscountGoodsDetail string `json:"discount_goods_detail,omitempty"`
	BuyerLogonID        string `json:"buyer_logon_id,omitempty"`
	SellerID            string `json:"seller_id,omitempty"`
	SellerEmail         string `json:"seller_email,omitempty"`
	GmtRefund           int    `json:"gmt_refund,omitempty"`
	GmtClose            int    `json:"gmt_close,omitempty"`
	RefundFee           int    `json:"refund_fee,omitempty"`
	OutBizNo            string `json:"out_biz_no,omitempty"`
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
	StartTime  int      `json:"start_time,omitempty"`
	DomainName string   `json:"domain_name,omitempty"`
	Logs       []string `json:"logs,omitempty"`
	Level      int      `json:"level,omitempty"`
	Channel    int      `json:"channel,omitempty"`
}

type CloudPayMachineInfo struct {
	Ip                string `json:"ip,omitempty"`
	CpuNum            int    `json:"cpu_num,omitempty"`
	TotalMem          int    `json:"total_mem,omitempty"`
	MemUsage          int    `json:"mem_usage,omitempty"`
	TaskMemleakageCnt int    `json:"task_memleakage_cnt,omitempty"`
}

type UncompressedMonitorInfo struct {
	ClientIntResults []*ClientIntResult `json:"client_int_results,omitempty"`
	MachineInfo      string             `json:"machine_info,omitempty"`
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

type WxpayRefundOrderContentExt struct {
	State               WxpayRefundOrderState `json:"state,omitempty"`
	CashRefundFee       int                   `json:"cash_refund_fee,omitempty"`
	SettlementRefundFee int                   `json:"settlement_refund_fee,omitempty"`
	CouponRefundFee     int                   `json:"coupon_refund_fee,omitempty"`
	CouponRefundCount   int                   `json:"coupon_refund_count,omitempty"`
	CouponRefundInfos   []*WxpayCouponInfo    `json:"coupon_refund_infos,omitempty"`
	RefundAccount       string                `json:"refund_account,omitempty"`
	RefundChannel       string                `json:"refund_channel,omitempty"`
	RefundRecvAccount   string                `json:"refund_recv_account,omitempty"`
}

type RefundContent struct {
	OutTradeNo    string `json:"out_trade_no,omitempty"`
	OutRefundNo   string `json:"out_refund_no,omitempty"`
	TotalFee      int    `json:"total_fee,omitempty"`
	RefundFee     int    `json:"refund_fee,omitempty"`
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
	Amount         int    `json:"amount,omitempty"`
	MerchantAmount int    `json:"merchant_amount,omitempty"`
	OtherAmount    int    `json:"other_amount,omitempty"`
	Memo           string `json:"memo,omitempty"`
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
