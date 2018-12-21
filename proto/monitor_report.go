package proto

const MonitorReportURL = `upload_client_monitor_info` //监控上报url

type MonitorReportRequest struct {
	OutMchID    string       `json:"out_mch_id,omitempty"`
	OutSubMchID string       `json:"out_sub_mch_id,omitempty"`
	OrderClient *OrderClient `json:"order_client,omitempty"`
	Interval    uint         `json:"interval,omitempty"`
	NonceStr    string       `json:"nonce_str,omitempty"`
	IsCompress  bool         `json:"is_compress,omitempty"`
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
