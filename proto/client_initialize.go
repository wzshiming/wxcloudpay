package proto

const ClientInitializeURL = `client_initialize`

type ClientInitializeRequest struct {
	OutSubMchID           string                `json:"out_sub_mch_id,omitempty"`
	OutShopID             string                `json:"out_shop_id,omitempty"`
	AuthorizeTerminalType AuthorizeTerminalType `json:"authorize_terminal_type,omitempty"`
}

type ClientInitializeResponseDetail struct {
	Activator string `json:"activator,omitempty"`
	Token     string `json:"token,omitempty"`
	QrCodeUrl string `json:"qr_code_url,omitempty"`
	Ttl       int    `json:"ttl,omitempty"`
}

type ClientInitializeResponse struct {
	ResponseStatus
	InternalStatus   uint                            `json:"internal_status,omitempty"`
	LogID            uint                            `json:"log_id,omitempty"`
	ClientInitialize *ClientInitializeResponseDetail `json:"client_initialize,omitempty"`
}
