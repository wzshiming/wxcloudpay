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
	InterfaceQueryIdentityMapping      Interface = 213
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
