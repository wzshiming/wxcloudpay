package wxcloudpay

import (
	"encoding/json"
	"errors"
	"fmt"
	"unsafe"

	"github.com/wzshiming/requests"
	"github.com/wzshiming/wxcloudpay/authen"
	"github.com/wzshiming/wxcloudpay/proto"
)

var client = requests.NewClient().
	SetLogLevel(requests.LogMessageAll)

var (
	ErrResponseValidation = errors.New("cloudpay: Response validation failed")
)

type CloudPay struct {
	cli       *requests.Request
	authenKey []byte
}

func NewCloudPay(authenKey, signatureKey []byte) *CloudPay {
	return NewCloudPayWithURL(proto.BaseURL, authenKey, signatureKey)
}

func NewCloudPayWithURL(baseURL string, authenKey, signatureKey []byte) *CloudPay {
	return &CloudPay{
		cli:       client.NewRequest().SetURLByStr(baseURL),
		authenKey: authenKey,
	}
}

func (c *CloudPay) requests(url string, req, resp interface{}, auth bool) error {
	ctx, err := json.Marshal(req)
	if err != nil {
		return err
	}
	requestsContent := *(*string)(unsafe.Pointer(&ctx))

	authenInfo := &proto.AuthenInfo{
		A: &proto.Authen{
			AuthenCode: authen.HmacSha256(ctx, c.authenKey),
			AuthenType: proto.AuthenTypeHmacSha256,
		},
	}

	request := &proto.Sdk2CloudPayRequest{
		RequestContent: requestsContent,
		AuthenInfo:     authenInfo,
	}

	body, err := c.cli.Clone().
		SetJSON(request).
		Post(url)
	if err != nil {
		return err
	}
	if !auth {
		return json.Unmarshal(body.Body(), resp)
	}
	response := &proto.CloudPay2SdkResponse{}

	err = json.Unmarshal(body.Body(), &response)
	if err != nil {
		return err
	}

	if response.ResponseContent == "" {
		return ErrResponseValidation
	}

	responseContent := *(*[]byte)(unsafe.Pointer(&response.ResponseContent))

	// 返回验证
	if authenInfo := response.AuthenInfo; authenInfo != nil {
		if a := authenInfo.A; a != nil {
			if a.AuthenType == proto.AuthenTypeHmacSha256 {
				if authen.HmacSha256(responseContent, c.authenKey) != a.AuthenCode {
					return ErrResponseValidation
				}
			}
		}
	}

	err = json.Unmarshal(responseContent, resp)
	if err != nil {
		return err
	}
	if resp, ok := resp.(proto.IsResponseStatus); ok {
		responseStatus := resp.ResponseStatus()
		if responseStatus.Status != 0 {
			return fmt.Errorf("Error %d: %s", responseStatus.Status, responseStatus.Description)
		}
	}
	return nil
}

func (c *CloudPay) Ping() error {
	resp := &proto.ResponseStatus{}
	err := c.requests(proto.PingURL, nil, &resp, false)
	if err != nil {
		return err
	}
	return nil
}

func (c *CloudPay) CloseOrder(req *proto.CloseOrderRequest) (*proto.CloseOrderResponseDetail, error) {
	resp := &proto.CloseOrderResponse{}
	err := c.requests(proto.CloseOrderURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.CloseOrder, nil
}

func (c *CloudPay) MicroPay(req *proto.MicroPayRequest) (*proto.MicroPayResponseDetail, error) {
	resp := &proto.MicroPayResponse{}
	err := c.requests(proto.MicroPayURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.MicroPay, nil
}

func (c *CloudPay) OrderDetail(req *proto.OrderDetailRequest) (*proto.OrderDetailResponseDetail, error) {
	resp := &proto.OrderDetailResponse{}
	err := c.requests(proto.OrderDetailURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.OrderDetailQuery, nil
}

func (c *CloudPay) QueryOrder(req *proto.QueryOrderRequest) (*proto.QueryOrderResponseDetail, error) {
	resp := &proto.QueryOrderResponse{}
	err := c.requests(proto.QueryOrderURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.QueryOrder, nil
}

func (c *CloudPay) QueryRefund(req *proto.QueryRefundRequest) (*proto.QueryRefundResponseDetail, error) {
	resp := &proto.QueryRefundResponse{}
	err := c.requests(proto.QueryRefundURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.QueryRefundOrder, nil
}

func (c *CloudPay) QueryShopInfo(req *proto.QueryShopInfoRequest) (*proto.QueryShopInfoResponseDetail, error) {
	resp := &proto.QueryShopInfoResponse{}
	err := c.requests(proto.QueryShopInfoURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.QueryShopInfo, nil
}

func (c *CloudPay) QuerySubMchInfo(req *proto.QuerySubMchInfoRequest) (*proto.QuerySubMchInfoResponseDetail, error) {
	resp := &proto.QuerySubMchInfoResponse{}
	err := c.requests(proto.QuerySubMchInfoURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.QuerySubMchInfo, nil
}

func (c *CloudPay) Refund(req *proto.RefundRequest) (*proto.RefundResponseDetail, error) {
	resp := &proto.RefundResponse{}
	err := c.requests(proto.RefundURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.Refund, nil
}

func (c *CloudPay) ScanCodePay(req *proto.ScanCodePayRequest) (*proto.ScanCodePayResponseDetail, error) {
	resp := &proto.ScanCodePayResponse{}
	err := c.requests(proto.ScanCodePayURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.ScanCodePay, nil
}

func (c *CloudPay) SetShopInfo(req *proto.SetShopInfoRequest) (*proto.SetShopInfoResponseDetail, error) {
	resp := &proto.SetShopInfoResponse{}
	err := c.requests(proto.SetShopInfoURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.SetShopInfo, nil
}

func (c *CloudPay) QueryIdentityMapping(req *proto.QueryIdentityMappingRequest) (*proto.QueryIdentityMappingResponseDetail, error) {
	resp := &proto.QueryIdentityMappingResponse{}
	err := c.requests(proto.QueryIdentityMappingURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.QueryIdentityMapping, nil
}

func (c *CloudPay) ClientInitialize(req *proto.ClientInitializeRequest) (*proto.ClientInitializeResponseDetail, error) {
	resp := &proto.ClientInitializeResponse{}
	err := c.requests(proto.ClientInitializeURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.ClientInitialize, nil
}

func (c *CloudPay) ClientSubMchSalesOverview(req *proto.ClientSubMchSalesOverviewRequest) (*proto.ClientSubMchSalesOverviewResponseDetail, error) {
	resp := &proto.ClientSubMchSalesOverviewResponse{}
	err := c.requests(proto.ClientSubMchSalesOverviewURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.ClientSubMchSalesOverview, nil
}

func (c *CloudPay) ClientTryGetInitializers(req *proto.ClientTryGetInitializersRequest) (*proto.ClientTryGetInitializersResponseDetail, error) {
	resp := &proto.ClientTryGetInitializersResponse{}
	err := c.requests(proto.ClientTryGetInitializersURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.ClientTryGetInitializers, nil
}

func (c *CloudPay) UploadClientConfInfo(req *proto.UploadClientConfInfoRequest) (*proto.UploadClientConfInfoResponseDetail, error) {
	resp := &proto.UploadClientConfInfoResponse{}
	err := c.requests(proto.UploadClientConfInfoURL, req, &resp, true)
	if err != nil {
		return nil, err
	}
	return resp.UploadClientConfInfo, nil
}
