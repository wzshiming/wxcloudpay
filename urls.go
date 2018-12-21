package wxcloudpay

const (
	PingURL          = `ping`                       //ping
	MicroPayURL      = `micro_pay`                  //刷卡支付的url
	QueryOrderURL    = `query_order`                //查单的url
	ReverseURL       = `reverse`                    //撤单的url
	MonitorReportURL = `upload_client_monitor_info` //监控上报url
	MachineConfURL   = `upload_client_conf_info`    //机器配置上报url
	RefundURL        = `refund`                     //申请退款url
	QueryRefundURL   = `query_refund_order`         //退款查询url
	ScanCodePayURL   = `scan_code_pay`              //扫码支付url
	CloseOrderURL    = `close_order`                //关单url
	SetShopInfoURL   = `set_sub_mch_shop_info`      //设置门店的url
	QueryShopInfoURL = `query_sub_mch_shop_info`    //查询门店的url
	OrderDetailURL   = `client_order_detail`        //客户端订单查询 url
)
